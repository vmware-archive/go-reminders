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

// Opens a form for editing and subsequently saving a Reminder.
func (t *Template) EditHandler(w http.ResponseWriter, r *http.Request) {
	var data reminders.Reminder
	guid := r.FormValue("guid")
	if guid == "" {
		log.Panicf("guid: \"%s\"\n", guid)
	}

	path := filepath.Join(t.ContentRoot, r.URL.Path) + ".html"
	page := filepath.Base(path)
	log.Printf("page, path: %s, %s\n", page, path)

	tmpl, err := html_template.New(page).ParseFiles(path)
	if err == nil {
		data = t.getReminder(guid)
		if err := tmpl.ExecuteTemplate(w, page, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
