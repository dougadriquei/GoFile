# desafio Neoway

##  Informations

### init.sql
Script da entidade principal (Purchase) utilizada para represtar cada linha do arquivo.txt 

### purchase.go
Arquivo que tem as querys e o DTO (Pruchase struct)

## Running

### Docker Compose

```shell
docker-compose up

docker-compose down
```

#### Local / Dev
```shell
docker-compose up --build
```

### Docker

#### Development
```shell
docker build -t desafioneoway .

Go to file storage.go (change port postgress local)

docker run --publish 8080:8080 desafioneoway
```

```shell
docker stop desafioneoway

docker start desafioneoway
```

```shell
utils:
docker rm $(docker ps -a -q) 
docker rmi $(docker images -a -q)
```

```shell
docker build -t my-golang-app .
docker-compose down --remove-orphans --volumes 
```

## Reference

- [Faster builds in Docker with Go 1.11](https://blog.container-solutions.com/faster-builds-in-docker-with-go-1-11)
- [Getting started with Go modules](https://medium.com/@fonseka.live/getting-started-with-go-modules-b3dac652066d)