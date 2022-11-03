build:
	go build

kind:
	kubectx kind-kind

dev: kind
	tilt up