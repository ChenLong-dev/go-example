package main

/*
goctl docker -go hello.go
docker build -t hello:v1 -f Dockerfile .
docker run -it --rm hello:v1

*/

func main() {
	println("hello world!")
}
