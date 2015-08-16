// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	// added to the return string from DB.dbURI
	connectFmt = "%s?charset=utf8&parseTime=True"
)

// Return a properly formatted connection URI for the SQL db.
func (db *DB) ConnectURI() string {
	open := db.dbURI()
	return open + fmt.Sprintf(connectFmt, db.name)
}
