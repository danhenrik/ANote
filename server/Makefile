help:
	@echo "=========================================================================\n \
		run - Run backend server\n \
		infra [up|down] - Run infrastructure (DB, Background services, ES)\n \
		full - Run infrastructure and server application\n \
		build - Build the server executable in bin folder \
		\n=========================================================================" 

full:
	@echo "=========================================================================\n \
		Initializing infra.. \
		\n========================================================================="
	make infra up
	@echo "=========================================================================\n \
		Initializing server... \
		\n========================================================================="
	make run

infra:
	@echo "=========================================================================\n \
		Usage: make infra [up|down] \
		\n========================================================================="

up:
	docker compose up -d

down:
	docker compose down

run: 
	go run cmd/anote/main.go

build:
	go build -o bin/anote cmd/anote/main.go

	
