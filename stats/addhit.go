// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package stats

import (
	"log"
)

func (s *Stats) AddHit(request string) {
	s.lock.Lock()
	count := s.hits[request]
	s.hits[request] = count + 1
	log.Printf("Counting hit: %s -- up to %d\n", request, count)
	s.lock.Unlock()
}
