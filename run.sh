#!/usr/bin/env bash

sudo docker build . -t go-gin
sudo docker run -i -t -p 8080:8080 go-gin