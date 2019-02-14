// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package app

import (
	"flag"
	"log"
	"os"
	"strconv"
)

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
	cfgSrcExtraDefault = ""
	cfgSrcExtraUsage   = "Extra info for config source, e.g., execute URL for vRO"
)

func setEnvString(val *string, key string, dflt string) {
	str, ok := os.LookupEnv(key)
	if !ok {
		*val = dflt
	} else {
		*val = str
	}
}

func setEnvInt(val *int, key string, dflt int) {
	var str string
	sdflt := strconv.Itoa(dflt)
	setEnvString(&str, key, sdflt)
	if i, err := strconv.ParseInt(str, 0, 64); err != nil {
		*val = dflt
	} else {
		*val = int(i)
	}
}

func setEnvBool(val *bool, key string, dflt bool) {
	var str string
	var sdflt string
	if dflt {
		sdflt = "true"
	} else {
		sdflt = "false"
	}
	setEnvString(&str, key, sdflt)
	if b, err := strconv.ParseBool(str); err != nil {
		*val = dflt
	} else {
		*val = b
	}
}

func configureFromEnv() {
	log.Println("---- Setting Config From Environment ----")
	setEnvInt(&ListenPort, "LISTENPORT", listenPortDefault)
	log.Printf("Configure ListenPort to: %v\n", ListenPort)
	setEnvInt(&Port, "PORT", portDefault)
	log.Printf("Configure Port to: %v\n", Port)
	setEnvString(&Host, "HOST", hostDefault)
	log.Printf("Configure Host to: %v\n", Host)
	setEnvString(&Admin, "USER", adminDefault)
	log.Printf("Configure Admin to: %v\n", Admin)
	setEnvString(&Passwd, "PASSWD", passwordDefault)
	log.Println("Configured Passwd")
	setEnvString(&DBName, "DBNAME", dbNameDefault)
	log.Printf("Configure DBName to: %v\n", DBName)
	setEnvString(&DBType, "DBTYPE", dbTypeDefault)
	log.Printf("Configure DBType to: %v\n", DBType)
	setEnvString(&ContentRoot, "TPLPATH", contentRootDefault)
	log.Printf("Configure ContentRoot to: %v\n", ContentRoot)
	setEnvBool(&Insecure, "INSECURE", insecureDefault)
	log.Printf("Configure Insecure to: %v\n", Insecure)
	setEnvString(&CfgType, "CFGTYPE", cfgTypeDefault)
	log.Printf("Configure CfgType to: %v\n", CfgType)
	setEnvString(&CfgSrc, "CFGSRC", cfgSrcDefault)
	log.Printf("Configure CfgSrc to: %v\n", CfgSrc)
	setEnvString(&CfgSrcExtra, "CFGSRCEXTRA", cfgSrcExtraDefault)
	log.Printf("Configure CfgSrcExtra to: %v\n", CfgSrcExtra)
}

// Initialize the flags processor with default values and help messages.
func initFlags() {
	log.Println("---- Setting Config From Command Line ----")
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
	flag.StringVar(&DBType, "dbtype", dbTypeDefault, dbTypeUsage)
	flag.StringVar(&DBType, "d", dbTypeDefault, dbTypeUsage+" (shorthand)")
	flag.StringVar(&ContentRoot, "tplpath", contentRootDefault, contentRootUsage)
	flag.StringVar(&ContentRoot, "t", contentRootDefault, contentRootUsage+" (shorthand)")
	flag.BoolVar(&Insecure, "insecure", insecureDefault, insecureUsage)
	flag.BoolVar(&Insecure, "i", insecureDefault, insecureUsage+" (shorthand)")
	flag.StringVar(&CfgType, "cfgtype", cfgTypeDefault, cfgTypeUsage)
	flag.StringVar(&CfgType, "c", cfgTypeDefault, cfgTypeUsage+" (shorthand)")
	flag.StringVar(&CfgSrc, "cfgsrc", cfgSrcDefault, cfgSrcUsage)
	flag.StringVar(&CfgSrc, "g", cfgSrcDefault, cfgSrcUsage+" (shorthand)")
	flag.StringVar(&CfgSrcExtra, "cfgsrcextra", cfgSrcExtraDefault, cfgSrcExtraUsage)
	flag.StringVar(&CfgSrcExtra, "e", cfgSrcExtraDefault, cfgSrcExtraUsage+" (shorthand)")
}

// Process application (command line) flags.
func Init() {
	initFlags()
	flag.Parse()
	configureFromEnv()
}

func init() {
	log.Println("Initialized app package.")
}
