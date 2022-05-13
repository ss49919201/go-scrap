#!/bin/bash

go build -o ./examples/gpl/main ./examples/gpl/main.go

./examples/gpl/main

rm ./examples/gpl/main
