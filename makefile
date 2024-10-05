include .env
# run application server, use command "make run"
run:
	go run main.go

# build for production, use command "make build"
build:
	GOOS=linux GOARCH=amd64 go build -o ./ikct_ed
	upx ikct_ed
	tar -czf ikct_ed.tar.gz ikct_ed .env.sample 

db_up:
	migrate -path database/ -database "postgresql://$(DBUSER):$(DBPASS)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=disable" -verbose up

db_down:
	migrate -path database/ -database "postgresql://$(DBUSER):$(DBPASS)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=disable" -verbose down $(version)
