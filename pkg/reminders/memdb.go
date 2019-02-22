// Package reminders holds the application logic to manage reminders (tasks one seeks to remember).
//
// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package reminders

import (
	"errors"
	"time"
)

// MemDB emulates an in-memory database to store reminders.
type MemDB struct {
	byGUID map[string]*Reminder
	byID   map[int64]*Reminder
	index  int64
}

////// Storage Interface

// NewMemDB initialize and returns a new storage.
func NewMemDB() (Storage, error) {
	m := MemDB{}
	m.Close() // this creates the tables

	if err := m.InitDB(); err != nil {
		return &m, err
	}
	return &m, nil
}

//// Storage Implementation
func (db *MemDB) deleteReminder(r *Reminder) {
	delete(db.byGUID, r.GUID)
	delete(db.byID, r.ID)
}

// InitDB initialize the database and open it.
func (db *MemDB) InitDB() error {
	db.byGUID = make(map[string]*Reminder)
	db.byID = make(map[int64]*Reminder)
	db.index = 0
	return nil
}

// Close terminates the connection to the database.
func (db *MemDB) Close() error {
	db.InitDB()
	return nil
}

// Drop the database represented by DB.
func (db *MemDB) Drop() error {
	db.InitDB()
	return nil
}

// DeleteID removes the reminder with the given ID from the database.
func (db *MemDB) DeleteID(id int64) (Reminder, error) {
	if r, ok := db.byID[id]; ok {
		db.deleteReminder(r)
		return *r, nil
	}

	return Reminder{}, errors.New("ID does not exist")
}

// DeleteGUID removes the reminder with the given GUID from the database
func (db *MemDB) DeleteGUID(guid string) (Reminder, error) {
	if r, ok := db.byGUID[guid]; ok {
		db.deleteReminder(r)
		return *r, nil
	}

	return Reminder{}, errors.New("GUID does not exist")
}

// GetAll returns a list of all reminders in the database
func (db *MemDB) GetAll() (*[]Reminder, error) {
	r := make([]Reminder, 0)
	for _, v := range db.byID {
		r = append(r, *v)
	}

	return &r, nil
}

// GetID returns the reminder stored in the database specified by the given ID
func (db *MemDB) GetID(id int64) (Reminder, error) {
	if r, ok := db.byID[id]; ok {
		return *r, nil
	}

	return Reminder{}, errors.New("GUID does not exist")
}

// GetGUID returns the reminder stored in the database specified by the given GUID
func (db *MemDB) GetGUID(guid string) (Reminder, error) {
	if r, ok := db.byGUID[guid]; ok {
		return *r, nil
	}

	return Reminder{}, errors.New("GUID does not exist")
}

// Save stores the given reminder in the database
func (db *MemDB) Save(r Reminder) error {
	// Check if this exists already
	rem, err := db.GetGUID(r.GUID)
	if err != nil {
		if r.ID == 0 {
			db.index++
			r.ID = db.index
		}
	} else {
		r.ID = rem.ID
		r.CreatedAt = rem.CreatedAt
		r.DeletedAt = rem.DeletedAt
		r.UpdatedAt = time.Now()
	}

	db.byGUID[r.GUID] = &r
	db.byID[r.ID] = &r
	return nil
}
