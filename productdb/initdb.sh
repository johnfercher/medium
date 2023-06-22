#!/bin/bash

read -p 'Container ID: ' containerId
echo
echo Loading data into $containerId.

cat create-scripts/create_db.sql | docker exec -i $containerId mysql -u root --password=supersecret ProductDb