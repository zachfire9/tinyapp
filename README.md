go run main.go

docker build -t tinyapp .
docker run -p 8080:8080 tinyapp

http://localhost:8080/