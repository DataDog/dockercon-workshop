Add the Datadog Container

Create a Datadog trial account if you don't have one. Go to https://app.datadoghq.com/account/settings#api and add your API Key to the environment variable API_KEY

Review docker-compose.yaml to see the datadog container

Review the configuration files in datadog-conf directory

Start the containers

    docker-compose up

Open a shell on the datadog container in a new terminal

    docker ps -f "name=datadog"  // make note of the container id
    docker exec -it <container-id> bash

Run the Datadog info command

    service datadog-agent info

OR, Join all the commands into one:

    docker exec -it `docker ps -f "name=datadog" -q` service datadog-agent info

