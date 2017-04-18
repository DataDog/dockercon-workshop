# Step 01 - Getting Started with Docker

The first thing we want to do is run a container. You do this with the `docker run` command.

Start with this command:

    docker run -d --hostname rabbit --name rabbit rabbitmq:3.6

-d means run as a daemon. rabbitmq:3.6 means to use the rabbitmq image on DockerHub. Specifically use the one tagged with '3.6'.

See your running containers:

    docker ps

See the output of the container:

    docker logs <container-id | container-name>
    // you can use either the id or the name of the container

Stop the container:

    docker stop <container-id | container-name>

Get rid of `-d` to not run as daemon:

    docker run --hostname rabbit --name rabbit rabbitmq:3.6

Notice that you cannot run this again since the container exists. Remove it using the following command.

    docker rm <container-id | container-name>

Now run that docker run command without the -d.

Go to https://hub.docker.com, search for rabbitmq. Click on the first item. Review the instructions for using the image.

Stop and remove the rabbit container. Run again using 3.6-management. Look at the options for exposing a port

    docker run --hostname rabbit --name rabbit -p 8080:15672 rabbitmq:3.6-management

-p 8080:15672 means expose the container's port 15672 as port 8080 on the host.

Open http://localhost:8080 and play around with the interface. The default user and password can be found in the Docker Hub page for rabbitmq.

Stop and remove the container before moving on.
