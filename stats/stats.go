// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package stats

import (
	"log"
	"sync"
)

// store hits per URL
type Stats struct {
	hits map[string]int `json:"hits"`
	lock sync.RWMutex
}

func init() {
	log.Println("Initialized stats package.")
}

func New() Stats {
	return Stats{
		hits: make(map[string]int),
		lock: sync.RWMutex{},
	}
}
