// Copyright 2015 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
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
