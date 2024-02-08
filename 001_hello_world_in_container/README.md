# go
Container for golang

https://hub.docker.com/_/golang

This example is just going to get hello world running as described here https://go.dev/doc/tutorial/getting-started

This just builds the image locally and then you can exec into the image to run the hello world



from this directory you run

docker build -t hello-go .

which builds the image

You can then run

docker run --rm -ti golang /bin/sh

which will get you into a running container of this image.

From there its the go run .

# go run .
Hello, World!

And that is it. Go running within a container