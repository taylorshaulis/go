# go
Continue with the create module tutorial
https://go.dev/doc/tutorial/create-module

This has not been going well breaking up these parts

Ok combined the greets and hello code from the previous steps into the respective greetings and hello directories.

We wanna run some go commands so going to run the docker container from this directory.

> docker run -ti --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp go /bin/bash

Follow the walk threw to setup the modules replace line. Kinda interesting. Basically override the the 'example.com/greetings' reference to '../greetings'.

After all that goodness we are able to run our go run . command and get the desired output.