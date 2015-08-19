// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package template

import (
	"bytes"
	"encoding/json"
	"github.com/tdhite/go-reminders/reminders"
	"github.com/tdhite/go-reminders/stats"
	"io/ioutil"
	"log"
	"net/http"
)

type Template struct {
	ContentRoot string
	APIHost     string
}

// Return a new Template object initialized -- convenience function.
func New(contentRoot string, apiHost string) Template {
	return Template{
		ContentRoot: contentRoot,
		APIHost:     apiHost,
	}
}

func init() {
	log.Println("Initialized Template.")
}

func (t *Template) generateAPIUrl(path string) string {
	return "http://" + t.APIHost + path
}

// Retrieve a Reminder from storage via REST call.
func (t *Template) getReminder(guid string) reminders.Reminder {
	url := t.generateAPIUrl("/api/reminders/" + guid)
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
	url := t.generateAPIUrl("/api/reminders")
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
	url := t.generateAPIUrl("/api/reminders/" + guid)
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

	url := t.generateAPIUrl("/api/reminders/" + r.Guid)
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

	url := t.generateAPIUrl("/api/reminders")
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

// Retrieve a Reminder from storage via REST call.
func (t *Template) getStatsHits() map[string]int {
	url := t.generateAPIUrl("/stats/hits")
	log.Println("url: " + url)

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	data, err := stats.HitsFromJson(body)
	perror(err)

	return data
}
