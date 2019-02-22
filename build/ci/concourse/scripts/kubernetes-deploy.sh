#!/bin/bash
# go-reminders kubernetes-deploy.sh
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

set -x

# Save current directory
TOP="$(pwd)"

# install kubectl
LATEST="$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)"
if [ $? -ne 0 ]; then
    echo "Failed to obtain latest kubectl release. Aborting!"
    exit 1
fi

curl -LO https://storage.googleapis.com/kubernetes-release/release/${LATEST}/bin/linux/amd64/kubectl
if [ $? -ne 0 ]; then
    echo "Failed to download kubectl. Aborting!"
    exit 1
fi
chmod +x kubectl

tag=$(cat version/version)

# validate parameters
ret=0
if [ -z "${cluster_ca}" ]; then
    echo "ERROR: cluster_ca not supplied. Aborting!"
    ret=1
fi
if [ -z "${cluster_url}" ]; then
    echo "ERROR: cluster_url not supplied. Aborting!"
    ret=1
fi
if [ -z "${namespace}" ]; then
    echo "ERROR: namespace not supplied. Aborting!"
    ret=1
fi
if [ -z "${resource_type}" ]; then
    echo "ERROR: resource_type not supplied. Aborting!"
    ret=1
fi
if [ -z "${resource_name}" ]; then
    echo "ERROR: resource_name not supplied. Aborting!"
    ret=1
fi
if [ -z "${container}" ]; then
    echo "ERROR: container not supplied. Aborting!"
    ret=1
fi
if [ -z "${admin_key}" ]; then
    echo "ERROR: admin_key not supplied. Aborting!"
    ret=1
fi
if [ -z "${admin_cert}" ]; then
    echo "ERROR: admin_cert not supplied. Aborting!"
    ret=1
fi
if [ -z "${admin_token}" ]; then
    echo "ERROR: admin_token not supplied. Aborting!"
    ret=1
fi
if [ -z "${tag}" ]; then
    echo "ERROR: tag (version) not supplied. Aborting!"
    ret=1
fi
if [ $ret -ne 0 ]; then
    exit $ret
fi

echo "build credentials"

echo "$cluster_ca" | base64 -d > ca.pem

# if using a bearer token or using minikube, admin keys are useless
if [ -z "$admin_token" -o "${admin_token}" == "MINIKUBE" ]; then
    echo "$admin_key" | base64 -d > key.pem
    echo "$admin_cert" | base64 -d > cert.pem
fi

# list the directory now for debugging purposes
ls -lat

# build kubectl command line
KUBECTL="./kubectl --server=${cluster_url} --namespace=$namespace --certificate-authority=ca.pem"
if [ -z "$admin_token" -o "${admin_token}" == "MINIKUBE" ]; then
    KUBECTL="${KUBECTL} --client-key=key.pem --client-certificate=cert.pem"
else
    KUBECTL="${KUBECTL} --token=${admin_token}"
fi

# run the correct commands
${KUBECTL} get ${resource_type}/${resource_name} >/dev/null 2>&1
if [ $? -eq 0 ]; then
    # Note: assume resource_name and the container name equate.
    $KUBECTL set image ${resource_type}/${resource_name} ${resource_name}=${container}:${tag}
else
    # No deployment yet, start it.
    $KUBECTL create -f kubernetes/deployment.yml
fi

${KUBECTL} get service/${resource_name} >/dev/null 2>&1
if [ $? -ne 0 ]; then
    # No service yet, start it.
    $KUBECTL create -f kubernetes/service.yml
fi
