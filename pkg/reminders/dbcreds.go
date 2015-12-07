// Copyright 2019 VMware, Inc. All Rights Reserved.
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

	"github.com/vmwaresamples/go-reminders/pkg/globals"
)

type DBCreds struct {
	addr   string
	port   int
	admin  string
	passwd string
	name   string
}

func (c *DBCreds) Init(host string, port int, admin string, passwd string, dbname string) {
	c.addr = host
	c.port = port
	c.admin = admin
	c.passwd = passwd
	c.name = dbname
}

// [NOTE]: credentials should be pre-populated at config source.
func (c *DBCreds) FetchCredentials(ctype string, d *DBCreds) error {
	switch ctype {
	case "etcd":
		k := NewEtcd(globals.CfgSrc)
		k.GetDBCreds(d)
		log.Println("DB: %v", *d)
	case "none":
		log.Println("No credential type supplied, assuming command line args.")
	case "vro":
		v := NewVro(globals.Admin, globals.Passwd, globals.Insecure)
		err := v.GetDBCredsBasicAuth(globals.CfgSrc, d)
		if err != nil {
			log.Fatalf("Failed to connect and obtain creds from vRO. %v.\n", err)
		}
		log.Printf("DB: %v", *d)
	default:
		msg := fmt.Sprintf("Unsupported configuration source of type %s.\n", globals.CfgSrc)
		log.Println(msg)
		return errors.New(msg)
	}

	return nil
}

/////// getters / setters

func (db *DBCreds) Address() string {
	return db.addr
}

func (db *DBCreds) SetAddress(host string) {
	if host != "" {
		db.addr = host
	}
}

func (db *DBCreds) Port() int {
	return db.port
}

func (db *DBCreds) SetPort(port int) {
	if port > 0 {
		db.port = port
	}
}

func (db *DBCreds) Admin() string {
	return db.admin
}

func (db *DBCreds) SetAdmin(admin string) {
	if admin != "" {
		db.admin = admin
	}
}

func (db *DBCreds) Passwd() string {
	return db.passwd
}

func (db *DBCreds) SetPasswd(passwd string) {
	if passwd != "" {
		db.passwd = passwd
	}
}

func (db *DBCreds) Name() string {
	return db.name
}

func (db *DBCreds) SetName(name string) {
	if name != "" {
		db.name = name
	}
}
