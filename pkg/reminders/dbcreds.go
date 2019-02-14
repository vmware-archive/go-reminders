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

func (c *DBCreds) Init(host string, port int, admin string, passwd string, dbtype string, dbname string, cfgsrc string, extra string) {
	c.addr = host
	c.port = port
	c.admin = admin
	c.passwd = passwd
	c.name = dbname
	c.dbtype = dbtype
	c.cfgsrc = cfgsrc
	c.extra = extra
}

// [NOTE]: credentials should be pre-populated at config source.
func (c *DBCreds) FetchCredentials(insecure bool) error {
	switch c.cfgsrc {
	case "etcd":
		k := NewEtcd(c.cfgsrc)
		if err := k.GetDBCreds(c); err != nil {
			log.Fatalf("Failed to connect and obtain creds from etcd. %v.\n", err)
			return err
		}
	case "none":
		log.Println("No credential type supplied, assuming command line args.")
	case "vro":
		v := NewVro(*c, insecure)
		err := v.GetDBCredsBasicAuth(c)
		if err != nil {
			log.Fatalf("Failed to connect and obtain creds from vRO. %v.\n", err)
			return err
		}
	default:
		msg := fmt.Sprintf("Unsupported configuration source of type %s.\n", c.dbtype)
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
	if host == "" {
		log.Println("Setting DBCreds host to empty string.")
	}
	db.addr = host
}

func (db *DBCreds) Port() int {
	return db.port
}

func (db *DBCreds) SetPort(port int) {
	if port == 0 {
		log.Println("Setting DBCreds port to empty string.")
	}
	db.port = port
}

func (db *DBCreds) Admin() string {
	return db.admin
}

func (db *DBCreds) SetAdmin(admin string) {
	if admin != "" {
		log.Println("Setting DBCreds admin to empty string.")
	}
	db.admin = admin
}

func (db *DBCreds) Passwd() string {
	return db.passwd
}

func (db *DBCreds) SetPasswd(passwd string) {
	if passwd == "" {
		log.Println("Setting DBCreds passwd to empty string.")
	}
	db.passwd = passwd
}

func (db *DBCreds) Name() string {
	return db.name
}

func (db *DBCreds) SetName(name string) {
	if name == "" {
		log.Println("Setting DBCreds name to empty string.")
	}
	db.name = name
}

func (db *DBCreds) DBType() string {
	return db.dbtype
}

func (db *DBCreds) SetDBType(t string) {
	if t == "" {
		log.Println("Setting DBCreds dbtype to empty string.")
	}
	db.dbtype = t
}

func (db *DBCreds) CfgSrc() string {
	return db.cfgsrc
}

func (db *DBCreds) SetCfgSrc(src string) {
	if src == "" {
		log.Println("Setting DBCreds cfgsrc to empty string.")
	}
	db.cfgsrc = src
}

func (db *DBCreds) Extra() string {
	return db.extra
}

func (db *DBCreds) SetExtra(e string) {
	if e == "" {
		log.Println("Setting DBCreds cfgsrc to empty string.")
	}
	db.extra = e
}
