#!/bin/sh
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

if [ -f /doesnotexist ]; then
# Get Access to Fusion command-line utils
# export PATH=$PATH:"/Applications/VMware Fusion.app/Contents/Library"

echo "Starting minikube cluster with 8G memory, 100G disk..."
minikube start --memory 8000 --cpus 4 --disk-size 100g --dns-domain corp.local --vm-driver hyperkit

eval $(minikube docker-env)

# assure ingress is enabled
echo "Enabling ingress..."
minikube addons enable ingress

# initialize helm
echo "Initializing helm..."
helm init
sleep 5 # settling time

fi

# start a concourse engine
echo "Starting concourse..."
helm install --name reminders-concourse stable/concourse

# Notify the user how to use the local concourse
echo To use concourse locally, you will want to do the following:
echo
echo export CPOD=\$\(kubectl get pods --namespace default -l \"app=reminders-concourse-web\" -o jsonpath=\"\{.items\[0\].metadata.name\}\"\)
echo kubectl port-forward --namespace default \$CPOD 8080:8080
echo fly -t k8s-cluster login --team-name main --concourse-url http://127.0.0.1:8080 -u test -p test
echo
echo
#export CPOD=$(kubectl get pods --namespace default -l "app=reminders-concourse-web" -o jsonpath="{.items[0].metadata.name}")
#kubectl port-forward --namespace default $CPOD 8080:8080

