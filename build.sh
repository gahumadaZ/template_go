#!/bin/bash

go -C cmd/my_app build -o ../../my_project
env GOOS=linux GOARCH=amd64 go -C cmd/my_app build -o ../../my_project_x64
