# APIOps demo for ngrok

This repo contains a simple Go app that creates an API with a single route. This project is designed to be containerized and deployed to a Kubernetes cluster (see `deployment.yaml` and `gateway.yaml`), with ingress handled by the [ngrok Kubernetes Operator](https://github.com/ngrok/kubernetes-ingress-controller).

There is no database, so changes will not persist between executions.

To start the API:

```bash
git clone git@github.com:joelhans/ngrok-apiops-demo.git
cd ngrok-apiops-demo
go mod init legends-api
go get golang.ngrok.com/ngrok github.com/gorilla/mux
go run main.go
```

## Usage

Add a new entry:

```bash
curl \
  -X POST \ 
  -H "Content-Type: application/json" \
  -d '{"name":"test","type":"dragon","origin":"somewhere"}' \    
  http://localhost:5000/legend
```

See existing entries:

```bash
curl \
  -X GET \ 
  -H "Content-Type: application/json" \
  http://localhost:5000/legend
```
