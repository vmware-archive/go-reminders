# go-reminders minimal container definition

# Start from scratch
FROM scratch
MAINTAINER Tom Hite <thite@vmware.com>

# Add the microservice
ADD go-reminders /go-reminders

# Add the content (html and templates)
ADD html /html
