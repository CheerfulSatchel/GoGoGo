#!/bin/bash

build:
	docker build -t pseudonym-api .

run:
	docker run --env-file ./.env.list -it -p 8080:8080 pseudonym-api
