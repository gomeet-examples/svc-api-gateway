# svc-api-gateway usage

## Basic usage

- Run server

```shell
svc-api-gateway serve --address <server-address>

# serve gRPC and HTTP multiplexed on localhost:3000
svc-api-gateway serve --address localhost:3000

# serve gRPC on localhost:3000 and HTTP on localhost:3001
svc-api-gateway serve --grpc-address localhost:3000 --http-address localhost:3001

# more info
svc-api-gateway help serve
```

- Run cli client

  ```shell
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

- Run console client

```shell
$ svc-api-gateway console --address=localhost:3000
INFO[0000] svc-api-gateway console  Exit=exit HistoryFile="/tmp/svc-api-gateway-62852.tmp" Interrupt="^C"
└─┤svc-api-gateway-0.1.8+dev@localhost:13000├─$ help
INFO[0002] HELP :

	┌─ version
	└─ call version service

	┌─ services_status
	└─ call services_status service

	┌─ echo <uuid [string]> <content [string]>
	└─ call echo service

	┌─ create_profile <gender [UNKNOW|MALE|FEMALE]> <email [string]> <name [string]> <birthday [string]>
	└─ call create_profile service

	┌─ read_profile <uuid [string]>
	└─ call read_profile service

	┌─ list_profile <page_number [uint32]> <page_size [uint32]> <order [string]> <exclude_soft_deleted [bool]> <soft_deleted_only [bool]> <gender [UNKNOW|MALE|FEMALE]>
	└─ call list_profile service

	┌─ update_profile <uuid [string]> <gender [UNKNOW|MALE|FEMALE]> <email [string]> <name [string]> <birthday [string]> <created_at [string]> <updated_at [string]> <deleted_at [string]>
	└─ call update_profile service

	┌─ delete_profile <uuid [string]>
	└─ call delete_profile service

	┌─ service_address
	└─ return service address

	┌─ jwt [<token>]
	└─ display current jwt or save none if it's set

	┌─ console_version
	└─ return console version

	┌─ tls_config
	└─ display TLS client configuration

	┌─ help
	└─ display this help

	┌─ exit
	└─ exit the console


└─┤svc-api-gateway-0.1.8+dev@localhost:13000├─$ unknow
WARN[0003] Bad arguments : "unknow" unknow
└─┤svc-api-gateway-0.1.8+dev@localhost:13000├─$ exit
```

- HTTP/1.1 usage (with curl):

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

- Get help

```shell
svc-api-gateway help

# or get help directly for a command
svc-api-gateway help <command[serve|cli|console]>
```

## Tests

- Use make directive

```shell
make test
```

- Unit tests

```shell
cd service
go test
```

- Functional tests (with an embedded server)

```shell
svc-api-gateway functest -e
```

- Load tests

```shell
svc-api-gateway loadtest --address <multiplexed server address> -n <number of sessions> -s <concurrency level>
```

## Mutual TLS authentication

- Create a Certificate Authority:

```shell
hack/gen-ca.sh gomeetexamples_ca
ls data/certs
```

- Create two key pairs with the common name "localhost":

```shell
hack/gen-cert.sh server gomeetexamples_ca
./gencert.sh client gomeetexamples_ca
ls data/certs
```

- Run the server with its TLS credentials:

```shell
svc-api-gateway serve \
    --address localhost:3000 \
    --ca data/certs/gomeetexamples_ca.crt \
    --cert data/certs/server.crt \
    --key data/certs/server.key
```

- Run the clients with their TLS credentials:

```shell
svc-api-gateway cli <grpc_service> <params...> \
    --address localhost:3000 \
    --ca data/certs/gomeetexamples_ca.crt \
    --cert data/certs/client.crt \
    --key data/certs/client.key

svc-api-gateway console \
    --address localhost:3000 \
    --ca data/certs/gomeetexamples_ca.crt \
    --cert data/certs/client.crt \
    --key data/certs/client.key
```

## JSON Web Token support

JSON Web Token validation can be enabled on the server by providing a secret key:

```shell
svc-api-gateway serve --jwt-secret foobar
```

The token subcommand is used to generate a JWT from the secret key:

```shell
svc-api-gateway token --secret-key foobar
```

Then the cli and console subcommands can use the generated token for authentication against the JWT-enabled server:

```shell
svc-api-gateway cli --jwt <generated token> <grpc_service> <params...>
svc-api-gateway console --jwt <generated token>
```

JWT validation can be tested on the HTTP/1.1 endpoints by providing the bearer token in the "Authorization" HTTP header:

```shell
TOKEN=`svc-api-gateway token --secret-key foobar`
curl -H "Authorization: Bearer $TOKEN" -X <HTTP_VERB> http://localhost:13000/api/v1/<grpc_service> -d '<HTTP_REQUEST_BODY json format>'
```


