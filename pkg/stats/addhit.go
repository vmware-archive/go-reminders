// Package stats holds the statics logic of the go-reminders application
//
// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package stats

import (
	"log"
)

// AddHit increses the number of hits count to the reminders API.
func (s *Stats) AddHit(request string) {
	s.lock.Lock()
	count := s.Hits[request] + 1
	s.Hits[request] = count
	log.Printf("Counting hit: %s -- up to %d\n", request, count)
	s.lock.Unlock()
}
