CC = go
WORKER_OUTPUT = bin/worker 
MANAGER_OUTPUT = bin/manager

WORKER_INPUT = cmd/worker/main.go
MANAGER_INPUT = cmd/manager/main.go

build:
	$(CC) build -o $(WORKER_OUTPUT) $(WORKER_INPUT)
	$(CC) build -o $(MANAGER_OUTPUT) $(MANAGER_INPUT)
run:
	$(MAKE) build
	$(MANAGER_OUTPUT)