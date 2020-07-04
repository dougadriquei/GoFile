# desafioNeoway


## Running

### Docker Compose

```shell
docker-compose up

docker-compose down
docker-compose down --remove-orphans --volumes
```

#### Local / Dev
```shell
docker-compose up --build
```

### Docker

#### Development
```shell
docker build -t desafioNeoway .
docker run -p 8095:8095 --name desafioNeoway -it desafioNeoway
```

```shell
docker stop desafioNeoway
docker start desafioNeoway
```

## Connecting throuhg PGAdmin
```shell
docker container ls
docker inspect <postgres_image_id> | grep IPAddress
```

## Reference

- [Faster builds in Docker with Go 1.11](https://blog.container-solutions.com/faster-builds-in-docker-with-go-1-11)
- [Getting started with Go modules](https://medium.com/@fonseka.live/getting-started-with-go-modules-b3dac652066d)