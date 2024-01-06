# node-discovery

## run as Docker container

```
# create docker image
docker build -t node-discovery .

# start docker container
docker container run -d -p 3200:3200 7373:7373 7946:7946 --name node-discovery node-discovery
```


## health check
```
grpcurl --plaintext 127.0.0.1:3200 NodeDiscovery.GetMembers
```