// Package reminders holds the application logic to manage reminders (tasks one seeks to remember).
//
// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Copyright (c) 2013-2015 Antoine Imbert
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package reminders

import (
	"encoding/json"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/vmware/go-reminders/pkg/stats"
)

// Reminder is serializable as json (tagged) and also SQL tags provide for
// best fit database storage (see the Go sql provider for details).
type Reminder struct {
	ID        int64     `json:"id"`
	GUID      string    `sql:"size:48;unique_index:idx_guid;size=32" json:"guid"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}

// Reminders holds the Storage and Stats structs.
type Reminders struct {
	s     Storage
	stats stats.Stats
}

// NewReminders initialize and returns a new Reminders struct.
func NewReminders(creds DBCreds, stats stats.Stats, insecure bool) (Reminders, error) {
	r := Reminders{
		stats: stats,
	}

	s, err := newStorage(creds, insecure)
	if err != nil {
		return r, err
	}

	r.s = s
	return r, nil
}

func newTestReminders(creds DBCreds, insecure bool) (Reminders, error) {
	stats := stats.New()
	return NewReminders(creds, stats, insecure)

}

func newGuid() (uuid.UUID, error) {
	return uuid.NewV4()
}

func init() {
	log.Println("Initialized reminders package.")
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
