docker/build:
	docker build -t cfurl-dev .

docker/run:
	docker run \
		-v $(SSH_AUTH_SOCK):$(SSH_AUTH_SOCK) \
		-e SSH_AUTH_SOCK=$(SSH_AUTH_SOCK) \
		--name cfurl-dev --rm \
		-v $(PWD):/go/src/github.com/sh19910711/cfurl \
		-w /go/src/github.com/sh19910711/cfurl \
		-ti cfurl-dev
