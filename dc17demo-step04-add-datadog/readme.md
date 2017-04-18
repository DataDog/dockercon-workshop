Add the Datadog Container

Create a Datadog trial account if you don't have one. Go to https://app.datadoghq.com/account/settings#api and add your API Key to the environment variable API_KEY

Review docker-compose.yaml to see the datadog container

Review the configuration files in datadog-conf directory

Start the containers

    docker-compose up

Open a shell on the datadog container in a new terminal


    docker-compose exec datadog bash

Run the Datadog info command

    service datadog-agent info

OR, Join all the commands into one:

    docker-compose exec datadog service datadog-agent info

As you start working with the containers more, adding features to the containers, you will find it useful to pipe the commands together, like this:

    docker-compose rm -f; docker-compose up

