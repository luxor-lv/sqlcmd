build:
	go build -C ./src -o ../bin/sql

run: build
	./bin/sql -t "select * from qcore.transaction_validation_groups"