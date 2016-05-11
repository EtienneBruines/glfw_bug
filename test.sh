#!/usr/bin/env bash

Xvfb :10 -ac &
export DISPLAY=:10
go run main.go