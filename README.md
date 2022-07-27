This project's intent is to show how to create a Hello World dockerized Go Application and deploy it with Traefik's reverse proxy.

To use this, clone the repo and follow these steps:

1) Go to helloworld folder and build the Go Hello World docker image:

docker image build --tag helloworld:1.0 .
