// Copyright 2015 VMware, Inc. All Rights Reserved.
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

type Reminders struct {
	s Storage
}

func NewReminders() (Reminders, error) {
	r := Reminders{}

	s, err := newStorage()
	if err != nil {
		return r, err
	}

	r.s = s
	return r, nil
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
