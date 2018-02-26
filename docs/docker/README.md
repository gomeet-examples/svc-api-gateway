# svc-api-gateway docker usage

## Build docker image

### Regular Dockerfile

```shell
make docker
--or--
docker build -t gomeetexamples/svc-api-gateway:$(cat VERSION | tr +- __) .
```

## Use port binding on host

### 1. Launch server container

```shell
docker run -d \
    --net=network-grpc-gomeetexamples \
    -p 13000:13000 \
    --name=svc-svc-api-gateway-1 \
    -it gomeetexamples/svc-api-gateway:$(cat VERSION | tr +- __)
```

### 2. Use client on host

- Build and use cli tool

  ```shell
  $ make
  $ cd _build
  $ svc-api-gateway cli version
  $ svc-api-gateway cli services_status
  $ svc-api-gateway cli echo <uuid [string]> <content [string]>
  $ svc-api-gateway cli create_profile <gender [UNKNOW|MALE|FEMALE]> <email [string]> <name [string]> <birthday [string]>
  $ svc-api-gateway cli read_profile <uuid [string]>
  $ svc-api-gateway cli list_profile <page_number [uint32]> <page_size [uint32]> <order [string]> <exclude_soft_deleted [bool]> <soft_deleted_only [bool]> <gender [UNKNOW|MALE|FEMALE]>
  $ svc-api-gateway cli update_profile <uuid [string]> <gender [UNKNOW|MALE|FEMALE]> <email [string]> <name [string]> <birthday [string]> <created_at [string]> <updated_at [string]> <deleted_at [string]>
  $ svc-api-gateway cli delete_profile <uuid [string]>
  $ svc-api-gateway cli --address localhost:42000 version

  # more info
  svc-api-gateway help cli
  ```

- Or use HTTP/1.1 api

  ```shell
  $ curl -X GET    http://localhost:13000/api/v1/version
  $ curl -X GET    http://localhost:13000/api/v1/services/status
  $ curl -X POST   http://localhost:13000/api/v1/echo -d '{"uuid": "<string>", "content": "<string>"}'
  $ curl -X POST   http://localhost:13000/api/v1/profile/create -d '{"gender": "UNKNOW|MALE|FEMALE", "email": "<string>", "name": "<string>", "birthday": "<string>"}'
  $ curl -X POST   http://localhost:13000/api/v1/profile/read -d '{"uuid": "<string>"}'
  $ curl -X POST   http://localhost:13000/api/v1/profile/list -d '{"page_number": <number>, "page_size": <number>, "order": "<string>", "exclude_soft_deleted": <boolean>, "soft_deleted_only": <boolean>, "gender": "UNKNOW|MALE|FEMALE"}'
  $ curl -X POST   http://localhost:13000/api/v1/profile/update -d '{"uuid": "<string>", "gender": "UNKNOW|MALE|FEMALE", "email": "<string>", "name": "<string>", "birthday": "<string>", "created_at": "<string>", "updated_at": "<string>", "deleted_at": "<string>"}'
  $ curl -X POST   http://localhost:13000/api/v1/profile/delete -d '{"uuid": "<string>"}'
  $ curl -X GET    http://localhost:13000/
  $ curl -X GET    http://localhost:13000/version
  $ curl -X GET    http://localhost:13000/metrics
  $ curl -X GET    http://localhost:13000/status
  $ curl -X GET    http://localhost:42000/version
  ```

## Do not use port binding

### 1. Create a docker's network

```shell
docker network create \
    --driver bridge network-grpc-gomeetexamples &> /dev/null
```

### 2. Run server container with the previous created network

```shell
docker run -d \
    --net=network-grpc-gomeetexamples \
    --name=svc-svc-api-gateway \
    -it gomeetexamples/svc-api-gateway:$(cat VERSION | tr +- __)
```

### 3. Run clients with docker

#### Console

```shell
docker run -d \
    --net=network-grpc-gomeetexamples \
    --name=console-svc-api-gateway \
    -it gomeetexamples/svc-api-gateway:$(cat VERSION | tr +- __) console --address=svc-svc-api-gateway:13000
```

Detach console with `Ctrl + p Ctrl + q` and attach with :

```shell
docker attach console-svc-api-gateway
```

#### Client with docker

```shell
docker run \
    --net=network-grpc-gomeetexamples \
    -it gomeetexamples/svc-api-gateway:$(cat VERSION | tr +- __) cli --address=svc-svc-api-gateway:13000 <grpc_service> <params...>
```

#### Curl with docker use gomeet/gomeet-curl

[Docker Hub](https://hub.docker.com/r/gomeet/gomeet-curl/) - [Source](https://github.com/gomeet/gomeet-curl)

```shell
# use HTTP/1.1 api
docker run \
    --net=network-grpc-gomeetexamples \
    -it gomeet/gomeet-curl -X POST http://svc:13000/api/v1/-X <HTTP_VERB> http://localhost:13000/api/v1/<grpc_service> -d '<HTTP_REQUEST_BODY json format>'

# status and metrics
docker run \
    --net=network-grpc-gomeetexamples \
    -it gomeet/gomeet-curl http://svc-svc-api-gateway:13000/status

docker run \
    --net=network-grpc-gomeetexamples \
    -it gomeet/gomeet-curl http://svc-svc-api-gateway:13000/metrics
```
