# APIOps demo for ngrok

This repo contains:

- A simple Go app that creates a simple API with a single route.

To run the Go-based API:

```
git clone git@github.com:joelhans/ngrok-apiops-demo.git
cd ngrok-apiops-demo
go mod init legends-api
go get golang.ngrok.com/ngrok github.com/gorilla/mux
go run main.go
```
