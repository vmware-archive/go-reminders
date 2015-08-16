// Copyright (c) 2013-2015 Antoine Imbert
// Copyright (c) 2015 VMware
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package reminders

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

// Update (REST Put) a Reminder using id as the key.
func (s *Storage) Put(w rest.ResponseWriter, r *rest.Request) {

	id := r.PathParam("id")
	reminder := Reminder{}
	if s.DB.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Reminder{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reminder.Message = updated.Message

	if err := s.DB.Save(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&reminder)
}

// Update (REST Put) a Reminder using guid as the key.
func (s *Storage) PutGuid(w rest.ResponseWriter, r *rest.Request) {

	guid := r.PathParam("guid")
	reminder := Reminder{}
	if s.DB.Where(&Reminder{Guid: guid}).First(&reminder).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Reminder{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reminder.Message = updated.Message

	if err := s.DB.Save(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&reminder)
}
