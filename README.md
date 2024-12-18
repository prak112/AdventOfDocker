# Overview
- Learning Docker and its related aspects to build and deploy containers from [AdventOfDocker](https://adventofdocker.com/)
- Amazing resource! Amazing timing! Loving it ALL!

## Docker Port Mapping

- Docker port mapping syntax :
 - Port number declared in `EXPOSE` layer in the Docker Image, known as `CONTAINER_PORT` 
 - Port hosting location on local server known as `HOST_PORT`

```pwsh
  docker run -p HOST_PORT:CONTAINER_PORT IMAGE_NAME
```

- To bind to a specific interface or hosting location :

```pwsh
  docker run -p 127.0.0.1:8080:80 nginx
```

- To publish all exposed ports on the host :

```pwsh
  docker run -P IMAGE_NAME
```


<hr>
</br>


## Docker Layers

- Instructions `CMD` and `EXPOSE` do not create new layers, they just modify the Image's metadata

- The following instructions create different layers in the Image :
  - `FROM`
  - `COPY`
  - `RUN`
  - `ADD`

- These instructions typically involve adding or modifying files in the filesystem of the image, which results in the creation of new layers.

<hr>
</br>


## Passing Arguments

### During Runtime
- Environment variables can be passed during runtime as a file or seperate variables.
- However, passing it as a file is more secure and mapping the ports with regard to `EXPOSE`d port inside the Image with port in the `.env`

```pwsh
  docker run --env-file .env -p HOST_PORT:CONTAINER_PORT IMAGE_NAME
```

### During Build-time
- For example, output file execution process can be altered based on the Environment for building the Image.
- `RUN` command changes to ignoring debugging information using the flags `-ld flags "-s -w"` from file output to reduce compiled binary size

```dockerfile
  FROM golang
  COPY . .
  ARG IS_PRODUCTION=false
  # if production, add compilation flag
  RUN if [ "$IS_PRODUCTION" = "true" ]; then go build -o main main.go -ldflags "-s -w"; else go build -o main main.go; fi
  CMD ["./main"]
```

- Docke Image is built as follows :

```pwsh
  docker build -t helloworld-go-http --build-arg IS_PRODUCTION=true .
```

<hr>
</br>

