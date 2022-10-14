TARGET=main test

start: 
	docker-compose up -d
stop: 
	docker-compose down
in:
	docker exec -i -t formappgo_app_1 bash
$(TARGET):
	docker-compose exec app go run $@.go