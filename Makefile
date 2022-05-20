all: start
.PHONY: all

start:
	sudo docker-compose up --build
rebuild:
	sudo docker build -t kirigaikabuto/n50-teacher-api:latest .
	sudo docker-compose up --build
git:
	git add .
	git commit -m "feat:add update"
	git push
