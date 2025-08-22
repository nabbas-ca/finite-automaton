all: build 

build: 
	go build -o ./bin/fsm main.go

test: 
	go test -coverprofile=fsm.coverage github.com/nabbas-ca/finite-automaton/fsm

