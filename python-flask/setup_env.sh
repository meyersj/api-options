#!/bin/bash

# get directory of script
DIR="$(cd "$(dirname "$0")" && pwd)"

# create python virtual environment and install dependencies
rm -rf $DIR/env
virtualenv $DIR/env
$DIR/env/bin/pip install -r $DIR/requirements.txt
