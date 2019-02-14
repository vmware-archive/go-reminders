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

Once up and running, copy and edit the
[examples/params-minikube.yml](examples/params-minikube.yml) file to match
your to params.yml. The latter will be ignored by git, so no worries on
subsequent commits of your changes. Then run:

    ./00_set-pipeline-minikube.sh

### PKS
    TODO

