// Copyright (c) 2015 VMware
// Author: Tim Green (greent@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//

package etcd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/tdhite/go-reminders/db"
	"golang.org/x/net/context"
)

type Connection struct {
	kapi client.KeysAPI
}

func New(host string) Connection {
	var conn Connection
	h := fmt.Sprintf("http://%s", host)
	fmt.Println("Attempting to connect to etcd at ", h)
	cfg := client.Config{
		Endpoints: []string{h},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	conn.kapi = client.NewKeysAPI(c)
	return conn
}

func (conn *Connection) GetDBCreds(d *db.DB) {
	dbInfo := map[string]string{"/host": "", "/port": "", "/user": "", "/passwd": ""}
	for key, _ := range dbInfo {
		resp, err := conn.kapi.Get(context.Background(), key, nil)
		if err != nil {
			log.Fatal(err)
		}
		dbInfo[key] = resp.Node.Value
	}
	d.SetAddress(dbInfo["/host"])
	port, _ := strconv.Atoi(dbInfo["/port"])
	d.SetPort(port)
	d.SetAdmin(dbInfo["/user"])
	d.SetPasswd(dbInfo["/passwd"])
}

func (conn *Connection) SetDBCreds(d *db.DB) {
	port := strconv.Itoa(d.Port())

	conn.kapi.Set(context.Background(), "/host", d.Address(), nil)
	conn.kapi.Set(context.Background(), "/port", port, nil)
	conn.kapi.Set(context.Background(), "/user", d.Admin(), nil)
	conn.kapi.Set(context.Background(), "/passwd", d.Passwd(), nil)
}
