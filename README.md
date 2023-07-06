# Go + Integração Contínua
Github Actions, Golangci-lint e Codecov

## run entire project
```
docker-compose up
```

## product-api
### build
```
make build
```

### run
```
make run
```

## product-db
### build
```
docker build -t product-db .
```

### run
```
docker run --volume=$HOME/datadir:/var/lib/mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=supersecret -e MYSQL_DATABASE=ProductDb -e MYSQL_USER=AdminUser -e MYSQL_PASSWORD=AdminPassword product-db
```