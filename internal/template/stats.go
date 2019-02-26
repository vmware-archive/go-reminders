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

type StatsData struct {
	Stats   map[string]int
	UrlRoot string
}

// Generate the main (home) page of the site.
func (t *Template) StatsHitsHandler(w http.ResponseWriter, r *http.Request) {
	stats := t.getStatsHits()

	path := filepath.Join(t.ContentRoot, r.URL.Path) + ".html"
	page := filepath.Base(path)
	log.Printf("page, path: %s, %s\n", page, path)

	tmpl, err := html_template.New(page).ParseFiles(path)
	if err == nil {
		sd := StatsData{
			Stats:   stats,
			UrlRoot: t.VHost,
		}
		if err := tmpl.ExecuteTemplate(w, page, sd); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
