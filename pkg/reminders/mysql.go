// Package reminders holds the application logic to manage reminders (tasks one seeks to remember).
//
// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package reminders

import (
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"

	// Import MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	//MySQL is the name of the driver to use
	MySQL = "mysql"

	// added to the return string from DB.dbURI
	connectFmt = "%s?charset=utf8&parseTime=True"
)

// MySQLDB holds connection information and credentials for the MySQL database.
type MySQLDB struct {
	db    *gorm.DB
	creds DBCreds
}

// Return a properly formed connection URI for connecting to the server, but
// not a specific database. Useful for, for example, creating the database
// rather than running queries on an already created database.
func (db *MySQLDB) dbURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		db.creds.Admin(), db.creds.Passwd(), db.creds.Address(), db.creds.Port())
}

// Return a properly formatted connection URI for the SQL db.
func (db *MySQLDB) connectURI() string {
	open := db.dbURI()
	return open + fmt.Sprintf(connectFmt, db.creds.Name())
}

// Initialize the SQL database schema and open it.
func (db MySQLDB) initSchema() error {
	err := db.db.AutoMigrate(&Reminder{}).Error
	return err
}

// Execute a command on the DB (similar to mysql -e '...' ...).
func (db MySQLDB) exec(cmd string) error {
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

// Create the database represented by DB.
func (db MySQLDB) create() error {
	log.Printf("Creating database: %s\n", db.creds.Name())
	return db.exec("CREATE DATABASE " + db.creds.Name())
}

////// Storage Interface

// NewMySQL Initialize and returns a new storage.
func NewMySQL(c DBCreds) (Storage, error) {
	m := MySQLDB{}
	m.creds = c
	if err := m.InitDB(); err != nil {
		return &m, err
	}
	if err := m.initSchema(); err != nil {
		return &m, err
	}
	return &m, nil
}

//// Storage Implementation

// InitDB Initialize the database and open it.
func (db *MySQLDB) InitDB() error {
	wantNewDB := false
	if len(db.creds.Name()) == 0 {
		name, _ := randomName()
		db.creds.SetName(name)
		wantNewDB = true
	}

	if wantNewDB {
		err := db.create()
		if err != nil {
			log.Printf("Failed to create new database: %v.\n", err)
			return err
		}
	}

	var err error
	db.db, err = gorm.Open(MySQL, db.connectURI())
	if err != nil {
		log.Fatalf("Database connect error: '%v'.", err)
	}

	db.db.LogMode(true)

	return err
}

// Close terminates the connection to the database.
func (db *MySQLDB) Close() error {
	err := db.db.Close()
	return err
}

// Drop the database represented by DB.
func (db *MySQLDB) Drop() error {
	log.Printf("Dropping database: %s\n", db.creds.Name())
	return db.exec("DROP DATABASE " + db.creds.Name())
}

// DeleteID removes the reminder with the given ID from the database.
func (db *MySQLDB) DeleteID(id int64) (Reminder, error) {
	r := Reminder{}
	var err error
	if err = db.db.First(&r, id).Error; err != nil {
		return r, err
	}
	if err = db.db.Delete(r).Error; err != nil {
		r = Reminder{}
	}
	return r, err
}

// DeleteGUID removes the reminder with the given GUID from the database.
func (db *MySQLDB) DeleteGUID(guid string) (Reminder, error) {
	r := Reminder{}
	if err := db.db.Where(&Reminder{GUID: guid}).First(&r).Error; err != nil {
		return Reminder{}, err
	}
	err := db.db.Delete(&r).Error
	return r, err
}

// GetAll returns a list of all reminders in the database.
func (db *MySQLDB) GetAll() (*[]Reminder, error) {
	r := []Reminder{}
	err := db.db.Find(&r).Error
	return &r, err
}

// GetID returns the reminder with the given ID. If not found,
// it will return a not found error.
func (db *MySQLDB) GetID(id int64) (Reminder, error) {
	r := Reminder{}
	err := db.db.First(&r, id).Error
	return r, err
}

// GetGUID returns the reminder with the given UUID. If not found,
// it will return a not found error.
func (db *MySQLDB) GetGUID(guid string) (Reminder, error) {
	r := Reminder{}
	err := db.db.Where(&Reminder{GUID: guid}).First(&r).Error
	return r, err
}

// Save stores the given reminder in the database.
func (db *MySQLDB) Save(r Reminder) error {
	err := db.db.Save(r).Error
	return err
}
