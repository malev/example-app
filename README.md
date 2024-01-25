# example-app

Minimalistic example app for K8s

`go get github.com/joho/godotenv`
`go run main.go`

Testing:

`docker build . -t malev/example-app`
`docker run -e PORT=4000 -e PROJECT_NAME=nginx -p 4000:4000 malev/example-app`
