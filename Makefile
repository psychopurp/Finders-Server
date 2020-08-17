RUN_NAME=finder_server
CURDIR=$(shell pwd)
RUNTIME_ROOT=${CURDIR}/output
RUNTIME_CONF_ROOT=${RUNTIME_ROOT}/conf
RUNTIME_LOG_ROOT=${RUNTIME_ROOT}/log

run:
	cd output
	exec ${RUNTIME_ROOT}/bin/${RUN_NAME} -conf=${RUNTIME_CONF_ROOT} -log=${RUNTIME_LOG_ROOT}

build:
	mkdir -p output output/bin output/conf output/log
	go build -o output/bin/${RUN_NAME}

linux-build:
	mkdir -p output output/bin output/conf output/log
	rm -f output/conf/*
	cp conf/*_product* output/conf/
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o output/bin/${RUN_NAME}

linux-run:
	cd output
	fresh ${RUNTIME_ROOT}/bin/${RUN_NAME} -conf=${RUNTIME_CONF_ROOT} -log=${RUNTIME_LOG_ROOT}


clean:
	rm -rf output

dev:build
	rm -f output/conf/*
	cp conf/*_dev* output/conf/
	cd output
	exec ${RUNTIME_ROOT}/bin/${RUN_NAME} -conf=${RUNTIME_CONF_ROOT} -log=${RUNTIME_LOG_ROOT}

	
product:export GIN_MODE=release
product:build
	rm -f output/conf/*
	cp conf/*_product* output/conf/
	# cd output
	fresh ${RUNTIME_ROOT}/bin/${RUN_NAME} -conf=${RUNTIME_CONF_ROOT} -log=${RUNTIME_LOG_ROOT}

out:
	echo "${RUNTIME_ROOT}  ${RUNTIME_CONF_ROOT} ${RUNTIME_LOG_ROOT}"

sync:
	scp -r ./output finder@123.56.104.212:~/app/ 

sync-make:
	scp  Makefile finder@123.56.104.212:~/app/ 