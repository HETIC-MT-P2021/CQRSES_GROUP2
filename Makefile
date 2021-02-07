DOCKER_COMPOSE ?= docker-compose
GO_WORKDIR ?= ""

.PHONY: help
help: # Provides help information on available commands
	@printf "Usage: make <command>\n\n"
	@printf "Commands:\n"
	@awk -F ':(.*)## ' '/^[a-zA-Z0-9%\\\/_.-]+:(.*)##/ { \
	  printf "  \033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)

.PHONY: compose/build
compose/build: ## Build all Docker images of the project
	@$(DOCKER_COMPOSE) build

.PHONY: compose/up
compose/up: ## Start all containers (in the background)
	@$(DOCKER_COMPOSE) up -d

.PHONY: compose/down
compose/down: ## Stops and deletes containers and networks created by "up".
	@$(DOCKER_COMPOSE) down

.PHONY: compose/restart
compose/restart: compose/down compose/up ## Restarts all containers

.PHONY: compose/start
compose/start: ## Starts existing containers for a service
	@$(DOCKER_COMPOSE) start

.PHONY: compose/stop
compose/stop: ## Stops containers without removing them
	@$(DOCKER_COMPOSE) stop

.PHONY: compose/purge/local
compose/purge/local:
	@$(DOCKER_COMPOSE) down -v --rmi local

.PHONY: compose/purge
compose/purge: compose/purge/local ## Stops and deletes containers, volumes, images (local) and networks created by "up".

.PHONY: compose/purge/all
compose/purge/all: ## Stops and deletes containers, volumes, images (all) and networks created by "up".
	@$(DOCKER_COMPOSE) down -v --rmi all

.PHONY: compose/rebuild
compose/rebuild: compose/down compose/build compose/up ## Rebuild the project

.PHONY: compose/rebuild/%
compose/rebuild/%: ## Rebuild a specific service ex: make compose/rebuild/<service_name>
	@$(DOCKER_COMPOSE) build $*
	@$(MAKE) compose/up

.PHONY: compose/top
compose/top: ## Displays the running processes.
	@$(DOCKER_COMPOSE) top

DOCKER_MONITOR_TABLE ?= 'table {{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.NetIO}}\t{{.BlockIO}}'
.PHONY: compose/monitor
compose/monitor: ## Display of container(s) resource usage statistics
	@$(DOCKER_COMPOSE) ps -q | tr '\n' ' ' | docker stats --format $(DOCKER_MONITOR_TABLE) --no-stream

.PHONY: compose/monitor/follow
compose/monitor/follow: ## Display a live stream of container(s) resource usage statistics
	@$(DOCKER_COMPOSE) ps -q | tr '\n' ' ' | docker stats --format $(DOCKER_MONITOR_TABLE)

DOCKER_INPECT_FORMAT__AWK ?= "'\''{{.Name}} : {{range $$p, $$confs := .NetworkSettings.Ports}}{{range $$conf := $$confs}}{{$$p}} -> {{$$conf.HostIp}}:{{$$conf.HostPort}}{{else}}{{$$p}} -> Not exposed{{end}}\t{{else}}No exposed ports{{end}}'\''"
.PHONY: compose/urls
compose/urls: ## Get project's URL
	@echo "------------------------------------------------------------"
	@echo "You can access your project at the following URLS:"
	@echo "------------------------------------------------------------"
	@$(DOCKER_COMPOSE) ps -q | awk '{ \
		cmd_docker_inspect = sprintf("docker inspect --format=%s %s", ${DOCKER_INPECT_FORMAT__AWK}, $$0) ; \
		cmd_docker_inspect | getline docker_inspect ; close(cmd_docker_inspect) ; \
		gsub(/0.0.0.0/, "http://localhost", docker_inspect) ; \
		split(docker_inspect, urls, "\t") ; \
		printf "\n%s\n", urls[1] ; \
		i = 2 ; while (i < length(urls)) { \
		index_tab = index(docker_inspect,":") ; \
		printf "%*s %s\n", index_tab, "", urls[i]; i++ \
		} ; \
	}'
