#!/bin/sh

trap killgroup 2

SCRIPT=$(readlink -f "$0")
SCRIPTPATH=$(dirname "$SCRIPT")

GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
[ -z $GOPATH ] && GOPATH=$(go env GOPATH)

[ -z $GOMEET_EXEC_TYPE ] && GOMEET_EXEC_TYPE="make" # go, make

[ -z $GOMEET_JWT_SECRET ] && GOMEET_JWT_SECRET=""

if [ -z $GOMEET_SVC_PROFILE_DB_DSN ]
then
	[ -z $GOMEET_SVC_PROFILE_DB_USERNAME ] && GOMEET_SVC_PROFILE_DB_USERNAME="gomeet"
	[ -z $GOMEET_SVC_PROFILE_DB_PASSWORD ] && GOMEET_SVC_PROFILE_DB_PASSWORD="totomysql"
	[ -z $GOMEET_SVC_PROFILE_DB_SERVER ] && GOMEET_SVC_PROFILE_DB_SERVER="localhost"
	[ -z $GOMEET_SVC_PROFILE_DB_PORT ] && GOMEET_SVC_PROFILE_DB_PORT="3306"
	[ -z $GOMEET_SVC_PROFILE_DB_DATABASE ] && GOMEET_SVC_PROFILE_DB_DATABASE="svc_profile"
	GOMEET_SVC_PROFILE_DB_DSN="$GOMEET_SVC_PROFILE_DB_USERNAME:$GOMEET_SVC_PROFILE_DB_PASSWORD@tcp($GOMEET_SVC_PROFILE_DB_SERVER:$GOMEET_SVC_PROFILE_DB_PORT)/$GOMEET_SVC_PROFILE_DB_DATABASE"
fi;

killgroup(){
	echo killing...
	kill 0
}

# SUB-SERVICES DEFINITION : run.sh
cd $GOPATH/src/github.com/gomeet-examples/svc-echo
CGO_ENABLED=0 go run $GO_RUN_FLAGS \
	-ldflags '-extldflags "-lm -lstdc++ -static"' \
	-ldflags "-X github.com/gomeet-examples/svc-echo/service.version=$(cat VERSION) -X github.com/gomeet-examples/svc-echo/service.name=gomeet-svc-echo" \
	main.go serve \
		-d \
		--jwt-secret "$GOMEET_JWT_SECRET" \
		--address ":13001" &

cd $GOPATH/src/github.com/gomeet-examples/svc-profile
CGO_ENABLED=0 go run $GO_RUN_FLAGS \
	-ldflags '-extldflags "-lm -lstdc++ -static"' \
	-ldflags "-X github.com/gomeet-examples/svc-profile/service.version=$(cat VERSION) -X github.com/gomeet-examples/svc-profile/service.name=svc-profile" \
	main.go serve \
		-d \
		--jwt-secret "$GOMEET_JWT_SECRET" \
		--mysql-migrate \
		--mysql-dsn "$GOMEET_SVC_PROFILE_DB_DSN" \
		--address ":13002" &
# END SUB-SERVICES DEFINITION : run.sh

cd $SCRIPTPATH/../
SERVER_OPTS='serve
				-d
				--jwt-secret "$GOMEET_JWT_SECRET"
				--svc-echo-address "localhost:13001"
				--svc-profile-address "localhost:13002"
				--address ":13000"'

case $GOMEET_EXEC_TYPE in
	"go")
		CMD='CGO_ENABLED=0 go run
			-ldflags "-extldflags \"-lm -lstdc++ -static\""
			-ldflags "-X github.com/gomeet-examples/svc-profile/service.version=$(cat VERSION) -X github.com/gomeet/svc-profile/service.name=svc-profile"
			main.go'
		eval $CMD $SERVER_OPTS
		break
		;;
	"make")
		make
		eval _build/svc-api-gateway $SERVER_OPTS
		break
		;;
	*)
		echo "Error : unknow $GOMEET_EXEC_TYPE value for GOMEET_EXEC_TYPE [go|make] allowed"
		;;
esac

wait

