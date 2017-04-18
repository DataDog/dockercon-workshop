# Step 02 - Using Docker Compose instead of docker run

Review **docker-compose.yaml** file

    version: '2'
    services:
      rabbit:
        image: rabbitmq:3.6-management
        ports:
         - "8080:15672"

The first line mentions the version number of the compose yaml file syntax. This will be updated to 3 soon. Then you have a list of services. This file only has 1 service, called rabbit. You can call services anything you like. This service, aka 'container', is based on the image rabbitmq:3.6-management. Finally port 15672 is exposed as port 8080.

Start the container

    docker-compose up

To stop if CTRL-C doesn't stop it:

    docker-compose stop <service>
    docker-compose stop    // to stop all services in the docker-compose file

To remove the containers defined in the local docker-compose.yaml

    docker-compose rm   //add -f to force remove without a confirmation prompt




