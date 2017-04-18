1.  Install Docker.  (https://www.docker.com/)

2.  We are going to use a bunch of docker hub images. I'll explain what these are in a bit, but for now, run:

        docker pull rabbitmq:3
        docker pull rabbitmq:3-management
        docker pull golang:1.7-alpine
        docker pull datadog/docker-dd-agent:latest
        docker pull logstash
        docker pull elasticsearch
        docker pull kibana


