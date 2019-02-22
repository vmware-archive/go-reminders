// Package reminders holds the application logic to manage reminders (tasks one seeks to remember).
//
// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
// This software is licensed to you under the MIT license (the "License").
// You may not use this product except in compliance with the MIT License.
//
package reminders

import (
	"errors"
	"fmt"
	"log"
)

// DBCreds holds database credentials and connection information.
type DBCreds struct {
	addr   string
	port   int
	admin  string
	passwd string
	name   string
	dbtype string
	cfgsrc string
	extra  string
}

// Init initializes the DBCreds struct.
func (db *DBCreds) Init(host string, port int, admin string, passwd string, dbtype string, dbname string, cfgsrc string, extra string) {
	db.addr = host
	db.port = port
	db.admin = admin
	db.passwd = passwd
	db.name = dbname
	db.dbtype = dbtype
	db.cfgsrc = cfgsrc
	db.extra = extra
}

// FetchCredentials obtains the database credentials from either Etcd or vRO if
// not provided on the command line
// [NOTE]: credentials should be pre-populated at config source.
func (db *DBCreds) FetchCredentials(insecure bool) error {
	switch db.cfgsrc {
	case "etcd":
		k := NewEtcd(db.cfgsrc)
		if err := k.GetDBCreds(db); err != nil {
			log.Fatalf("Failed to connect and obtain creds from etcd. %v.\n", err)
			return err
		}
	case "none":
		log.Println("No credential type supplied, assuming command line args.")
	case "vro":
		v := NewVro(*db, insecure)
		err := v.GetDBCredsBasicAuth(db)
		if err != nil {
			log.Fatalf("Failed to connect and obtain creds from vRO. %v.\n", err)
			return err
		}
	default:
		msg := fmt.Sprintf("Unsupported configuration source of type %s.\n", db.dbtype)
		log.Println(msg)
		return errors.New(msg)
	}

	return nil
}

/////// getters / setters

// Address is a getter for the databse host in the DBCreds struct.
func (db *DBCreds) Address() string {
	return db.addr
}

// SetAddress is a setter for the database host in the DBCreds struct.
func (db *DBCreds) SetAddress(host string) {
	if host == "" {
		log.Println("Setting DBCreds host to empty string.")
	}
	db.addr = host
}

// Port is a getter for the database host port in the DBCreds struct.
func (db *DBCreds) Port() int {
	return db.port
}

// SetPort is a setter for the database host port in the DBCreds struct.
func (db *DBCreds) SetPort(port int) {
	if port == 0 {
		log.Println("Setting DBCreds port to empty string.")
	}
	db.port = port
}

// Admin is a getter for the database user in the DBCreds struct.
func (db *DBCreds) Admin() string {
	return db.admin
}

// SetAdmin is a setter for the database user in the DBCreds struct.
func (db *DBCreds) SetAdmin(admin string) {
	if admin != "" {
		log.Println("Setting DBCreds admin to empty string.")
	}
	db.admin = admin
}

// Passwd is a getter for the database password in the DBCreds struct.
func (db *DBCreds) Passwd() string {
	return db.passwd
}

// SetPasswd is a setter for the database password in the DBCreds struct.
func (db *DBCreds) SetPasswd(passwd string) {
	if passwd == "" {
		log.Println("Setting DBCreds passwd to empty string.")
	}
	db.passwd = passwd
}

// Name is a getter for the database name in the DBCreds struct.
func (db *DBCreds) Name() string {
	return db.name
}

// SetName is a setter for the database name in the DBCreds struct.
func (db *DBCreds) SetName(name string) {
	if name == "" {
		log.Println("Setting DBCreds name to empty string.")
	}
	db.name = name
}

// DBType is a getter for the database type in the DBCreds struct.
func (db *DBCreds) DBType() string {
	return db.dbtype
}

// SetDBType is a setter for the database type in the DBCreds struct.
// Acceptable values are "mem" or "mysql".
func (db *DBCreds) SetDBType(t string) {
	if t == "" {
		log.Println("Setting DBCreds dbtype to empty string.")
	}
	db.dbtype = t
}

// CfgSrc is a getter for the cfgsrc field in the DBCreds struct.
func (db *DBCreds) CfgSrc() string {
	return db.cfgsrc
}

// SetCfgSrc is a setter for the cfgsrc field in the DBCreds struct.
func (db *DBCreds) SetCfgSrc(src string) {
	if src == "" {
		log.Println("Setting DBCreds cfgsrc to empty string.")
	}
	db.cfgsrc = src
}

// Extra is a getter for the extra field in the DBCreds struct.
func (db *DBCreds) Extra() string {
	return db.extra
}

// SetExtra is a setter for the extra field in the DBCreds struct.
func (db *DBCreds) SetExtra(e string) {
	if e == "" {
		log.Println("Setting DBCreds cfgsrc to empty string.")
	}
	db.extra = e
}
