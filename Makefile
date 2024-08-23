# COLORS
GREEN		= \033[1;32m
RED 		= \033[1;31m
ORANGE		= \033[1;33m
CYAN		= \033[1;36m
RESET		= \033[0m

# FOLDERS
SRCS_DIR	= ./
ENV_FILE	= ${SRCS_DIR}.env
DOCKER_DIR	= ${SRCS_DIR}docker-compose.yml
DEV_DIR     = ${SRCS_DIR}docker-compose.override.yml

# COMMANDS
DOCKER = docker compose -f ${DOCKER_DIR} --env-file ${ENV_FILE} -p toolkit_go
DOCKER_DEV = docker compose -f ${DOCKER_DIR} -f ${DEV_DIR} --env-file ${ENV_FILE} -p toolkit_go_dev

%:
	@:

all: up

start: up

up:
	@echo "${GREEN}Starting containers in production mode...${RESET}"
	@${DOCKER} up -d --remove-orphans

up-dev:
	@echo "${ORANGE}Starting containers in development mode...${RESET}"
	@${DOCKER_DEV} up -d --remove-orphans

down:
	@echo "${RED}Stopping containers...${RESET}"
	@${DOCKER} down

down-dev:
	@echo "${RED}Stopping containers in development mode...${RESET}"
	@${DOCKER_DEV} down

stop:
	@echo "${RED}Stopping containers...${RESET}"
	@${DOCKER} stop

stop-dev:
	@echo "${RED}Stopping containers in development mode...${RESET}"
	@${DOCKER_DEV} stop

rebuild:
	@echo "${GREEN}Rebuilding containers...${RESET}"
	@${DOCKER} up -d --remove-orphans --build

rebuild-dev:
	@echo "${GREEN}Rebuilding containers in development mode...${RESET}"
	@${DOCKER_DEV} up -d --remove-orphans --build

delete:
	@echo "${RED}Deleting containers...${RESET}"
	@${DOCKER} down -v --remove-orphans

delete-dev:
	@echo "${RED}Deleting containers in development mode...${RESET}"
	@${DOCKER_DEV} down -v --remove-orphans

rebuild-no-cache:
	@echo "${GREEN}Rebuilding containers with no cache...${RESET}"
	@${DOCKER} build --no-cache
	@${DOCKER} up -d --remove-orphans --build

rebuild-no-cache-dev:
	@echo "${GREEN}Rebuilding containers in development mode with no cache...${RESET}"
	@${DOCKER_DEV} build --no-cache
	@${DOCKER_DEV} up -d --remove-orphans --build

.PHONY: all start up down stop rebuild delete rebuild-no-cache up-dev down-dev stop-dev rebuild-dev delete-dev rebuild-no-cache-dev