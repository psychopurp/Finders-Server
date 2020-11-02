RUN_NAME=finder_server
CURDIR=$(shell pwd)
RUNTIME_ROOT=${CURDIR}/output
RUNTIME_CONF_ROOT=${RUNTIME_ROOT}/conf
RUNTIME_LOG_ROOT=${RUNTIME_ROOT}/log

build:
	mkdir -p output output/bin output/conf output/log
	go build -o output/bin/${RUN_NAME}


clean:
	rm -rf output

dev:build
	rm -f output/conf/*
	cp conf/*_dev* output/conf/
	exec ${RUNTIME_ROOT}/bin/${RUN_NAME} -conf=${RUNTIME_CONF_ROOT} -log=${RUNTIME_LOG_ROOT}


product:export GIN_MODE=release
product:build
	rm -f output/conf/*
	cp conf/*_product* output/conf/
	exec "${RUNTIME_ROOT}/bin/${RUN_NAME}" -conf=${RUNTIME_CONF_ROOT} -log=${RUNTIME_LOG_ROOT}

out:
	echo "${RUNTIME_ROOT}  ${RUNTIME_CONF_ROOT} ${RUNTIME_LOG_ROOT}"

app:
	go build main.go
	nohup ./main &