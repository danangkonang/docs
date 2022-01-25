# Docker

```bash
# build image
docker build --tag web:1.0 .

# build tag name for upload to docker hub
docker tag <image name or id> <repo name>/<image name or id>:1.0

# example
docker tag b027ac1f4e81 danangkonang/web:3.0

# build container
docker container create --name <container name> -p <to port>:<from port> <from image>

# example
docker container create --name dev -p 5000:80 danangkonang/web:3.0

# running container
docker container start <container name or id>

# remove image
docker rmi <image id>

```