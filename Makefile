all: start
.PHONY: all

start:
	docker-compose up --build
rebuild:
	docker build -t kirigaikabuto/n50-teacher-api:latest .
	docker-compose up --build
git:
	git add .
	git commit -m "feat:add update"
	git push