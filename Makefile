export IMAGE := "afarid/k8s-events-logger:0.0.3"
export NAMESPACE := "kube-system"

build:
	docker build -t ${IMAGE} .

push:
	docker push ${IMAGE}
load:
	kind load docker-image ${IMAGE}

deploy:
	envsubst < deployment/deployment.yaml | kubectl apply -f -

remove:
	envsubst < deployment/deployment.yaml | kubectl  delete -f -
