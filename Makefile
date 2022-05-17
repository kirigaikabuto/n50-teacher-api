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

GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/setdata-questionnaire-api

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web