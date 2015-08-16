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

// Handle REST Delete request, which presumes Id as the identifying key.
func (s *Storage) Delete(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	reminder := Reminder{}
	if s.DB.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := s.DB.Delete(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Handle REST Delete request using Guid as the identifying key.
func (s *Storage) DeleteGuid(w rest.ResponseWriter, r *rest.Request) {
	guid := r.PathParam("guid")
	reminder := Reminder{}
	if s.DB.Where(&Reminder{Guid: guid}).First(&reminder).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := s.DB.Delete(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
