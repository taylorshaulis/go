# go
Now playing with the second part of the starter tutorial
https://go.dev/doc/tutorial/getting-started
Calling an external package in the code

A new stage has been added to the Dockerfile
> go mod tidy

This seems to tell go to download the mod that we are including.


Now we just build the image and execute the go run
> docker build -t hello-go .

> docker run --rm -ti hello-go go run .

