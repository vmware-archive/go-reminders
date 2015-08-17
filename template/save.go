// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package template

import (
	"github.com/tdhite/go-reminders/reminders"
	html_template "html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Saves a Reminder object given a POST request with the buid and message set
// in the form data.
func (t *Template) SaveHandler(w http.ResponseWriter, r *http.Request) {
	var data reminders.Reminder
	guid := r.FormValue("guid")
	message := r.FormValue("message")
	if message == "" {
		log.Panicf("message: \"%s\"\n", message)
	}

	if guid == "" {
		// new reminder (no guid at this time)
		data = reminders.Reminder{
			Message: message,
		}
		// send REST request to create
		t.createReminder(data)
	} else {
		// pump REST request to retrieve the object
		data = t.getReminder(guid)
		data.Message = message
		// send REST request to save
		t.saveReminder(data)
	}

	// run the index (home page) template to show all reminders again
	path := filepath.Join(t.ContentRoot, filepath.Dir(r.URL.Path), "index.html")
	page := filepath.Base(path)
	log.Printf("page, path: %s, %s\n", page, path)

	tmpl, err := html_template.New(page).ParseFiles(path)
	if err == nil {
		alldata := t.getAllReminders()
		if err := tmpl.ExecuteTemplate(w, page, alldata); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
