<table width="100%">
	<tr>
		<td align="left" width="70%">
			<strong>CQRSES_GROUP2</strong><br>
			<spam>Attempt to implement CQRS and Event Sourcing in Go</spam>
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

## Wiki Documentation

You can find the full documentation in the Wiki Pages :

* [Configuration & Usage](https://github.com/HETIC-MT-P2021/CQRSES_GROUP2/wiki/1-%E2%80%A2-Configuration-&-Usage)
* [Command line](https://github.com/HETIC-MT-P2021/CQRSES_GROUP2/wiki/2-%E2%80%A2-Command-line)
* [Docker structure](https://github.com/HETIC-MT-P2021/CQRSES_GROUP2/wiki/3-%E2%80%A2-Docker-structure)
* [Architecture](https://github.com/HETIC-MT-P2021/CQRSES_GROUP2/wiki/4-%E2%80%A2-Architecture)
* [Database](https://github.com/HETIC-MT-P2021/CQRSES_GROUP2/wiki/5-%E2%80%A2-Database)
* [CQRS & Event Sourcing](https://github.com/HETIC-MT-P2021/CQRSES_GROUP2/wiki/6-%E2%80%A2-CQRS-&-Event-Sourcing)
* [ADR & MVC](https://github.com/HETIC-MT-P2021/CQRSES_GROUP2/wiki/7-%E2%80%A2-ADR-&-MVC)
* [RabbitMQ](https://github.com/HETIC-MT-P2021/CQRSES_GROUP2/wiki/8-%E2%80%A2-RabbitMQ)
* [Contributing](https://github.com/HETIC-MT-P2021/CQRSES_GROUP2/wiki/9-%E2%80%A2-Contributing)

## Requirement

To use this app in an optimal way, docker is required.

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/install/)

## Quick use

Start the project with :

```shell
$ cp .env.example .env
$ make compose/build compose/up
```

You can now access the API: [http://localhost:1323/](http://localhost:1323/).
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
