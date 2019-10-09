# go-reminders

## Overview
Sample microservice to manage reminders (tasks one seeks to remember).

## Try it out
To try out the microservice, it must be built and executed in a suitable
environment. Details are [further below](#build-and-run).

### Prerequisites
There are some basic requirements in order build and use the go-reminders
microservice:

1. A Kubernetes cluster sufficient to run go-reminders and its backing services (e.g., mysql);
2. A Docker registry and relevant credentials for pushing and pulling containers;
3. A backing service (generally mysql) running in the Kubernetes cluster;
4. A Pivotal Concourse setup in order to run the build and deploy processes;
5. [Helm](https://helm.sh) for managing go-reminders deployments, also used by the ci/cd pipelines.

#### Kubernetes Cluster
Running go-reminders generally targets a kubenernetes cluster. It runs in the
default namespace.

In addition, as with any Kubernetes cluster, credentials should be setup.
Secrets for pushing and pulling containers from the Docker registry must exist
before building the microservice. The way to add such credentials is as
follows:

```kubectl create secret docker-registry go-reminders-registry-creds --docker-server=<your-registry-server> --docker-username=<your-name> --docker-password=<your-pword> --docker-email=<your-email>```

As well, in one form or another, credentials for go-reminders to make use of a
backing store like MySQL is required. Those can be obtained in various ways,
such as described [below](#obtaining-backing-storage-secrets). Another way is
to use the [concourse credential
management](https://concourse-ci.org/creds.html) facility if you want a more
secure model. To do so, integrate your service of choice and set the
environment variables or command line switches in the Kubernetes deployment
for running go-reminders.

#### Backing Service
The microservice needs a backing service in which to store the reminders.
MySQL or internal facilities are common scenarios.

##### MySQL Backing Service
A working MySQL server can be used for the backing service. The microservice
requires rights to create and drop databases, as well as normal CRUD on
tables.

##### MemDB Backing Service
A memory only backing can be used for testing or demonstrative purposes. The
MemDB service is purely in memory and provides no persistence at all, nor
distribution of the reminders across go-reminder instances. It is intended
mainly for testing and development purposes. To use that, invoke go-reminders
similarly to the following:

```DBTYPE=mem ./cmd/go-reminders/go-reminders```

or use a similar environment setting in a Kubernetes deployment.

##### Obtaining Backing Storage Secrets
It is possible to obtain MySQL secrets from a VMware vRealize Orchestrator
backing service or an etd cluster.

###### Obtaining MySQL Creds from vRO
When utilizing the vRO capabilities, the service depends on the vRO workflow
to provide a valid database host address, admin login and login password where
the admin user has rights to create a database and tables. The vRO code may
need to change based on various workflows, but the gist is in the
[vro.go](pkg/reminders/vro.go) file.

###### Obtaining MySQL Creds from Etcd
When using etcd support, the expectation is that the etcd service is
deployed, and key/value pairs exist for:

* /host
* /port
* /user
* /passwd

All of these entries are required for go-reminders to successfully make
a connection to the MySQL server to support stateful operations.

### Build and Run
Building the code involves very few steps and there are some CI/CD options to
help out.

1. Get the code
2. Build the code
3. Optionally using CI/CD Pipelines

#### Get the Code
To get the code, get it similarly to the following:

    git clone https://github.com/vmware/go-reminders
    cd go-reminders
    git submodule init
    git submodule update --recursive

#### Build the Code
The project includes a Makefile for use in building the go-reminders
microservice. To build a Docker container, you must provide a container name
for pushing to a registry.

    export CONTAINER=myreponame/go-reminders
    make

That will build and push the container image assuming you have already logged
in to your registry with "docker login". If you are using
[concourse](https://concourse-ci.org) to perform your builds, you can issue the
makefile as follows to preclude immediately pushing the Docker image, instead
using the concourse resource to perform that task:

    export CONTAINER=myreponame/go-reminders:version-tag
    make cmd/go-reminders/go-reminders

The output of the concourse task would be the go-reminders executable.

To build for mac use the following.  Examples elsewhere in this document should
switch to referencing `go-reminders-darwin` instead of `go-reminders`

    make cmd/go-reminders/go-reminders-darwin

#### CI/CD Pipelines
The code includes various pipeines in the [build/ci](build/ci) directory. They
can be used assuming the appropriate tool (i.e., Jenkins, Concourse or the
like) is in place. Some alteration to the configuration may be necessary
depending on your setup.

##### Notes on Using Concourse
A set of pipelines and sample parameter files are provided in the
[build/ci/concourse](build/ci/concourse) directory. There is also a README
therein to explain using Concourse with this project.

##### Notes on Using Jenkins
A set of config.xml files for jobs and the Gerrit trigger plugin exist
in the [jenkins](build/ci/jenkins) directory. The config files can be used as templates
to setup a flow that, from a merge, for example, of the sources triggers a
full build, push of the go-reminders container to Docker hub, generate a HEAT
template to run it on an OpenStack instance, and thereafter kick off the HEAT
stack.

#### Run
The microservice will run standalone, in a container, or more often via
Kubernetes deployments. Whether by command line for testing or via Kubernetes
deployment, there are a number of options available to configure its
operations. All options are available via the command line or environment,
with the environment taking higher priority to lean a bit more towards
[12 factor](https://12factor.net). All command line switches are looked up in
the environment with the long form switch name in all capital letters.

##### Standalone
For instance To invoke go-reminders with a mysql database at
mysql.corp.local:3306, with user credentials as root/rootpasswd, issue the
following command:

    HOST=mysql.corp.local:3306 USER=root PASSWD=rootpasswd DBTYPE=mysql cmd/go-reminders/go-reminders

Alternatively:

    cmd/go-reminders/go-reminders --host=mysql.corp.local:3306 --user=root --passwd=rootpasswd --dbtype=mysql

To run go-reminders stand-alone, execute it as follows with appropriate
environment variables or command line switches:

    cmd/go-reminders/go-reminders ...

To get help on the available options, execute it as follows:

    cmd/go-reminders/go-reminders --help

##### In a Container
Once pushed to Docker, you can run the microservice similarly to:

	docker run -p 8080:8080 -E DBTYPE=mem myregistryrepo/go-reminders

##### In Kubernetes
In a Kubernetes environment, the command is part of the deployment manifest,
for example:

    ...
    command:
        - "/go-reminders"
    env:
        - name: DBTYPE
          value: "mysql"
        - name: HOST
          value: "mysql.corp.local"
        - name: USER
          value: "root"
        - name: PASSWD
          value: "rootpasswd"
    ...

###### Sample Kubernetes Deployment Using Kustomize
Sample Kubernetes manifests based on the use of
[Kustomize](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization/)
are provided in the [deployment/kubernetes](deployments/kubernetes) directory.
Prior to running kustomize, you will want to first create the go-reminders
Kubernetes secrets if you are using DBTYPE=mysql.

```
apiVersion: v1
kind: Secret
metadata:
  name: go-reminders-secrets
type: Opaque
data:
  admin: {{base64 encoded admin user name}}
  password: {{base64 encoded admin password}}
  host: {{base64 encoded database IP address or DNS name}}
  port: {{base64 encoded database port}}
```

You will also want to set the container repository and version tag. For an
example of doing that, take a look at the Concourse [task that replaces the
container
image](https://github.com/vmware/go-reminders/blob/master/build/ci/concourse/scripts/deployment-files.sh#L32-L43).

For  example, to Kustomize and apply the manifests, you might issue commands
similar to the following to start go-reminders in a 'dev' environment:

```
cd deployments/kubernetes
cat >kubernetes/base/kustomization.yaml <<EOF
  resources:
  - service.yml
  - deployment.yml
  images:
  - name: docker-registry-repo
    newTag: 1.0.12
    newName: concourse.corp.local/go-reminders
EOF
kubectl kustomize overlays/dev | kubectl apply -f -
```

###### Sample Kubernetes Deployment Using Helm
The facility generally used to deploy and manage go-reminders is
[Helm](https://helm.sh). If you are using the concourse pipelines to build and
deploy, helm will be used automatically.

If you want to deploy manually using helm, switch to the
[deployments/helm](deployment/helm] directory. There you will find the key
files to fill out -- either of values-minikube.yml or values-pks.yml. Copy one
of those files values.yml

    cp values-minikube.yml values.yml

and edit the values.yml file to match your deployment requirements. The file
values.yml is ignored by git, so will not be part of subequent commits unless
forced (i.e., git will not track values.yml).

To deploy using helm, issue a command such as:

    cd .../deployments/helm
    helm install --name go-reminders .

and that should create the deployment, which you can check with

    helm status go-reminders

Helm is a powerful tool for Kubernetes deployment management and a good read
of [its documentation](https://helm.sh/docs/) is in order for those not yet
familiar.

## Documentation
The microservice provides two mechanisms for interacting with the service:

1. an API for creating, modifying and deleting reminders and checking stats.
2. a web interface for visually doing the same.

### The API
Upon running the service, e.g.:

- etcd:
    docker run -d -p 8080:8080 go-reminders /go-reminders -cfgsrc etcd_host:2379

- vRO:
    docker run -d -p 8080:8080 go-reminders /go-reminders -cfgtype vro -cfgsrc 172.16.78.227

and assuming that provided you a Docker generated container address as
172.17.0.1, the REST API exists at http://172.17.0.1:8080/api/reminders and paths further thereafter pursuant to the pattern:

- GET /api/reminders:
Returns all reminders currently in the database.

- POST api/reminders:
Given a JSON body similar to:

    {
      "message": "message text"
    }

creates a new reminder.

- GET /api/reminders/:guid
- GET /api/reminders/byid/:id

Gets the JSON representing the Reminder with Guid ":guid" or ":id" as
appropriate.  The :id value is the database record id.

- PUT /api/reminders/:guid
- PUT /api/reminders/byid/:id

Given a JSON body similar to:

    {
      "message": "message text"
    }

Updates the Reminder with Guid ":guid" or ":id" as appropriate.
The :id value is the database record id.

- DELETE /api/reminders/:guid
- DELETE /api/reminders/byid/:id

Deletes the Reminder with Guid ":guid" or ":id" as appropriate.
The :id value is the database record id.

#### The HTML Interface
To reach the HTML interface (given the same sample as above), browse to:
http://172.17.0.1/html/tmpl/index and the bulk  of the HTML paths are
available from that page or others as appropriate given traversal of the 'site.'

Another HTML area is the /stats/hits, which provides a view of hit counts on
the various URLs involved in the service (API and HTML).

## Releases & Major Branches

## Contributing

The go-reminders project team welcomes contributions from the community. Before you start working with go-reminders, please read our [Developer Certificate of Origin](https://cla.vmware.com/dco). All contributions to this repository must be signed as described on that page. Your signature certifies that you wrote the patch or have the right to pass it on as an open-source patch. For more detailed information, refer to [CONTRIBUTING.md](CONTRIBUTING.md).

## License
Copyright: Copyright 2015-2019 VMware, Inc. All Rights Reserved.

SPDX-License-Identifier: [https://spdx.org/licenses/MIT.html](https://spdx.org/licenses/MIT.html)
