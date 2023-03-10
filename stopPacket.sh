#!/bin/bash

function stopPackets(){
  # Read the container IDs from the file
  container_ids=($(cat container_ids.txt))

  # Loop through the container IDs and run each container
  for container_id in "${container_ids[@]}"
  do
    # Stop the container using the Docker CLI
    docker stop -d $container_id
  done
}