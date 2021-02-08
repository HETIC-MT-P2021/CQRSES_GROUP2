<h1 align="center"> CQRSES </h1>


## About

This application based on the HTTP protocol allows, through an API, to manage the following issues:

* Individual user account with authentication, authorization and account creation system.


## Requirement

To use this app in an optimal way, docker is required.

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/install/)


## Usage

Start the project with :

```shell
$ cp .env.example .env
$ make compose/build compose/up
```

You can now access the API: [http://localhost:1323/](http://localhost:1323/).

## Use the command line

To list available commands, either run make with no parameters or execute make help:

```shell
$ make help
Usage: make <command>

Commands:
  compose/build                  Build all Docker images of the project
  compose/up                     Start all containers (in the background)
  compose/down                   Stops and deletes containers and networks created by "up".
  compose/restart                Restarts all containers
  compose/start                  Starts existing containers for a service
  compose/stop                   Stops containers without removing them
  compose/purge                  Stops and deletes containers, volumes, images (local) and networks created by "up".
  compose/purge/all              Stops and deletes containers, volumes, images (all) and networks created by "up".
  compose/rebuild                Rebuild the project
  compose/rebuild/%              Rebuild a specific service ex: make compose/rebuild/<service_name>
  compose/top                    Displays the running processes.
  compose/monitor                Display of container(s) resource usage statistics
  compose/monitor/follow         Display a live stream of container(s) resource usage statistics
  compose/urls                   Get project's URL
```

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
