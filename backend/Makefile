# Manually define main variables

ifndef APP_PORT
export APP_PORT = 8080
endif

ifndef APP_HOST
export APP_HOST = 127.0.0.1
endif

# parse additional args for commands

ROOT_FOLDER = .

args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

run:  ##@Application Run application server
	go run $(ROOT_FOLDER)/cmd/golang_graphs --rootPath $(ROOT_FOLDER)

run_test:  ##@Application Run application server
	TESTING="testing" go run $(ROOT_FOLDER)/cmd/golang_graphs --rootPath $(ROOT_FOLDER)

# Это просто пример того как запускать бенчмарки, если запустить из makefile, то он не отработает
bench:
	go test -bench=BenchmarkSimplest -benchmem -benchtime=1x

create_swagger:
	swag init -g cmd/golang_graphs/main.go
