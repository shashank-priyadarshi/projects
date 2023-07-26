#!/bin/bash

# Check if at least one argument is received
if [ $# -lt 1 ]; then
  echo "Error: GitHub token is required."
  exit 1
fi

# Set the value of GITHUB_TOKEN
GITHUB_TOKEN=$1
DEV_SETUP=0
BRANCH_NAME=dev

# Check if two arguments are received
if [ $# -eq 2 ]; then
  # Set the value of DEV_SETUP
  DEV_SETUP=$2
  BRANCH_NAME=main
  echo "Running setup in production mode: $DEV_SETUP, $BRANCH_NAME"
else
  echo "Running setup in development mode: $DEV_SETUP , $BRANCH_NAME"
fi

# Check Go version
go_version=$(go version | awk -F '[ .]' '{print $3 $4}' | sed 's/go//')
if [[ $go_version -lt 120 ]]; then
  echo "Go version must be greater than or equal to 1.20. Current version: $(go version)"
  exit 1
fi

cd upgraded-disco && go mod tidy && go get -u && git commit -am "updated dep packages" && export GH=$GITHUB_TOKEN && export SETUP=$DEV_SETUP && gnome-terminal -- bash -c "air; if [ \$? -ne 0 ]; then exit; fi"
cd ../portfolio-core-ui && npm i --f --legacy-peer-deps && gnome-terminal -- bash -c "npm run start; if [ \$? -ne 0 ]; then exit; fi"
