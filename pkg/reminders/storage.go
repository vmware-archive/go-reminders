// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package reminders

import (
	"log"
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

func newStorage(creds DBCreds, insecure bool) (Storage, error) {
	// Use external configuration source.
	if creds.CfgSrc() != "" {
		if err := creds.FetchCredentials(insecure); err != nil {
			log.Printf("Failed to connect to db creds source: %v.\n", err)
			return nil, err
		}
		// No external configuration source, so success depend on arguments.
	}

	var s Storage
	var err error
	switch creds.DBType() {
	case "mem":
		if s, err = NewMemDB(); err != nil {
			log.Fatalf("Failed to open database: %v.\n", err)
		}
	case "mysql":
		if s, err = NewMySQL(creds); err != nil {
			log.Fatalf("Failed to open database: %v.\n", err)
		}
	default:
		log.Fatalf("Unsupported database type %s.\n", creds.DBType())
	}

	return s, err
}
