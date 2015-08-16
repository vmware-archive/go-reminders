// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package stats

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func (s *Stats) Get(w rest.ResponseWriter, r *rest.Request) {
	m := make(map[string]interface{})
	s.AddHit(r.RequestURI)
	s.lock.RLock()
	m["hits"] = s.hits
	s.lock.RUnlock()
	err := w.WriteJson(m)
	if err != nil {
		rest.Error(w, err.Error(), 503)
	}
}
