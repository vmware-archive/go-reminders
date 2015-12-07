// Copyright 2015 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package template

import (
	html_template "html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/vmwaresamples/go-reminders/pkg/reminders"
)

// Generate the main (home) page of the site.
func (t *Template) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var data []reminders.Reminder

	guid := r.URL.Query().Get("guid")
	if guid == "" {
		log.Printf("guid: \"%s\"\n", guid)
		data = t.getAllReminders()
	} else {
		r := t.getReminder(guid)
		data = make([]reminders.Reminder, 1)
		data[1] = r
	}

	path := filepath.Join(t.ContentRoot, r.URL.Path) + ".html"
	page := filepath.Base(path)
	log.Printf("page, path: %s, %s\n", page, path)

	tmpl, err := html_template.New(page).ParseFiles(path)
	if err == nil {
		if err := tmpl.ExecuteTemplate(w, page, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
