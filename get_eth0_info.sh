#!/bin/bash

# Run the command and store the output in a variable
output=$(ip -s -d link show eth0)

# Save the output as JSON
json_output=$(echo "$output" | jq -Rs '.')

# Save the JSON output to a file
echo "$json_output" > output.json

