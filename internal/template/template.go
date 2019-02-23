// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package template

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/vmware/go-reminders/pkg/reminders"
	"github.com/vmware/go-reminders/pkg/stats"
)

const apiRoot = "/api/reminders"

type Template struct {
	ContentRoot string
	VHost       string
	APIBase     string
	stats       stats.Stats
}

type RemindersData struct {
	Reminders []reminders.Reminder
	UrlRoot   string
}

// Return a new Template object initialized -- convenience function.
func New(contentRoot string, vhost string, apiroot string,
	stats stats.Stats) Template {
	return Template{
		ContentRoot: contentRoot,
		APIBase:     apiroot,
		VHost:       vhost,
		stats:       stats,
	}
}

func init() {
	log.Println("Initialized Template.")
}

func (t *Template) generateAPIUrl(p string) string {
	u, err := url.Parse(t.APIBase)
	if err != nil {
		log.Printf("ERROR: failed to parse APIRoot %s!\n", t.APIBase)
		return p
	} else {
		log.Printf("generateAPIUrl parsed APIRoot as %s!\n", u.String())
	}
	u.Path = path.Join(u.Path, p)
	log.Printf("generateAPIUrl created final URL as %s!\n", u.String())
	return u.String()
}

// Retrieve a Reminder from storage via REST call.
func (t *Template) getReminder(guid string) reminders.Reminder {
	url := t.generateAPIUrl(apiRoot + "/" + guid)
	log.Println("url: " + url)

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	data, err := reminders.FromJson(body)
	perror(err)

	return data
}

// Retrieve all Reminders from storage via REST call.
func (t *Template) getAllReminders() []reminders.Reminder {
	url := t.generateAPIUrl(apiRoot)
	log.Println("url: " + url)

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	data, err := reminders.ArrayFromJson(body)
	perror(err)

	return data
}

// Delete the Reminder, to which guid refers, from storage via REST call.
func (t *Template) deleteReminder(guid string) {
	url := t.generateAPIUrl(apiRoot + "/" + guid)
	log.Println("url: " + url)

	req, err := http.NewRequest("DELETE", url, nil)
	perror(err)

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	_, err = ioutil.ReadAll(rsp.Body)
	perror(err)
}

// Save the Reminder, to which guid refers, to storage via REST call.
func (t *Template) saveReminder(r reminders.Reminder) {
	jsonData, err := json.Marshal(r)
	perror(err)

	url := t.generateAPIUrl(apiRoot + "/" + r.GUID)
	log.Println("url: " + url)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	perror(err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	_, err = ioutil.ReadAll(rsp.Body)
	perror(err)
}

// Save the Reminder, to which guid refers, to storage via REST call.
func (t *Template) createReminder(r reminders.Reminder) {
	jsonData, err := json.Marshal(r)
	perror(err)

	url := t.generateAPIUrl(apiRoot)
	log.Println("url: " + url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	perror(err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	_, err = ioutil.ReadAll(rsp.Body)
	perror(err)
}

// Retrieve stats via REST call.
func (t *Template) getStatsHits() map[string]int {
	url := t.generateAPIUrl("/stats/hits")
	log.Println("url: " + url)

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	data, err := stats.HitsFromJSON(body)
	perror(err)

	return data
}
