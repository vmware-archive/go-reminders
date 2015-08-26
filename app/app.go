// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package app

import (
	"flag"
	"github.com/tdhite/go-reminders/stats"
	"log"
)

// Global application context variables.
var (
	ListenPort  int
	Port        int
	Host        string
	Admin       string
	Passwd      string
	DBName      string
	ContentRoot string
	APIAddress  string
	VROUrl      string
	Insecure    bool
	Stats       stats.Stats = stats.New()
)

// Initialize the flags processor with default values and help messages.
func initFlags() {
	const (
		listenPortDefault  = 8080
		listenPortUsage    = "port on which to listen for HTTP requests"
		portDefault        = 3306
		portUsage          = "port to use for connections"
		hostDefault        = "localhost"
		hostUsage          = "database (host) address"
		adminDefault       = "vmware"
		adminUsage         = "username for authentication -- note: db user must have power to create databases)"
		passwordDefault    = "vmware"
		passwordUsage      = "password for authentication"
		dbNameDefault      = ""
		dbNameUsage        = "database name to connect to (empty for new database)"
		contentRootDefault = "."
		contentRootUsage   = "path to (content) templates, skeleton, etc."
		insecureDefault    = false
		insecureUsage      = "allow insecure auth (skip tls verify)"
		vROUrlDefault      = ""
		vROUrlUsage        = "URL to authenticate against REST API for MySQL Access"
	)

	flag.IntVar(&ListenPort, "listenport", listenPortDefault, listenPortUsage)
	flag.IntVar(&ListenPort, "l", listenPortDefault, listenPortUsage+" (shorthand)")
	flag.IntVar(&Port, "port", portDefault, portUsage)
	flag.IntVar(&Port, "p", portDefault, portUsage+" (shorthand)")
	flag.StringVar(&Host, "host", hostDefault, hostUsage)
	flag.StringVar(&Host, "h", hostDefault, hostUsage+" (shorthand)")
	flag.StringVar(&Admin, "user", adminDefault, adminUsage)
	flag.StringVar(&Admin, "u", adminDefault, adminUsage+" (shorthand)")
	flag.StringVar(&Passwd, "passwd", passwordDefault, passwordUsage)
	flag.StringVar(&Passwd, "s", passwordDefault, passwordUsage+" (shorthand)")
	flag.StringVar(&DBName, "dbname", dbNameDefault, dbNameUsage)
	flag.StringVar(&DBName, "n", dbNameDefault, dbNameUsage+" (shorthand)")
	flag.StringVar(&ContentRoot, "tplpath", contentRootDefault, contentRootUsage)
	flag.StringVar(&ContentRoot, "t", contentRootDefault, contentRootUsage+" (shorthand)")
	flag.StringVar(&VROUrl, "vrourl", vROUrlDefault, vROUrlUsage)
	flag.StringVar(&VROUrl, "v", vROUrlDefault, vROUrlUsage+" (shorthand)")
	flag.BoolVar(&Insecure, "insecure", insecureDefault, insecureUsage)
	flag.BoolVar(&Insecure, "i", insecureDefault, insecureUsage+" (shorthand)")
}

// Process application (command line) flags. Note this happens automatically.
// No need to explicitly call this function (in fact that is a bad idea).
func init() {
	initFlags()
	flag.Parse()
	log.Printf("Initialized app package.")
}
