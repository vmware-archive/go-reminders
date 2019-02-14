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

type MemDB struct {
	byGuid map[string]*Reminder
	byId   map[int64]*Reminder
	index  int64
}

////// Storage Interface

// Initialize and returns a new storage.
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
	delete(db.byGuid, r.Guid)
	delete(db.byId, r.Id)
}

// Initialize the database and open it.
func (db *MemDB) InitDB() error {
	db.byGuid = make(map[string]*Reminder)
	db.byId = make(map[int64]*Reminder)
	db.index = 0
	return nil
}

func (db *MemDB) Close() error {
	db.InitDB()
	return nil
}

// Drop the database represented by DB.
func (db *MemDB) Drop() error {
	db.InitDB()
	return nil
}

func (db *MemDB) DeleteId(id int64) (Reminder, error) {
	if r, ok := db.byId[id]; ok {
		db.deleteReminder(r)
		return *r, nil
	}

	return Reminder{}, errors.New("Id does not exist.")
}

func (db *MemDB) DeleteGuid(guid string) (Reminder, error) {
	if r, ok := db.byGuid[guid]; ok {
		db.deleteReminder(r)
		return *r, nil
	}

	return Reminder{}, errors.New("Guid does not exist.")
}

func (db *MemDB) GetAll() (*[]Reminder, error) {
	r := make([]Reminder, 0)
	for _, v := range db.byId {
		r = append(r, *v)
	}

	return &r, nil
}

func (db *MemDB) GetId(id int64) (Reminder, error) {
	if r, ok := db.byId[id]; ok {
		return *r, nil
	}

	return Reminder{}, errors.New("Guid does not exist.")
}

func (db *MemDB) GetGuid(guid string) (Reminder, error) {
	if r, ok := db.byGuid[guid]; ok {
		return *r, nil
	}

	return Reminder{}, errors.New("Guid does not exist.")
}

func (db *MemDB) Save(r Reminder) error {
	// Check if this exists already
	rem, err := db.GetGuid(r.Guid)
	if err != nil {
		if r.Id == 0 {
			db.index++
			r.Id = db.index
		}
	} else {
		r.Id = rem.Id
		r.CreatedAt = rem.CreatedAt
		r.DeletedAt = rem.DeletedAt
		r.UpdatedAt = time.Now()
	}

	db.byGuid[r.Guid] = &r
	db.byId[r.Id] = &r
	return nil
}
