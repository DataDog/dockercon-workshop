1.  Install Docker. If you are on a Mac or Windows, the easiest option is the Docker Toolbox (https://www.docker.com/products/docker-toolbox). There is a Docker for Mac beta program that uses xhyve instead of VirtualBox and VMWare Fusion, but there are still some issues with it. This session will use the standard released version of Docker instead.

2.  Create a new machine with docker-machine. Having 4G RAM for the machine is recommended. If using VMWare Fusion, use this command:


        docker-machine create -d vmwarefusion --vmwarefusion-memory-size 4196 --vmwarefusion-cpu-count 2 dockercon2016

    If using VirtualBox, use this command:

        docker-machine create -d virtualbox --virtualbox-memory 4196 --virtualbox-cpu-count 2 vbdockercon2016

    To see a list of your currently running docker-machines, run:

        docker-machine ls

    To start using the created docker-machine for all of your docker and docker-compose commands, run:

        docker
