<h1 align="center"> CQRSES </h1>


## About

This application based on the HTTP protocol allows, through an API, to manage the following issues:

* Individual user account with authentication, authorization and account creation system.


## Requirement

To use this app in an optimal way, docker is required.

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/install/)


## Configuration

You only have the *.env* file to configure :

```
APP_URL=http:localhost:1323
APP_NAME=NAME
DB_HOST=db
DB_PORT=3306
DB_NAME=database
DB_USER=user
DB_PASSWORD=password
```

## Usage

Start the project with :

```
docker-compose up --build
```

You can now access the API: [http://localhost:1323/](http://localhost:1323/).

## Test

This project has a *tests* directory containing all the test files for each controller.

Simply use the test go command to run your tests :

```
$ cd tests
$ go test
PASS
ok      project/tests        0.029s
```

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **Axel Rayer** - *Project author*
* **Hugo Tinghino** - *Project author*

See also the list of contributors who participated in this project.

## Licence

This project is licensed under the [MIT License](https://opensource.org/licenses)  - see the LICENSE.md file for details.
