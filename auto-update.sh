#!/bin/sh

set -e

git submodule update --remote

go run ./lib
