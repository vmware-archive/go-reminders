// Copyright 2017 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package reminders

import (
	"log"

	"github.com/vmwaresamples/go-reminders/pkg/globals"
)

type Storage interface {
	InitDB() error
	Close() error
	Drop() error
	DeleteId(int64) (Reminder, error)
	DeleteGuid(string) (Reminder, error)
	GetAll() (*[]Reminder, error)
	GetId(int64) (Reminder, error)
	GetGuid(string) (Reminder, error)
	Save(r Reminder) error
}

func newStorage() (Storage, error) {
	var d DBCreds

	// Use external configuration source.
	if globals.CfgSrc != "" {
		d = DBCreds{}
		d.SetName(globals.DBName)
		if err := d.FetchCredentials(globals.CfgType, &d); err != nil {
			log.Printf("Failed to connect to DB: %v.\n", err)
			return nil, err
		}
		// No external configuration source, so success depend on arguments.
	} else {
		d.Init(globals.Host, globals.Port, globals.Admin, globals.Passwd, globals.DBName)
	}

	var s Storage
	var err error
	switch globals.DBType {
	case "mem":
		if s, err = NewMemDB(); err != nil {
			log.Fatalf("Failed to open database: %v.\n", err)
		}
	case "mysql":
		if s, err = NewMySQL(d); err != nil {
			log.Fatalf("Failed to open database: %v.\n", err)
		}
	default:
		log.Fatalf("Unsupported database type %s.\n", globals.DBType)
	}

	return s, err
}
