// Copyright (c) 2013-2015 Antoine Imbert
// Copyright (c) 2015 VMware
//
// License: MIT (see https://github.com/tdhite/vmorld/LICENSE).
//
package reminders

import (
	"github.com/ant0ine/go-json-rest/rest"
)

// Rerieve all Reminders via REST Get request.
func (s *Storage) GetAll(w rest.ResponseWriter, r *rest.Request) {
	reminders := []Reminder{}
	s.DB.Find(&reminders)
	w.WriteJson(&reminders)
}

// Retrieve one Reminder via REST Get request using id as key.
func (s *Storage) Get(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("guid")
	reminder := Reminder{}
	if s.DB.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&reminder)
}

// Retrieve one Reminder via REST Get request using guid as key.
func (s *Storage) GetGuid(w rest.ResponseWriter, r *rest.Request) {
	guid := r.PathParam("guid")
	reminder := Reminder{}
	if s.DB.Where(&Reminder{Guid: guid}).First(&reminder).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&reminder)
}
