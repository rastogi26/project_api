# GoFr Netflix Project

## Overview

This project is a simple Netflix-like application built with Go and GoFr. It provides basic CRUD (Create, Read, Update, Delete) operations for managing a watchlist of movies. The application uses MongoDB as the database to store movie information.

## Table of Contents

- [Installation](#installation)
- [Dependencies](#dependencies)
- [Usage](#usage)
  - [Running the Application](#running-the-application)
- [Endpoints](#endpoints)
  - [GET /mymovies](#get-mymovies)
  - [POST /addmovie](#post-addmovie)
  - [PUT /watched/{id}](#put-watchedid)
  - [DELETE /removemovie/{id}](#delete-removemovieid)
  - [DELETE /removeallmovies](#delete-removeallmovies)
- [MongoDB Atlas](#mongodb-atlas)
- [UML Diagram](#uml-diagram)


## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/YOUR-USERNAME/YOUR-REPOSITORY
cd project_api
go mod init github.com/example

```

## Dependencies
Ensure you have Go installed. Install project dependencies using Go modules:


```bash
go get gofr.dev
go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv
```
## Usage

### Running the Application

To run the application, use the following command:

```bash
go run main.go
```
The application will start, and you can access it at http://localhost:8000.

## Endpoints

### GET /mymovies

Retrieve all movies from the watchlist.

Example:

```bash
curl http://localhost:8000/mymovies
```

### POST /addmovie

Add a new movie to the watchlist.

Example:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "Movie Title", "genre": "Action", "watched": false}' http://localhost:8000/addmovie
```
### PUT /watched/{id}

Mark a movie as watched.

Example:

```bash
curl -X PUT http://localhost:8000/watched/12334
```

### DELETE /removemovie/{id}

Remove a specific movie from the watchlist.

Example:

```bash
curl -X DELETE http://localhost:8000/removemovie/12653
```

### DELETE /removeallmovies

Remove all movies from the watchlist.

Example:

```bash
curl -X DELETE http://localhost:8000/removeallmovies
```

## MongoDB Atlas

The project uses MongoDB as a database, and the connection details are stored in the `.env` file. Ensure you have a MongoDB Atlas account and update the `.env` file with your connection string. Refer to `.env.sample` file for help.

## UML Diagram
![Alt text](<UML GO.png>)

![Alt text](<Flow Structure.png>)














