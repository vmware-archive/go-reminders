# Concourse pipelines for go-reminders

## Prerequisites
In order to use these pipelines, you must have:

- a running [Concourse](https://concourse-ci.org) setup
- a kubernetes cluster with relevant credential
- a git repo and the related SSH key for housing a file for [semantic version management](https://concoursetutorial.com/miscellaneous/versions-and-buildnumbers/#semver-semantic-versioning)
- a valid git repo for the go-reminders packate (e.g., a github fork or similar) and related the private SSH key.

The key is to fill out a params.yml file similar to by copying a template from
[examples](examples). There is a sample for minikube and pks based pipelines.

## Try and run
Generally, you will 'fly login' and execute `00_set-pipeline.sh`, which will
create/update the pipeline in concourse. 

### Minikube
Start a minikube cluster that has access to a concourse CI engine. A sample
script for this exist in [scripts/minikube.sh](../../../scripts/minikube.sh).

### Credential Pipeline Variables

There are a number of variables that params.yml needs to specify, and as with
they should be base64 encoded. These are:

- k8s-cluster-url: something of the form https://192.168.64.55:8443
- k8s-cluster-ca:  base64 encoded certificate authority for the target cluster
- k8s-admin-cert:  base64 encoded user cert for the target cluster
- k8s-admin-key:   base64 encoded user key for the target cluster
- k8s-admin-token: base64 admin token, if any. If none, use "MINIKUBE"

In order to help out a bit with those, a
[script](examples/append-creds-to-params.sh) exists in the
[examples](examples) directory that will attempt to form the appropriate
values from your ~/.kube/config and append the variables with their values to
teh file "params.yml". It is likely you will need to modify the script a bit
to set its internal variables, which are documented within.

Once up and running, copy and edit the
[examples/params-minikube.yml](examples/params-minikube.yml) file to match
your to params.yml. The latter will be ignored by git, so no worries on
subsequent commits of your changes. Then run:

    ./00_set-pipeline-minikube.sh

### PKS
    TODO
