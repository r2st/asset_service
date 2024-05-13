#!/bin/bash

# Build and run containers
docker-compose up --build -d

# Wait for everything to stabilize
sleep 10

# Remove unused and dangling images
docker image prune -f
