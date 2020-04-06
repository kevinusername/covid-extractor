#!/bin/sh

set -e

git submodule update --remote

go run ./lib

git commit -am "Automated Update $(date --rfc-3339=date)"
