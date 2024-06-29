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

# remove all image
docker rmi $(docker images -q)

# remove all container running
docker rm $(docker ps --filter status=exited -q)

# stop
docker stop $(docker ps -aq)

```

### Docker - Bind for 0.0.0.0:4000 failed: port is already allocated

```bash
sudo service docker stop

sudo rm -f /var/lib/docker/network/files/local-kv.db
```

cat backup.sql | docker exec -i mariadb /bin/bash -c "/usr/bin/mysql -u root --password=Uk1F0r3xP@ssw0rd dev_ukiforex"
docker exec -i mariadb /bin/bash -c "/usr/bin/mysql -u root --password=Uk1F0r3xP@ssw0rd dev_ukiforex" < backup.sql
### backup databases
```bash
# mysql backup
sudo su
docker exec -i CONTAINER /bin/bash -c "mysqldump -u root --password=root DATABASE" > data.sql

# mysql restore
cat backup.sql | docker exec -i CONTAINER /usr/bin/mysql -u root --password=root DATABASE

docker exec -i mariadb /bin/bash -c "/usr/bin/mysql -u root --password=Uk1F0r3xP@ssw0rd dev_ukiforex" < backup.sql

docker cp backup.sql mysql-container:/backup.sql
docker exec mysql-container bash -c "mysql -u root -p<your_password> < /backup.sql"

# postgres
docker exec -i CONTAINER /bin/bash -c "PGPASSWORD=password pg_dump --username root DATABASE" > dump.sql
```