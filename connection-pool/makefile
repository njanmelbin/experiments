# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

postgres:
	docker run --name c_postgres \
	-e POSTGRES_PASSWORD=mysecretpassword \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_DB=postgres \
	-v my_pgdata:/var/lib/postgresql/data \
	-p 5431:5432 \
	-d postgres:17.6
