build:
	go build

docker-build:
	docker build .

kind:
	kubectx kind-kind

dev: kind
	tilt up