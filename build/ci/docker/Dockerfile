# go-reminders minimal container definition
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

# Start from scratch
FROM golang:1.12.9
MAINTAINER Tom Hite <thite@vmware.com>

# Add the microservice
ADD kubectl /usr/local/bin/kubectl
ADD linux-amd64 /usr/local/bin/linux-amd64
RUN ln -s /usr/local/bin/linux-amd/bin/helm /usr/local/bin/helm
