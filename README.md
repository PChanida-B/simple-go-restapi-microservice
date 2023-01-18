# Sample REST APIs Microservices

This project is a sample REST APIs with microservices for my study and my practice.

## Technologies used

- Golang
- Gin-Gonic framework
- Gorm framework
- MySQL(MariaDB)
- RESTful API
- Microservices Architecture
- Hexagonal Architecture
- Clean Architecture
- Docker

## Requirements

API lists
**GET** /tokenz
**GET** /resources
**GET** /resources/ID
**POST** /resources
**PULL** /resources/ID
**DELETE** /resources/ID

## Non-Functional Requirements

- Use Gin-Gonic framework
- Configuration in environment
- Gracefully shutting down
- JWT authentications
- Dockerfile

## VS Code Extensions

- for Dababase connection: mtxr.sqltools-driver-mysql
- for API testing: humao.rest-client

## Makefile

maria: to start MariaDB docker container

restapi-image: to build sample REST APIs docker image

restapi-container: to start sample REST APIs docker container

## Run sample REST APIs

start MariaDB, REST API contianer

run command:
> make maria
> make restapi-image
> make restapi-container


## API testing

run file test example for API testing
> test/create_data.http
> test/read_data.http
> test/update_data.http
> test/delete_data.http
