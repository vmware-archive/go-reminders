// Copyright 2015 VMware, Inc. All Rights Reserved.
// Copyright (c) 2013-2015 Antoine Imbert
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package reminders

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
	uuid "github.com/satori/go.uuid"
	"github.com/vmwaresamples/go-reminders/pkg/globals"
)

func idFromString(s string) (int64, error) {
	id, err := strconv.ParseInt(s, 10, 64)
	return id, err
}

// Retrieve all Reminders via REST Get request.
func (rem *Reminders) GetAll(w rest.ResponseWriter, r *rest.Request) {
	all, err := rem.s.GetAll()
	if err != nil {
		rest.NotFound(w, r)
		return
	}

	w.WriteJson(&all)
}

// Retrieve one Reminder via REST Get request using id as key.
func (rem *Reminders) GetId(w rest.ResponseWriter, r *rest.Request) {
	globals.Stats.AddHit(r.RequestURI)

	id, err := idFromString(r.PathParam("id"))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reminder, err := rem.s.GetId(id)
	if err != nil {
		rest.NotFound(w, r)
		return
	}

	w.WriteJson(&reminder)
}

// Retrieve one Reminder via REST Get request using guid as key.
func (rem *Reminders) GetGuid(w rest.ResponseWriter, r *rest.Request) {
	globals.Stats.AddHit(r.RequestURI)

	guid := r.PathParam("guid")
	reminder, err := rem.s.GetGuid(guid)
	if err != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&reminder)
}

// Create a new Reminder (REST Post request). This also injects a guid.
func (rem *Reminders) Post(w rest.ResponseWriter, r *rest.Request) {
	globals.Stats.AddHit(r.RequestURI)

	reminder := Reminder{}
	if err := r.DecodeJsonPayload(&reminder); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(reminder.Guid) == 0 {
		u, err := uuid.NewV4()
		if err != nil {
			log.Printf("Error creating new UUID: %v\n", err)
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			reminder.Guid = u.String()
		}
	}

	if err := rem.s.Save(reminder); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&reminder)
}

func (rem *Reminders) update(reminder Reminder, w rest.ResponseWriter, r *rest.Request) error {
	updated := Reminder{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	reminder.Message = updated.Message

	err := rem.s.Save(reminder)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

// Update (REST Put) a Reminder using id as the key.
func (rem *Reminders) Put(w rest.ResponseWriter, r *rest.Request) {
	globals.Stats.AddHit(r.RequestURI)

	id, err := idFromString(r.PathParam("id"))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	reminder, err := rem.s.GetId(id)
	if err != nil {
		rest.NotFound(w, r)
		return
	}

	if err = rem.update(reminder, w, r); err != nil {
		return
	}

	w.WriteJson(&reminder)
}

// Update (REST Put) a Reminder using guid as the key.
func (rem *Reminders) PutGuid(w rest.ResponseWriter, r *rest.Request) {
	guid := r.PathParam("guid")
	reminder, err := rem.s.GetGuid(guid)
	if reminder.Guid == "" {
		rest.NotFound(w, r)
		return
	}
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = rem.update(reminder, w, r); err != nil {
		return
	}

	w.WriteJson(&reminder)
}

// Handle REST Delete request, which presumes Id as the identifying key.
func (rem *Reminders) Delete(w rest.ResponseWriter, r *rest.Request) {
	globals.Stats.AddHit(r.RequestURI)

	id, err := idFromString(r.PathParam("id"))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reminder, err := rem.s.DeleteId(id)

	// If no record was found, Guid will remain empty
	if reminder.Guid == "" {
		rest.NotFound(w, r)
		return
	}

	// A record was found, but some error ocurred.
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handle REST Delete request using Guid as the identifying key.
func (rem *Reminders) DeleteGuid(w rest.ResponseWriter, r *rest.Request) {
	guid := r.PathParam("guid")

	reminder, err := rem.s.DeleteGuid(guid)

	// If no record was found, Guid will remain empty
	if reminder.Guid == "" {
		rest.NotFound(w, r)
		return
	}

	// A record was found, but some error ocurred.
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
