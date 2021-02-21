<table width="100%">
	<tr>
		<td align="left" width="70%">
			<strong>CQRSES_GROUP2</strong><br>
			Attempt to implement CQRS and Event Sourcing in Go
		</td>
		<td align="right" width="25%">
            <a href="https://goreportcard.com/report/github.com/HETIC-MT-P2021/CQRSES_GROUP2">
                <img src="https://goreportcard.com/badge/github.com/HETIC-MT-P2021/CQRSES_GROUP2" alt="Go Report Card">
			</a>
            <a href="https://github.com/HETIC-MT-P2021/aio-group2-proj01/blob/master/LICENSE">
                <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License: MIT"/>
            </a>
		</td>
	</tr>
	<tr>
		<td>A Hetic student project.</td>
		<td align="center">
			<img src="https://user-images.githubusercontent.com/27848278/80025966-ab059800-84e1-11ea-9e37-41a3ddcbda89.png" width="100"/>
		</td>
	</tr>
</table>


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

### Rebuild only the service container `go`

```shell
$ make compose/rebuild/go
```

## Use the command line

To list available commands, either run `make` with no parameters or execute `make help`:

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
  go/lint                        Run golangci-lint (All-In-One config)
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

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/fredgoum">
        <img src="https://github.com/fredgoum.png" width="100px;"/><br>
        <sub><b>Alfred Goumou</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/Akecel">
        <img src="https://github.com/Akecel.png" width="100px;"/><br>
        <sub><b>Axel Rayer</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/t-hugo">
        <img src="https://github.com/t-hugo.png" width="100px;"/><br>
        <sub><b>Hugo Tinghino</b></sub>
      </a>
    </td>
  </tr>
</table>

## Licence

This project is licensed under the [MIT License](https://opensource.org/licenses)  - see the LICENSE.md file for
details.
