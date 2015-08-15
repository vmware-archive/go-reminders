// Copyright (c) 2013-2015 Antoine Imbert
// Copyright (c) 2015 VMware
//
// License: MIT (see https://github.com/tdhite/vmorld/LICENSE).
//
package reminders

import (
	"github.com/ant0ine/go-json-rest/rest"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// Create a new Reminder (REST Post request). This also injects a guid.
func (s *Storage) Post(w rest.ResponseWriter, r *rest.Request) {
	reminder := Reminder{}
	if err := r.DecodeJsonPayload(&reminder); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(reminder.Guid) == 0 {
		u := uuid.NewV4()
		reminder.Guid = u.String()
	}
	if err := s.DB.Save(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&reminder)
}
