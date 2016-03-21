#!/bin/bash

set -e

FILES=(gitcookies inago-integration-test.pem)
SECRET_FILE=secrets.tar

tar cvf $SECRET_FILE ${FILES[*]}
travis encrypt-file $SECRET_FILE
