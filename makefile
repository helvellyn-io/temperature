
TARGET="${REGISTRY}/${REPO}:${VERSION}"

all: 

build-app: build-app
	docker build -f builds/Dockerfile.build.app -t ${TARGET} -t latest .

push-app:  push-app
	docker push ${TARGET}

#deploy-dev: deploy-dev
#	curl -X POST http://aa21a4849593a4b72b69dfde10241cff-1397073425.us-east-2.elb.amazonaws.com/webhooks/webhook/demo



