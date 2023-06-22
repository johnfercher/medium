#!/bin/bash

docker run --volume=$HOME/datadir:/var/lib/mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=supersecret -e MYSQL_DATABASE=ProductDb -e MYSQL_USER=AdminUser -e MYSQL_PASSWORD=AdminPassword product-db