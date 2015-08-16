// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package app

import (
	"flag"
	"log"
)

// Global application context variables.
var (
	ListenPort  int
	DBPort      int
	DBHost      string
	DBAdmin     string
	DBPasswd    string
	DBName      string
	ContentRoot string
	APIAddress  string
)

// Initialize the flags processor with default values and help messages.
func initFlags() {
	const (
		listenPortDefault  = 8080
		listenPortUsage    = "port on which to listen for HTTP requests"
		dbPortDefault      = 3306
		dbPortUsage        = "port to use for DB connections"
		dbHostDefault      = "localhost"
		dbHostUsage        = "database (host) address"
		dbAdminDefault     = "vmware"
		dbAdminUsage       = "database admin user (with power to create databases)"
		dbPasswordDefault  = "vmware"
		dbPasswordUsage    = "database admin password"
		dbNameDefault      = ""
		dbNameUsage        = "database name to connect to (empty for new database)"
		contentRootDefault = "."
		contentRootUsage   = "path to (content) templates, skeleton, etc."
		apiAddressDefault  = "localhost"
		apiAddressUsage    = "address for reminders api (default localhost)"
	)

	flag.IntVar(&ListenPort, "listenport", listenPortDefault, listenPortUsage)
	flag.IntVar(&ListenPort, "l", listenPortDefault, listenPortUsage+" (shorthand)")
	flag.IntVar(&DBPort, "dbport", dbPortDefault, dbPortUsage)
	flag.IntVar(&DBPort, "p", dbPortDefault, dbPortUsage+" (shorthand)")
	flag.StringVar(&DBHost, "dbhost", dbHostDefault, dbHostUsage)
	flag.StringVar(&DBHost, "a", dbHostDefault, dbHostUsage+" (shorthand)")
	flag.StringVar(&DBAdmin, "dbuser", dbAdminDefault, dbAdminUsage)
	flag.StringVar(&DBAdmin, "u", dbAdminDefault, dbAdminUsage+" (shorthand)")
	flag.StringVar(&DBPasswd, "dbpasswd", dbPasswordDefault, dbPasswordUsage)
	flag.StringVar(&DBPasswd, "s", dbPasswordDefault, dbPasswordUsage+" (shorthand)")
	flag.StringVar(&DBName, "dbname", dbNameDefault, dbNameUsage)
	flag.StringVar(&DBName, "n", dbNameDefault, dbNameUsage+" (shorthand)")
	flag.StringVar(&ContentRoot, "tplpath", contentRootDefault, contentRootUsage)
	flag.StringVar(&ContentRoot, "t", contentRootDefault, contentRootUsage+" (shorthand)")
}

// Process application (command line) flags. Note this happens automatically.
// No need to explicitly call this function (in fact that is a bad idea).
func init() {
	initFlags()
	flag.Parse()
	log.Printf("Initialized app package.")
}
