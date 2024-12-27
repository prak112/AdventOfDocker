# Overview
- Learning Docker and its related aspects to build and deploy containers from [AdventOfDocker](https://adventofdocker.com/)
- Amazing resource! Amazing timing! Loving it ALL!

## Topics
- [Container vs VMs](#containers-vs-vms)

- [Docker Port Mapping](#docker-port-mapping)

- [Docker Layers](#docker-layers)

- [Docker Volumes](#docker-volumes)

- [Passing Arguments](#passing-arguments)
  - [During Runtime](#during-runtime)
  - [During Build-time](#during-build-time)

- [Docker Networking](#docker-networking)
  - [Network Types](#network-types)

- [Docker Compose](#docker-compose)

- [Multistage Builds]()  /*TO BE DOCUMENTED*/


<hr>
<hr>
</br>


## Containers vs VMs
- *Containers*
  - Uses host machine to run applications
  - Source code is exposed to the host, thus a possible security vulnerability. 
  - Light-weight software to build and run applications to prevent 'Works-on-my-machine' problems

- *Virtual Machines (VMs)*
  - Isolates the source code from host
  - Runs the application in an isolated environment, thus reducing security risk to host.
  - Heavy on the host machine due to compartmentalized Memory, CPU and Disk usage.

<hr>
</br>

## Docker Port Mapping

- Docker port mapping syntax :
 - Port number declared in `EXPOSE` layer in the Docker Image, known as `CONTAINER_PORT` 
 - Port hosting location on local server known as `HOST_PORT`

```pwsh
  docker run -p <HOST_PORT>:<CONTAINER_PORT> <IMAGE_NAME>
```

- To bind to a specific interface or hosting location :

```pwsh
  docker run -p 127.0.0.1:8080:80 nginx
```

- To publish all exposed ports on the host :

```pwsh
  docker run -P <IMAGE_NAME>
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


## Docker Volumes
- Containers are stateless or ephemeral, i.e., while running data is written on a temporary layer on top of the container, once it stops the data layer is wiped clean.
- Volumes are used to store data written on top of a container by mounting the declared volume to a defined path inside the application
- This can be done during runtime or could be orchestrated using `docker-compose.yaml`
- This helps the data to persist in the volume even after the container is stopped.


- Using docker commands ;
```pwsh
  # create volume
  docker volume create myvolume

  # mount volume on container
  docker run -p 8080:8080 -v myvolume:/app_data/dir myapp

  # delete volume, if not needed
  docker volume rm myvolume
```

- Using Docker Compose, define the YAML file ;
```yaml
  services:
  api:
    build: .
    ports:
      - "8080:5050"
    env_file:
      - .env
    volumes:
      - mydata:/data
    networks:
      - myapp-network
    develop:
      watch:
        - path: ./
          action: rebuild
  volumes:
    mydata:
  networks:
    myapp-network:
```

- Execute the YAML file in docker;
```pwsh
  docker compose up
```
- For more details, check [Docker Compose](#docker-compose)


### Volumes Types
- *Anonymous Volumes*
  - Random volume is created and data is written into it. 
  - Harder to manage but useful for temporary data.
  
  ```pwsh
    docker run -v /app_data/dir myapp
  ```

- *Managed Volumes*
  - These are the volumes discussed above, where a volume is labelled and addressed specifically to mount onto the container for storing data

- *Bind Mounts*
  - Mounts host machine directory onto the container.
  - Recommended for development purposes for sharing information between Host and Container, such as:
    - Source code,
    - Configuration files,
    - Generate files

  ```pwsh
    # docker run -v <HOST_PATH>:<CONTAINER_PATH> <IMAGE_NAME>
    docker run -v D:/Projects/project1/app/info.txt:/info.txt myapp
  ```

  - Read [Docker docs](https://docs.docker.com/engine/storage/bind-mounts/) for more variations on Bind Mounts.


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


## Docker Networking

- To create a new network :

```pwsh
  #SYNTAX : docker network create <NETWORK_NAME>
  docker network create myapp-network
```

- Multiple containers can be labeled and run within the same network. This also allows communication between these containers through HTTP by using container name in the URL.

```pwsh
  #SYNTAX : docker run -d --name <CONTAINER_NAME> --network <NETWORK_NAME> <IMAGE_NAME>
  docker run -d --name api --network myapp-network

  #SYNTAX for example: to tunnel host an API : 
  # docker run -d --name <CONTAINER_NAME> --network <NETWORK_NAME> nginx 
  docker run -d --name frontend --network myapp-netowork nginx:latest

  # execute communication between containers through HTTP using hostnames
  docker exec -it curl http://api:8080
```

- Custom hostnames can be set instead of using container names :

```pwsh
  docker run --name api --hostname custom-api --network myapp-network <IMAGE_NAME>
```

- The hostnames can be verified after running the container from within :

```pwsh
  docker exec -it api hostname
  # returns custom-api
```


### Network Types 
- Docker supports several network types:

  - `bridge`: The default network driver. Good for containers on a single host (used above)
  - `host`: Removes network isolation, container uses host’s network directly
  - `none`: Disables networking completely
  - `overlay`: For connecting containers across multiple Docker hosts
  - `macvlan`: Assigns a MAC address to containers, making them appear as physical devices

- [Official Documentation about Networking](https://docs.docker.com/network/drivers/)


<hr>
</br>


## Docker Compose
- To define and run multiple containers involved in running an application (such as web-client, mobile-client, server, admin) Docker Compose is used. 

- Docker Compose YAML file, `docker-compose.yaml` defined at the root of the project hosting repositiory which defines the `services` involved such as `client`, `server`, `api` 
- Also, Ports mapping, Networks or Volumes created in the containers can be defined in the `docker-compose.yaml`

- Docker Compose will automatically map the ports, create networks, create volumes, build image and run container as specified in `Dockerfile` image configuration of the service.

### Example
- In the case of **ICD Symptom Checker** application, it consists of :
  - *Web Client*
  - *Mobile Client*
  - *NodeJS Server*
  - *Blazor Admin Client*

- Each of these repositories at their root directories will have their own `Dockerfile` to configure the Docker Image.
- Finally, the hosting repository, *ICD Symptom Checker* will have the `docker-compose.yaml` at its root directory to configure the build and execution of the containers defined above.


### Basic Docker Compose Commands#

- `docker compose up` - Start services

- `docker compose up -d` - Start in detached mode

- `docker compose down` - Stop and remove containers

- `docker compose ps` - List running services

- `docker compose logs` - View service logs

- `docker compose build` - Build or rebuild services


<hr>
</br>

