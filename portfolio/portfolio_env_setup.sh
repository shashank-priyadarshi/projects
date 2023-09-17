#!/bin/bash

# Check if at least one argument is received
if [ $# -lt 1 ]; then
  echo "Error: GitHub token is required."
  exit 1
fi

# Set the value of GITHUB_TOKEN
GITHUB_TOKEN=$1
DEV_SETUP=0

# Check if two arguments are received
if [ $# -eq 2 ]; then
  # Set the value of DEV_SETUP
  DEV_SETUP=$2
fi

echo "Running setup in $(if [[ $DEV_SETUP -eq 0 ]]; then echo development; else echo production; fi) mode"

# Check Go version
#go_version=$(go version | awk -F '[ .]' '{print $3 $4}' | sed 's/go//')
#if [[ $go_version -lt 120 ]]; then
#  echo "Go version must be greater than or equal to 1.20. Current version: $(go version)"
#  exit 1
#fi

echo "export GH=$GITHUB_TOKEN" > ./tmp-vars.sh
echo "export SETUP=$DEV_SETUP" >> ./tmp-vars.sh
chmod +x tmp-vars.sh
#gnome-terminal -- bash -c "ls && source ./tmp-vars.sh && chmod +x tmp-vars.sh && echo tmp-vars.sh && echo $GH && echo $SETUP  && rm ./tmp-vars.sh && docker-compose up -d; if [ \$? -ne 0 ]; then exit; fi"
#gnome-terminal -- bash -c "echo $DEV_SETUP; echo $GITHUB_TOKEN; source ./tmp-vars.sh; ls; GH=$GITHUB_TOKEN SETUP=$DEV_SETUP docker-compose up -d && rm tmp-vars.sh; if [ \$? -ne 0 ]; then exit; fi"
source ./tmp-vars.sh
docker-compose down && docker-compose rm go-dev --remove-orphans && docker rm portfolio_angular-dev_1 --remove-orphans
GH=$GITHUB_TOKEN SETUP=$DEV_SETUP docker-compose up -d --build
rm tmp-vars.sh
#cd upgraded-disco && go mod tidy && go get -u && git commit -am "updated dep packages" && export GH=$GITHUB_TOKEN && export SETUP=$DEV_SETUP && gnome-terminal -- bash -c "air; if [ \$? -ne 0 ]; then exit; fi"
#cd ../portfolio-core-ui && npm i --f --legacy-peer-deps && gnome-terminal -- bash -c "npm run start; if [ \$? -ne 0 ]; then exit; fi" && gnome-terminal -- bash -c "google-chrome http://localhost:4200; if [ \$? -ne 0 ]; then exit; fi"
