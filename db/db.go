// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package db

import (
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const MySQL = "mysql"

type DB struct {
	addr   string
	port   int
	admin  string
	passwd string
	name   string
}

func init() {
	log.Println("Initialized db package.")
}

// Return a properly formed connection URI for connecting to the server, but
// not a specific database. Useful for, for example, creating the database
// rather than running queries on an already created database.
func (db *DB) dbURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		db.admin, db.passwd, db.addr, db.port)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Execute a command on the DB (similar to mysql -e '...' ...).
func (db *DB) exec(cmd string) error {
	conn, err := sql.Open(MySQL, db.dbURI())
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(cmd)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

// Generate a random string for use as database objects.
func randomName() (string, error) {
	b := make([]byte, 12)
	_, err := rand.Read(b)
	if err != nil {
		log.Println(err)
	}

	b64 := base64.URLEncoding.EncodeToString(b)

	s := "_" + fmt.Sprintf("%x", md5.Sum([]byte(b64)))

	return s, err
}

// Create and return a DB struct, also create the database if necessary.
func New(host string, port int, user string, passwd string, name string) (DB, error) {
	db := DB{
		addr:   host,
		port:   port,
		admin:  user,
		passwd: passwd,
		name:   name,
	}

	wantNewDB := false
	if len(db.name) == 0 {
		db.name, _ = randomName()
		wantNewDB = true
	}

	if wantNewDB {
		err := db.Create()
		if err != nil {
			log.Println("Error creating database: %v.", err)
			return db, err
		}
	}

	return db, nil
}

// Create the database represented by DB.
func (db *DB) Create() error {
	log.Printf("Creating database: %s\n", db.name)
	return db.exec("CREATE DATABASE " + db.name)
}

// Drop the database represented by DB.
func (db *DB) Drop() error {
	log.Printf("Dropping database: %s\n", db.name)
	return db.exec("DROP DATABASE " + db.name)
}
