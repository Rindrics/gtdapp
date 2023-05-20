#!/bin/sh

LAYER=$1

cd tffile/environment/$LAYER
terraform init -input=false -no-color
terraform plan -input=false -no-color
