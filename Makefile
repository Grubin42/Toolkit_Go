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

DOCKER_NAME = toolkit_go_dev-app-1
# MIGRATE
# Commande de base pour exécuter les migrations
# Cette variable contient la commande pour exécuter les migrations, en précisant le chemin des migrations
# et la base de données cible (ici, MySQL sur localhost avec l'utilisateur root et le mot de passe root).
MIGRATE = /go/bin/migrate -path ./migrations -database "mysql://root:root@tcp(mysql:3306)/mydatabase"

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

# Commande pour appliquer toutes les migrations
migrate-up:
	@echo "Applying all up migrations..."
	@docker exec -it $(DOCKER_NAME) $(MIGRATE) up

# Commande pour annuler la dernière migration
migrate-down:
	@echo "Reverting the last migration..."
	@docker exec -it $(DOCKER_NAME) $(MIGRATE) down 1

# Commande pour créer une nouvelle migration
# Cette règle permet de créer une nouvelle migration. Elle demande d'abord à l'utilisateur de saisir un nom pour la migration,
# puis exécute la commande pour créer un fichier de migration avec ce nom, dans le répertoire './migrations'.
migrate-new:
	@read -p "Enter migration name: " name; \
	@docker exec -it $(DOCKER_NAME) /go/bin/migrate create -ext sql -dir ./migrations -seq $$name

# Commande pour vérifier le statut des migrations
# Cette règle affiche la version actuelle des migrations appliquées dans la base de données.
migrate-status:
	@echo "Checking migration status..."
	@docker exec -it $(DOCKER_NAME) $(MIGRATE) version

.PHONY: all start up down stop rebuild delete rebuild-no-cache up-dev down-dev stop-dev rebuild-dev delete-dev rebuild-no-cache-dev migrate-up migrate-down migrate-new migrate-status