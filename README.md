## Docker Commands to setup Cassandra

docker run -p 9042:9042 --rm --name cassandra -d cassandra:4.0

Wait for Docker to create Cassandra Container and a superuser of name cassandra (can be viewved in dockerhub logs)

docker exec -it cassandra /bin/bash

cqlsh

CREATE KEYSPACE restfulapi WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE restfulapi;

CREATE TABLE users ( uuid uuid PRIMARY KEY , name text , email text , password text , seller boolean);

CREATE TABLE purchases ( uuid uuid PRIMARY KEY , useruuid uuid , bookuuid uuid , timestamp timestamp);

CREATE TABLE books ( uuid uuid PRIMARY KEY , thumbnail text , name text, description text, author text, genre text, ratings int, price int, availableQuantity int);
