up: 
	docker-compose up

build:
	docker-compose build

fire:
	docker-compose build && docker-compose up

user-service:
	kit generate service user --dmw --gorilla