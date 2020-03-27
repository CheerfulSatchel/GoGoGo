#!/bin/bash

build:
	docker build -t pseudonym_models_service .

run:
	docker run --env-file ./.env.list -it -p 8081:8081 pseudonym_models_service
