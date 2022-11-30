init:
	sudo service postgresql stop
	docker start postgres
	
run:
	sudo service postgresql stop
	docker start postgres
	go run main.go

mockgen:
	mockgen -package mockdb -destination  ./api/mocks/course_mock.go -source=api/repository/course_repository.go
