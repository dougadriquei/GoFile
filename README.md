# desafio Neoway

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

##  Informations

### init.sql
Script da entidade principal(Purchase) utilizada para represtar cada linha do arquivo.txt 

### purchase.go
Responsável pelas querys e a estrutura(struct) onde é armazenado cada linha do arquivo 

### readFile.go
Responsável pela extração de cada linha do arquivo.txt

### psqlInfo.go
Responsável pelas configurações do acesso ao banco de dados postegres, para quando for rodar em localhost o app.

### Docker commands (utils)

```
docker build -t desafioneoway .
docker run --publish 8080:8080 desafioneoway
docker stop desafioneoway
docker start desafioneoway
docker rm $(docker ps -a -q) 
docker rmi $(docker images -a -q)
docker-compose down --remove-orphans --volumes 
docker container ls
docker volume ls
docker ps -a
docker images
docker rmi <imagem>
```

## Reference

- [Faster builds in Docker with Go 1.11](https://blog.container-solutions.com/faster-builds-in-docker-with-go-1-11)
- [Getting started with Go modules](https://medium.com/@fonseka.live/getting-started-with-go-modules-b3dac652066d)