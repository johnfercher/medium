# Go + Integração Contínua
Github Actions, Golangci-lint e Codecov

## run entire project
```
docker-compose up
```

## product-api
### build
```
docker build -t product-api .
```

### run
```
docker run -p 8081:8081 product-api
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