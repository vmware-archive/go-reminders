// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/vmorld/LICENSE).
//
package stats

import (
	"log"
	"sync"
)

type Stats struct {
	hits     int
	requests []string
	lock     sync.RWMutex
}

func init() {
	log.Println("Initialized stats package.")
}

func New() Stats {
	return Stats{
		hits:     0,
		requests: make([]string, 0),
		lock:     sync.RWMutex{},
	}
}
