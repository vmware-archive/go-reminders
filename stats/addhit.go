// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/vmorld/LICENSE).
//
package stats

import (
	"log"
)

func (s *Stats) AddHit(request string) {
	log.Printf("Counting hit: %s\n", request)
	s.lock.Lock()
	s.hits++
	s.lock.Unlock()
}
