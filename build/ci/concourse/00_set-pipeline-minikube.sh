#!/bin/bash
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

PARAMS="params.yml"
PIPELINE="pipeline-minikube.yml"

fly -t k8s-cluster set-pipeline --pipeline=go-reminders --load-vars-from="${PARAMS}" --config="${PIPELINE}"

#if [ $? -eq 0 ]; then
#	fly -t k8s-cluster unpause-pipeline --pipeline go-reminders
#fi
