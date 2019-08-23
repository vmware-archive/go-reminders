# Concourse pipelines for go-reminders

## Prerequisites
In order to use these pipelines, you must have:

- a running [Concourse](https://concourse-ci.org) setup
- a kubernetes cluster with relevant credentials (k8s secrets) setup
- a git repo and the related SSH key for housing a file for [semantic version management](https://concoursetutorial.com/miscellaneous/versions-and-buildnumbers/#semver-semantic-versioning)
- a valid git repo for the go-reminders project (e.g., a github fork or similar) and related private SSH key.

The key is to fill out a params.yml file similar by copying a template from
[examples](examples). There are samples for minikube and PKS based pipelines.

## Try and run
Generally, you will 'fly login' and execute `00_set-pipeline.sh`, which will
create/update the pipeline in concourse. 

### Minikube
Start a minikube cluster that has access to a concourse CI engine. A sample
script for this exists in [scripts/minikube.sh](../../../scripts/minikube.sh).

### Credential Pipeline Variables

There are a number of pipeline variables that params.yml needs in order to
complete pipeline runs. These include:

- docker-registry-repo: the docker repo/container to use for push / pull
- docker-registry-user: the docker registry login user
- docker-registry-passwd: the docker registry login passwd
- docker-registry-email: the docker registry login e-mail
- helm_ver: the version of helm to install and use

In addition, there are a number of variables that params.yml needs to specify
that must be base64 encoded. These are:

- k8s-cluster-url: something of the form https://192.168.64.55:8443
- k8s-cluster-ca:  base64 encoded certificate authority for the target cluster
- k8s-admin-cert:  base64 encoded user cert for the target cluster
- k8s-admin-key:   base64 encoded user key for the target cluster
- k8s-admin-token: base64 admin token, if any. If none, use "MINIKUBE"

In order to help out a bit with the base64 encoding, a
[script](examples/append-creds-to-params.sh) exists in the
[examples](examples) directory that will attempt to form the appropriate
values from your ~/.kube/config and append the variables with their values to
the file "params.yml". The script may need modifications to set its internal
variables, which are documented directly within.

Once a kubernetes cluster is up and running, copy and edit the
[examples/params-minikube.yml](examples/params-minikube.yml) file
params.yml, for example:

    cp examples/params-minikube.yml params.yml
    vi params.yml

That should be ignored by git to help prevent potential
commits of your changes. Then run:

    ./00_set-pipeline.sh

### PKS
    TODO
