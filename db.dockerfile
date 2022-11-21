# start with base image
FROM mysql:8.0.23

# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
COPY ./database/*.sql /docker-entrypoint-initdb.d/



# docker build myimg
# docker run -e MYSQL_ROOT_PASSWORD=rootpass myimg
# docker ps
# docker exec -it nervous_margulis bash
# mysql -u root -prootpass
# ----mysql
# show databases;
