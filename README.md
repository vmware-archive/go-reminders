# go-reminders
Sample microservice to manage reminders (tasks one seeks to remember).

# Getting Started
See the Dependencies section in this document for initial setup.

## Building the code
Until such time as we release this repository to the public, clone it in
your GOPATH at src/github.com/tdhite/go-reminders, then build it:

For example:

    export GOPATH=${HOME}/go
    mkdir -p ${HOME}/go/src/github.com/tdhite
    cd ${HOME}/go/src/github.com/tdhite
    git clone http://gerrit.cloudbuilders.vmware.local/go-reminders
    cd go-reminders
    make

### Using Jenkins
A set of config.xml files for jobs and the Gerrit trigger plugin exist
in the jenkins directory. The config files can be used as templates to
setup a flow that, from a merge, for example, of the sources triggers
a full build, push of the go-reminders container to Docker hub, generate
a HEAT template to run it on an OpenStack instance, and thereafter kick
off the HEAT stack.

## The API
Upon running the service, e.g.:

    docker run -d -p 8080:8080 go-reminders /go-reminders -a 172.16.78.227

and assuming that provided you a Docker generated container address as
172.17.0.1, the REST API exists at http://172.17.0.1:8080/api/reminders and paths further thereuafter pursuant to the pattern:

- GET /api/reminders:
Returns all reminders currently in the database.

- POST api/reminders:
Given a JSON body similar to:

```
    {
      "message": "message text"
    }
```

creates a new reminder.

- GET /api/reminders/:guid
- GET /api/reminders/byid/:id

Gets the JSON representing the Reminder with Guid ":guid" or ":id" as
appropriate.  The :id value is the database record id.

- PUT /api/reminders/:guid
- PUT /api/reminders/byid/:id

Given a JSON body similar to:

```
    {
      "message": "message text"
    }
```

Updates the Reminder with Guid ":guid" or ":id" as appropriate.
The :id value is the database record id.

- DELETE /api/reminders/:guid
- DELETE /api/reminders/byid/:id

Deletes the Reminder with Guid ":guid" or ":id" as appropriate.
The :id value is the database record id.


## The HTML Interface
To reach the HTML interface (given the same sample as above), browse to:
http://172.17.0.1/html/tmpl/index and the bulk  of the HTML paths are
available from that page or others as appropriate given traversal of the 'site.'

Another HTML area is the /stats/hits, which provides a view of hit counts on
the various URLs involved in the service (API and HTML).

# Dependencies
This service requires a valid Go language environment and gnu make.

To install go, see https://golang.org/doc/install

A working MySQL server must also be available. The microservice requires
rights to create and drop databases, as well as normal CRUD on tables.

When utilizing the vRO capabilities, the service depends on the vRO workflow
to provide a valid database host, admin login  and login password where the
admin user has rights to create a database and tables. The vRO code may need
to change based on various workflows, but the gist is in the vro directory.

# License and Author
Copyright: Copyright (c) 2015 VMware, Inc. All Rights Reserved

Author: Tom Hite, VMware, Inc.

License: MIT

For details of the license, see the LICENSE file.

