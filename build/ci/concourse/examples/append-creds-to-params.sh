#!/bin/bash
# append-creds-to-params.sh
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#
# This script may help adding credentials to a params.yml file modified for
# use in setting concourse pipelines for the go-reminders project.
#
# The script makes attempts to pull cred info out of ~/.kube/config and set
# appropriate values (appends to) the file params.yml. It also sets a market
# in place to let the caller know what to delete from params.yml if it is
# executed again.
#
# What that tells all is this is by no means an idemptent process, rather it's
# just enough script to try to help setting variables for pipeline setup, but
# your mileage may vary.
#
DEBUG=1

# Set this to anything hat follows variables like "certificate-authority"
# For instance many time the kube config will use "certificate-authority-data".
# Minikube, just leaves no such postfix.
#CDATAPOSTFIX="-data"
CDATAPOSTFIX=""

# Set this to the name of the user you want to perform k8s deployments.
export K8SUSER="minikube"

# Set this to the cluster into which you want to deploy.
export K8SCLUSTER="minikube"

if [ -z "$(which jq)" ]; then
    echo "ERROR: jq must be available on the PATH, but is not."
    exit 1
fi

if [ -z "$(which base64)" ]; then
    echo "ERROR: base64 must be available on the PATH, but is not."
    exit 2
fi

case "${OS}" in
    darwin*)
        BASE64DECODE="base64 -D"
        ;;
    *)
        BASE64DECODE="base64 -d"
        ;;
esac

echo get cert data...
CADATA=$(cat ~/.kube/config | python -c 'import sys, yaml, json; json.dump(yaml.load(sys.stdin), sys.stdout, indent=4)' | jq ".clusters[] | select(.name==\"${K8SCLUSTER}\")" | jq -r ".cluster[\"certificate-authority${CDATAPOSTFIX}\"]")
CCDATA=$(cat ~/.kube/config | python -c 'import sys, yaml, json; json.dump(yaml.load(sys.stdin), sys.stdout, indent=4)' | jq ".users[] | select(.name==\"${K8SUSER}\")" | jq -r ".user[\"client-certificate${CDATAPOSTFIX}\"]")
CKDATA=$(cat ~/.kube/config | python -c 'import sys, yaml, json; json.dump(yaml.load(sys.stdin), sys.stdout, indent=4)' | jq ".users[] | select(.name==\"${K8SUSER}\")" | jq -r ".user[\"client-key${CDATAPOSTFIX}\"]")

echo get token...
if [ "${K8SCLUSTER}" == "minikube" ]; then
    TOKEN=MINIKUBE
else
    TOKEN=$(cat ~/.kube/config | python -c 'import sys, yaml, json; json.dump(yaml.load(sys.stdin), sys.stdout, indent=4)' | jq ".users[] | select(.name==\"$K8SUSER\")" | jq -r '.user["token"]')
fi

echo get k8s host...
K8SHOST=$(kubectl config view -o jsonpath="{.clusters[?(@.name == \"${K8SCLUSTER}\")].cluster.server}")

# Append all the things to params.yml
echo "#### Added by ${0}. Delete this line and below to rerun the script." >>params.yml
echo "####" >>params.yml
echo "appending k8s-cluster-url..."
echo -n "k8s-cluster-url: " >>params.yml
echo "${K8SHOST}" >>params.yml
echo

if [ "$K8SCLUSTER" == "minikube" ]; then
    echo "appending k8s-cluster-ca..."
    echo -n "k8s-cluster-ca: " >>params.yml
    cat "${CADATA}" | base64 >>params.yml
    echo >>params.yml

    echo "appending k8s-admin-cert..."
    echo -n "k8s-admin-cert: " >>params.yml
    cat ${CCDATA} | base64 >>params.yml
    echo >>params.yml

    echo "appending k8s-admin-key..."
    echo -n "k8s-admin-key: " >>params.yml
    cat ${CKDATA} | base64 >>params.yml
    echo >>params.yml

    echo "appending k8s-admin-token."
    echo -n "k8s-admin-token: " >>params.yml
    echo ${TOKEN} >>params.yml
    echo >>params.yml
else
    echo "appending k8s-cluster-ca..."
    echo -n "k8s-cluster-ca: " >>params.yml
    echo -n ${CADATA} >>params.yml
    echo >>params.yml

    echo "appending k8s-admin-cert..."
    echo -n "k8s-admin-cert: " >>params.yml
    echo -n ${CCDATA} >>params.yml
    echo >>params.yml

    echo "appending k8s-admin-key..."
    echo -n "k8s-admin-key: " >>params.yml
    echo -n ${CKDATA} >>params.yml
    echo >>params.yml

    echo "appending k8s-admin-token."
    echo -n "k8s-admin-token: " >>params.yml
    echo -n ${TOKEN} >>params.yml
    echo >>params.yml
fi
