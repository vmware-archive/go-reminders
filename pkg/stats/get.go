// Copyright 2015 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package stats

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func (s *Stats) Get(w rest.ResponseWriter, r *rest.Request) {
	s.AddHit(r.RequestURI)
	s.lock.Lock()
	err := w.WriteJson(s.hits)
	s.lock.Unlock()
	if err != nil {
		rest.Error(w, err.Error(), 503)
	}
}
