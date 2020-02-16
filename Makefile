hello.exe:
	go build .

hello:
	env GOOS=linux GOARCH=arm go build

docker: hello
	docker build -t phlatphrog/hello .

docker-run: docker
	#docker run -p 8888:8080 phlatphrog/hello
	docker run phlatphrog/hello