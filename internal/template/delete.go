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
)

// Deletes a Reminder from storage given a guid.
func (t *Template) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	t.stats.AddHit(r.RequestURI)

	guid := r.FormValue("guid")
	if guid == "" {
		log.Panicf("guid: \"%s\"\n", guid)
	}

	t.deleteReminder(guid)

	path := filepath.Join(t.ContentRoot, filepath.Dir(r.URL.Path), "index.html")
	page := filepath.Base(path)
	log.Printf("page, path: %s, %s\n", page, path)

	tmpl, err := html_template.New(page).ParseFiles(path)
	if err == nil {
		td := RemindersData{
			Reminders: t.getAllReminders(),
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
