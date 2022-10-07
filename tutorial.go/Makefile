TARGET=main test

start: 
	docker-compose up -d
stop: 
	docker-compose down
in:
	docker exec -i -t tutorialgo_app_1 bash
$(TARGET):
	docker-compose exec app go run $@.go