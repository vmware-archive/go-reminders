// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package db

func (db *DB) Address() string {
	return db.addr
}

func (db *DB) Port() int {
	return db.port
}

func (db *DB) Admin() string {
	return db.admin
}

func (db *DB) Passwd() string {
	return db.passwd
}

func (db *DB) Name() string {
	return db.name
}
