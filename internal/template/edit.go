// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
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

	"github.com/vmware/go-reminders/pkg/reminders"
)

// Opens a form for editing and subsequently saving a Reminder.
func (t *Template) EditHandler(w http.ResponseWriter, r *http.Request) {
	// setup the template path
	path := filepath.Join(t.ContentRoot, r.URL.Path) + ".html"
	page := filepath.Base(path)
	log.Printf("page, path: %s, %s\n", page, path)

	tmpl, err := html_template.New(page).ParseFiles(path)
	if err == nil {
		data := make([]reminders.Reminder, 1)
		guid := r.FormValue("guid")
		if guid == "" {
			// this is a new reminder request, so create one and
			data[0] = reminders.Reminder{}
		} else {
			data[0] = t.getReminder(guid)
		}
		td := RemindersData{
			Reminders: data,
			UrlRoot:   t.VHost,
		}
		if err := tmpl.ExecuteTemplate(w, page, td); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
