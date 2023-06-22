#!/bin/bash

read -p 'Container ID: ' containerId
echo
echo Loading data into $containerId.

cat create-scripts/create_tables.sql | docker exec -i $containerId mysql -u root --password=supersecret ProductDb
cat create-scripts/insert_data.sql | docker exec -i $containerId mysql -u root --password=supersecret ProductDb