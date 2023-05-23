#!/bin/bash

echo "Running mongo container with volume mount--------------------"
docker-compose -f docker-compose-mongoDB.yaml up -d

echo "Checking running containers--------------------"
docker ps

echo "running main.go---------------------"
go run main.go


# docker exec -it mongo_IDT /bin/bash
# mongosh