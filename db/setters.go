// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package db

func (db *DB) SetAddress(host string) {
	if host != "" {
		db.addr = host
	}
}

func (db *DB) SetPort(port int) {
	if port > 0 {
		db.port = port
	}
}

func (db *DB) SetAdmin(admin string) {
	if admin != "" {
		db.admin = admin
	}
}

func (db *DB) SetPasswd(passwd string) {
	if passwd != "" {
		db.passwd = passwd
	}
}

func (db *DB) SetName(name string) {
	if name != "" {
		db.name = name
	}
}
