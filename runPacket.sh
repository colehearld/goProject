#!/bin/bash

# Read the container IDs from the file
container_ids=($(cat container_ids.txt))

# Loop through the container IDs and run each container
for container_id in "${container_ids[@]}"; do
  # Run the container using the Docker CLI
  docker start "$container_id"
done
