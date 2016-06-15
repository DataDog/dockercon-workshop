Start with this command:

    docker run -d --hostname rabbit --name rabbit rabbitmq:3

See your running containers:

    docker ps

See the output of the container:

    docker logs <container-id>

Stop the container:

    docker stop <container-id>

Get rid of `-d` to not run as daemon

    docker run --hostname rabbit --name rabbit rabbitmq:3

Notice that you cannot run this again since the container exists. Remove it to run again

    docker rm <container-id>

Go to https://hub.docker.com, search for rabbitmq. Click on the first item. Review the instructions for using the image.

Stop and remove the rabbit container. Run again using 3-management. Look at the options for exposing a port

    docker run --hostname rabbit --name rabbit -p 8080:15672 rabbitmq:3-management


Open http://localhost:8080

