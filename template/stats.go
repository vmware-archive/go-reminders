// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package template

import (
	html_template "html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Generate the main (home) page of the site.
func (t *Template) StatsHitsHandler(w http.ResponseWriter, r *http.Request) {
	stats := t.getStatsHits()

	path := filepath.Join(t.ContentRoot, r.URL.Path) + ".html"
	page := filepath.Base(path)
	log.Printf("page, path: %s, %s\n", page, path)

	tmpl, err := html_template.New(page).ParseFiles(path)
	if err == nil {
		if err := tmpl.ExecuteTemplate(w, page, stats); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
