// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/stretchr/graceful"
	"github.com/tdhite/go-reminders/app"
	"github.com/tdhite/go-reminders/db"
	"github.com/tdhite/go-reminders/reminders"
	"github.com/tdhite/go-reminders/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Http handler functions for dealing with various html site requests for
// home page, editing, deleting and saving reminder objects.
//
// These are not all that necessary as they are just a trick to use the
// http.ServeMux to create a poor man's URL router. The json stuff uses
// the venerable go-json-router, but the site pages are so simple it's not
// worth writing up a whole router model just for that when we can just 'mux'
// things via separate handlers for each html (site) request.
func templateHomeHandler(w http.ResponseWriter, r *http.Request) {
	app.Stats.AddHit(r.RequestURI)
	t := template.New(app.ContentRoot, app.APIAddress+":"+strconv.Itoa(app.ListenPort))
	t.IndexHandler(w, r)
}

func templateEditHandler(w http.ResponseWriter, r *http.Request) {
	app.Stats.AddHit(r.RequestURI)
	t := template.New(app.ContentRoot, app.APIAddress+":"+strconv.Itoa(app.ListenPort))
	t.EditHandler(w, r)
}

func templateSaveHandler(w http.ResponseWriter, r *http.Request) {
	app.Stats.AddHit(r.RequestURI)
	t := template.New(app.ContentRoot, app.APIAddress+":"+strconv.Itoa(app.ListenPort))
	t.SaveHandler(w, r)
}

func templateDeleteHandler(w http.ResponseWriter, r *http.Request) {
	app.Stats.AddHit(r.RequestURI)
	t := template.New(app.ContentRoot, app.APIAddress+":"+strconv.Itoa(app.ListenPort))
	t.DeleteHandler(w, r)
}

func statsHitsHandler(w http.ResponseWriter, r *http.Request) {
	app.Stats.AddHit(r.RequestURI)
	t := template.New(app.ContentRoot, app.APIAddress+":"+strconv.Itoa(app.ListenPort))
	t.StatsHitsHandler(w, r)
}

// Called by main, which is just a wrapper for this function. The reason
// is main can't directly pass back a return code to the OS.
func realMain() int {
	db, err := db.New(app.DBHost, app.DBPort, app.DBAdmin, app.DBPasswd, app.DBName)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v.\n", err)
	}

	reminders := reminders.New(db)

	// setup JSON request handlers
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		// stats
		rest.Get("/stats/hits", app.Stats.Get),
		// reminders
		rest.Get("/api/reminders", reminders.GetAll),
		rest.Post("/api/reminders", reminders.Post),
		rest.Get("/api/reminders/:guid", reminders.GetGuid),
		rest.Put("/api/reminders/:guid", reminders.PutGuid),
		rest.Delete("/api/reminders/:guid", reminders.DeleteGuid),
		rest.Get("/api/reminders/byid/:id", reminders.Get),
		rest.Put("/api/reminders/byid/:id", reminders.Put),
		rest.Delete("/api/reminders/byid/:id", reminders.Delete),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	// setup the html page request handlers and mux it all
	mux := http.NewServeMux()
	mux.Handle("/api/", api.MakeHandler())
	mux.Handle("/stats/", api.MakeHandler())
	mux.Handle("/html/skeleton/", http.FileServer(http.Dir(app.ContentRoot)))
	mux.Handle("/html/tmpl/index", http.HandlerFunc(templateHomeHandler))
	mux.Handle("/html/tmpl/delete", http.HandlerFunc(templateDeleteHandler))
	mux.Handle("/html/tmpl/edit", http.HandlerFunc(templateEditHandler))
	mux.Handle("/html/tmpl/save", http.HandlerFunc(templateSaveHandler))
	mux.Handle("/html/stats/hits", http.HandlerFunc(statsHitsHandler))

	// this runs a server that can handle os signals for clean shutdown.
	server := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    ":" + strconv.Itoa(app.ListenPort),
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
	if app.DBName == "" {
		db.Drop()
	}

	return exitcode
}

func main() {
	// Delegate to realMain so defered operations can happen (os.Exit exits
	// the program without servicing defer statements)
	os.Exit(realMain())
}
