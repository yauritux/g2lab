#!/usr/bin/env bash
pkill main
source env
go run app/main.go &

inotifywait -m -r -e close_write adapters app domain engine providers | while read line
do
  pkill main && go run app/main.go &
done