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

