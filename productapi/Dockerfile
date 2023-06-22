# stage de build
FROM golang:1.20 AS build
WORKDIR /app
COPY .. /app
RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/main.go

# stage imagem final
FROM scratch
WORKDIR /app
COPY --from=build /app/api ./
EXPOSE 8081
CMD [ "./api" ]