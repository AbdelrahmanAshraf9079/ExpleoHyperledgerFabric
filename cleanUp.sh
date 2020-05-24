
######### Running Networks #########

docker network ls


######### Cleaning up #########

docker rm -f $(docker ps -aq)

docker volume prune

docker network prune



