#!/bin/sh

LAYER=$1
ENV=$2

cd tffile/environment/$LAYER
terraform init -input=false -no-color
terraform destroy -input=false -no-color -auto-approve
