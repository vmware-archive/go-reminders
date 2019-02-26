// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package app

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/stretchr/graceful"
	"github.com/vmware/go-reminders/internal/template"
	"github.com/vmware/go-reminders/pkg/reminders"
)

func getUrlRoot() string {
	if VHost == "" {
		return "http://localhost:" + strconv.Itoa(ListenPort)
	} else {
		return VHost
	}
}

// Http handler functions for dealing with various html site requests for
// home page, editing, deleting and saving reminder objects.
//
// These are not all that necessary as they are just a trick to use the
// http.ServeMux to create a poor man's URL router. The json stuff uses
// the venerable go-json-router, but the site pages are so simple it's not
// worth writing up a whole router model just for that when we can just 'mux'
// things via separate handlers for each html (site) request.
func templateHomeHandler(w http.ResponseWriter, r *http.Request) {
	Stats.AddHit(r.RequestURI)
	t := template.New(ContentRoot, getUrlRoot(), APIBaseUrl, Stats)
	t.IndexHandler(w, r)
}

func templateEditHandler(w http.ResponseWriter, r *http.Request) {
	Stats.AddHit(r.RequestURI)
	t := template.New(ContentRoot, getUrlRoot(), APIBaseUrl, Stats)
	t.EditHandler(w, r)
}

func templateSaveHandler(w http.ResponseWriter, r *http.Request) {
	Stats.AddHit(r.RequestURI)
	t := template.New(ContentRoot, getUrlRoot(), APIBaseUrl, Stats)
	t.SaveHandler(w, r)
}

func templateDeleteHandler(w http.ResponseWriter, r *http.Request) {
	Stats.AddHit(r.RequestURI)
	t := template.New(ContentRoot, getUrlRoot(), APIBaseUrl, Stats)
	t.DeleteHandler(w, r)
}

func statsHitsHandler(w http.ResponseWriter, r *http.Request) {
	Stats.AddHit(r.RequestURI)
	t := template.New(ContentRoot, getUrlRoot(), APIBaseUrl, Stats)
	t.StatsHitsHandler(w, r)
}

// Called by main, which is just a wrapper for this function. The reason
// is main can't directly pass back a return code to the OS.
func RealMain() int {
	Init()

	c := reminders.DBCreds{}
	c.Init(Host, Port, Admin, Passwd, DBType, DBName, CfgSrc, CfgSrcExtra)

	r, err := reminders.NewReminders(c, Stats, Insecure)
	if err != nil {
		return 1
	}

	// setup JSON request handlers
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		// stats
		rest.Get("/stats/hits", Stats.Get),
		// reminders
		rest.Get("/api/reminders", r.GetAll),
		rest.Get("/api/reminders/byid/:id", r.GetId),
		rest.Get("/api/reminders/:guid", r.GetGuid),
		rest.Post("/api/reminders", r.Post),
		rest.Put("/api/reminders/:guid", r.PutGuid),
		rest.Put("/api/reminders/byid/:id", r.Put),
		rest.Delete("/api/reminders/:guid", r.DeleteGuid),
		rest.Delete("/api/reminders/byid/:id", r.Delete),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	// setup the html page request handlers and mux it all
	mux := http.NewServeMux()
	mux.Handle("/api/", api.MakeHandler())
	mux.Handle("/stats/", api.MakeHandler())
	mux.Handle("/html/skeleton/", http.FileServer(http.Dir(ContentRoot)))
	mux.Handle("/html/tmpl/index", http.HandlerFunc(templateHomeHandler))
	mux.Handle("/html/tmpl/delete", http.HandlerFunc(templateDeleteHandler))
	mux.Handle("/html/tmpl/edit", http.HandlerFunc(templateEditHandler))
	mux.Handle("/html/tmpl/save", http.HandlerFunc(templateSaveHandler))
	mux.Handle("/html/stats/hits", http.HandlerFunc(statsHitsHandler))

	// this runs a server that can handle os signals for clean shutdown.
	server := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    ":" + strconv.Itoa(ListenPort),
			Handler: mux,
		},
		ListenLimit: 1024,
	}

	exitcode := 0
	err = server.ListenAndServe()
	if err != nil {
		log.Println("Shutdown caused by:" + err.Error())
		exitcode = 1
	}

	// Deletes the database -- not strictly necessary so comment out
	// if you want to keep the data. Not that if a database is in fact
	// provided on the command line flags, it does not get deleted, which
	// allows for multiple of this program (service) to run against the
	// same storage backend (mysql at present).
	//	if DBName == "" {
	//		r.Drop()
	//	}

	return exitcode
}
