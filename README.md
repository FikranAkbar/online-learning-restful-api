# Online Learning RESTful API
Golang Portofolio Project: RESTful API For Online Learning Application

<div id="top"></div>

# Table of Contents

<ol>
  <li><a href="#built-with">Built With</a></li>
  <li><a href="#project-structure">Project Structure</a></li>
  <li>
    <a href="#getting-started">Getting Started</a>
    <ul>
      <li><a href="#prerequisites">Prerequisites</a></li>
      <li><a href="#environment-variables">Environment Variables</a></li>
      <li><a href="#database-migration">Database Migration</a></li>
    </ul>
  </li>
  <li><a href="#documentation">Deployment and Documentation</a></li>
  <li><a href="#features-roadmap">Features Roadmap</a></li>
</ol>

# Built With

- Go: 1.19.3^
- MySQL : 8.0.32^
- Wire: 0.5.0^

<p align="right">(<a href="#top">back to table of contents</a>)</p>

# Project Structure

#### `/app/database`
Information about database migration and seeder.

#### `/app/router`
All the route of the api.

#### `/config`
Configuration file that loaded in form of struct from env file.

#### `/controller`
Directory for the functions to process request from registered routes.

#### `/di`
Dependency injection for making automation in function provider.

#### `/exception`
Error handling from the api failed request.

#### `/helper`
Small functions like converter, panic helper, etc.

#### `/model`
Objects that interprets the model in the database and formats the request or response.

#### `/repository`
Layer that handling logic interaction with the database.

#### `/service`
Layer for handling main business login of the request routes.

#### `/test`
Integration tests for all the api.

<p align="right">(<a href="#top">back to table of contents</a>)</p>

# Getting Started

## Prerequisites

1. First we need to clone the repo into our local machine

```
git clone https://github.com/FikranAkbar/online-learning-restful-api.git
```

2. then move to folder just created

```
cd online-learning-restful-api/
```

3. then download all the required package

```
go mod tidy
```

Run Project in Development Environment
```
go run main.go
```

## Environment Variables

Configuration file is write using environment variables. This file must be exists in root directory with complete example file as below:
```env
#filename: app.env

DB_HOST= <<YOUR DB HOST>>
DB_PORT= <<YOUR DB PORT>>
DB_USER= <<YOUR DB USER>>
DB_PASSWORD= <<YOUR DB PASSWORD>>
DB_NAME= <<YOUR DB NAME>>
DB_NAME_FOR_TEST= <<YOUR DB TEST NAME>>

MIDTRANS_CLIENT_KEY= <<YOUR MIDTRANS CLIENT KEY>>
MIDTRANS_SERVER_KEY= <<YOUR MIDTRANS SERVER KEY>>
```

## Database Migration

1. Make file app.env and copy content from example.env

2. Update DB_NAME_DEV, DB_USERNAME_DEV, and DB_PASSWORD_DEV based on your local

3. Run the test in app/database/db_test.go

```
go test app/database/db_test.go
```

<p align="right">(<a href="#top">back to table of contents</a>)</p>

## Deployment and Documentation

- [API Documentation (Swagger)](api_spec.json)
- [Database Documentation (DBDiagram)](https://dbdiagram.io/d/621f197f54f9ad109a438bf0)

<p align="right">(<a href="#top">back to table of contents</a>)</p>

## Features Roadmap

- [x] Authentication
- [x] Course (Module, Webinar, Quiz)
- [x] Expert Mentor
- [x] Industry Insight Article
- [x] Payment (Course Transaction)

<p align="right">(<a href="#top">back to table of contents</a>)</p>

