#! /bin/bash

# Moved to root projects
cd ..

#Compile golang files to binary executable
ENV CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

