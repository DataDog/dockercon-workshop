Review **docker-compose.yaml** file

    version: '2'
    services:
      rabbit:
        image: rabbitmq:3-management
        ports:
         - "8080:15672"
         - "5672:5672"

Start the container

    docker-compose up

To stop if CTRL-C doesn't stop it:

    docker-compose stop <service>

To remove the containers defined in the local docker-compose.yaml

    docker-compose rm   //add -f to force remove without a confirmation prompt



