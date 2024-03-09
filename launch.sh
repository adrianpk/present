#!/bin/bash

GO111MODULE=off go get -u golang.org/x/tools/present

git clone git@github.com:adrianpk/present.git slides

cd slides

present
