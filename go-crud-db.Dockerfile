FROM mysql:latest

#copy files
COPY db.sql /docker-entrypoint-initdb.d/db.sql

EXPOSE 3306