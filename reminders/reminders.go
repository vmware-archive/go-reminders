// Copyright (c) 2013-2015 Antoine Imbert
// Copyright (c) 2015 VMware
//
// License: MIT (see https://github.com/tdhite/vmorld/LICENSE).
//
package reminders

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/tdhite/vmworld/db"
	"log"
	"time"
)

// Reminder is serializable as json (tagged) and also SQL tags provide for
// best fit database storage (see the Go sql provider for details).
type Reminder struct {
	Id        int64     `json:"id"`
	Guid      string    `sql:"size:48;unique_index:idx_guid;size=32" json:"guid"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}

type Storage struct {
	DB gorm.DB
}

func init() {
	log.Println("Initialized reminders package.")
}

// Initialize and returns a new DB storage structure.
func New(db db.DB) Storage {
	s := Storage{}
	s.InitDB(db)
	s.InitSchema()
	return s
}

// Initialize the SQL database and open it.
func (s *Storage) InitDB(rDB db.DB) {
	var err error
	s.DB, err = gorm.Open(db.MySQL, rDB.ConnectURI())
	if err != nil {
		log.Fatalf("Database connect error: '%v'.", err)
	}
	s.DB.LogMode(true)
}

// Initialize the SQL database schema and open it.
func (s *Storage) InitSchema() {
	s.DB.AutoMigrate(&Reminder{})
}

// Convert a JSON array of Reminders to Go slice and return.
func ArrayFromJson(jsonData []byte) ([]Reminder, error) {
	var reminders []Reminder
	err := json.Unmarshal([]byte(jsonData), &reminders)
	if err != nil {
		log.Printf("%T\n%s\n%#v\n", err, err, err)
		switch v := err.(type) {
		case *json.SyntaxError:
			log.Println(string(jsonData[v.Offset-40 : v.Offset]))
		}
	}

	return reminders, err
}

// Convert a JSON Reminder to Go struct and return.
func FromJson(jsonData []byte) (Reminder, error) {
	var r Reminder
	err := json.Unmarshal([]byte(jsonData), &r)
	if err != nil {
		log.Printf("%T\n%s\n%#v\n", err, err, err)
		switch v := err.(type) {
		case *json.SyntaxError:
			log.Println(string(jsonData[v.Offset-40 : v.Offset]))
		}
	}

	return r, err
}
