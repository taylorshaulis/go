# go
Now on to making our first go module
https://go.dev/doc/tutorial/create-module

This seems very setup around running go locally.........
I am not running go locally......
So adding some extra steps to do all the go stuff from within docker

First lets build a go image thats real stripped down.
>FROM golang:1.22
>
>WORKDIR /usr/src/app

Really we don't need to build this. We could just docker run the golang:1.22 image but... why not add extra steps for the fun of it

So build the image and tag it as go
> docker build -t go .

Now run the image and set the entrypoint to /bin/bash so we can run commands in the image but still have it mounted our host's directory.
> docker run -ti --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp go /bin/bash

We are now in the running container and should even see we are root. Woot for root!

In the container run the next step of the tutorial to create the module file
> go mod init example.com/greetings

Again this is doing it with extra steps. This all could have been combined into....
> docker run -ti --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.22 go mod init example.com/greetings

But hey this is for learning!