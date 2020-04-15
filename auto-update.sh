#!/bin/sh

set -e

git submodule update --remote

go run ./lib

npx prettier --write ./out

git commit -am "Automated Update $(date --rfc-3339=date)"
