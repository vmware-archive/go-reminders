// Copyright 2015 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package app

import (
	"flag"
	"log"

	"github.com/vmwaresamples/go-reminders/pkg/globals"
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
		passwordDefault    = "VMware1!"
		passwordUsage      = "password for authentication"
		dbNameDefault      = ""
		dbNameUsage        = "database name to connect to (empty for new database)"
		dbTypeDefault      = "mem"
		dbTypeUsage        = "type of database to use (current support for mem/mysql)"
		contentRootDefault = "."
		contentRootUsage   = "path to (content) templates, skeleton, etc."
		insecureDefault    = false
		insecureUsage      = "allow insecure auth (skip tls verify)"
		cfgTypeDefault     = "none"
		cfgTypeUsage       = "select configuration source type (current support for none/vro/etcd)"
		cfgSrcDefault      = ""
		cfgSrcUsage        = "host (or ip address) of configuration source"
	)

	flag.IntVar(&globals.ListenPort, "listenport", listenPortDefault, listenPortUsage)
	flag.IntVar(&globals.ListenPort, "l", listenPortDefault, listenPortUsage+" (shorthand)")
	flag.IntVar(&globals.Port, "port", portDefault, portUsage)
	flag.IntVar(&globals.Port, "p", portDefault, portUsage+" (shorthand)")
	flag.StringVar(&globals.Host, "host", hostDefault, hostUsage)
	flag.StringVar(&globals.Host, "h", hostDefault, hostUsage+" (shorthand)")
	flag.StringVar(&globals.Admin, "user", adminDefault, adminUsage)
	flag.StringVar(&globals.Admin, "u", adminDefault, adminUsage+" (shorthand)")
	flag.StringVar(&globals.Passwd, "passwd", passwordDefault, passwordUsage)
	flag.StringVar(&globals.Passwd, "s", passwordDefault, passwordUsage+" (shorthand)")
	flag.StringVar(&globals.DBName, "dbname", dbNameDefault, dbNameUsage)
	flag.StringVar(&globals.DBName, "n", dbNameDefault, dbNameUsage+" (shorthand)")
	flag.StringVar(&globals.DBType, "dbtype", dbTypeDefault, dbTypeUsage)
	flag.StringVar(&globals.DBType, "d", dbTypeDefault, dbTypeUsage+" (shorthand)")
	flag.StringVar(&globals.ContentRoot, "tplpath", contentRootDefault, contentRootUsage)
	flag.StringVar(&globals.ContentRoot, "t", contentRootDefault, contentRootUsage+" (shorthand)")
	flag.BoolVar(&globals.Insecure, "insecure", insecureDefault, insecureUsage)
	flag.BoolVar(&globals.Insecure, "i", insecureDefault, insecureUsage+" (shorthand)")
	flag.StringVar(&globals.CfgType, "cfgtype", cfgTypeDefault, cfgTypeUsage)
	flag.StringVar(&globals.CfgType, "c", cfgTypeDefault, cfgTypeUsage+" (shorthand)")
	flag.StringVar(&globals.CfgSrc, "cfgsrc", cfgSrcDefault, cfgSrcUsage)
	flag.StringVar(&globals.CfgSrc, "e", cfgSrcDefault, cfgSrcUsage+" (shorthand)")
}

// Process application (command line) flags.
func Init() {
	initFlags()
	flag.Parse()
}

func init() {
	log.Println("Initialized app package.")
}
