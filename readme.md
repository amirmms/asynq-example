redis port on localhost : ``6379``

to run dashboard :
``
docker run --rm --name asynqmon -p 8080:8080 hibiken/asynqmon --redis-addr='host.docker.internal:6379'
``
