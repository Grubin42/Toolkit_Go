# Toolkit_Go

## Prérequis

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.22+](https://golang.org/dl/)

## Accéder à l’application

L’application est accessible sur http://localhost.

## Installation

### 1. Cloner le dépôt

```bash
git clone https://github.com/votre-utilisateur/toolkit_go.git
cd toolkit_go
```
### 2. Ajouter un fichier .env
```bash
touch .env 
```
Ajoutez les configurations en suivant le fichier .env.example

### 3. Construire et démarrer les conteneurs
```bash
make up
```

## Utilisation de Air

Air est un outil de rechargement à chaud qui permet de recompiler et redémarrer automatiquement votre application Go lorsque des changements sont détectés dans le code source. Cela améliore le flux de travail en développement.

### 1. Démarrage de Air

Air est automatiquement démarré lorsque vous lancez la commande make up ou docker-compose up. Il surveille les fichiers de votre projet pour tout changement et les compile automatiquement.

### 2. Vérification des Logs de Compilation en mode dev 

Pour voir les informations de compilation, les erreurs, et les autres logs produits par Air, vous pouvez utiliser la commande suivante :

```bash
docker logs -f toolkit_go_dev-client-app-1
docker logs -f toolkit_go_dev-admin-app-1
```
Cette commande vous permet de suivre en temps réel les logs du conteneur app, où Air est exécuté.

## Ajouter de nouvelles dépendances

### 	1.	Importer la dépendance dans le code :
```go
import "github.com/gin-gonic/gin"
```
### 	2.	Installer la dépendance :
```bash
go get github.com/gin-gonic/gin
go mod tidy
```

### 	3.	Rebuild les conteneurs :

```bash
make rebuild
```

## Supprimer une dépendance

### 1. Retirer l’importation de la dépendance dans le code :

Supprimez la ligne de code qui importe la dépendance.
```go
// import "github.com/gin-gonic/gin"
```

### 2. Nettoyer les dépendances inutilisées :

Exécutez la commande suivante pour retirer la dépendance du projet et mettre à jour les fichiers go.mod et go.sum.

```bash
go mod tidy
```

### 3. Rebuild les conteneurs :

Pour s’assurer que la dépendance supprimée n’est plus présente, reconstruisez les conteneurs Docker :

```bash
make rebuild
```

### archi actuelle

```bash
/project-root
│
├── /cmd                       # Point d'entrée pour les différentes applications
│   ├── /client-site           # Application du site client
│   │   └── main.go
│   └── /admin                 # Application de l'administration et du toolkit
│       └── main.go
│
├── /pkg                       # Contient les packages partagés (auth, middleware, etc.)
│   ├── /auth                  # Gestion de l'authentification
│   ├── /middleware            # Middlewares HTTP (logging, auth, etc.)
│   ├── /models                # Modèles de données partagés
│   └── /utils                 # Utilitaires généraux (helpers, validation, etc.)
│
├── /internal                  # Code non exporté (logique métier spécifique, non réutilisable)
│   ├── /client                # Logiciel spécifique au client (non partagé)
│   │   ├── /handlers          # Handlers HTTP spécifiques au site client
│   │   ├── /templates         # Templates HTML pour le site client
│   │   └── /assets            # Assets personnalisés (CSS, JS, images) pour le site client
│   └── /admin                 # Logiciel spécifique à l'administration
│       ├── /handlers          # Handlers HTTP pour l'interface admin
│       ├── /templates         # Templates HTML pour l'interface admin
│       └── /assets            # Assets pour l'interface admin
│
├── /configs                   # Fichiers de configuration pour les différents environnements
│   ├── /client-site.yaml      # Configuration du site client
│   └── /admin.yaml            # Configuration de l'administration
│
├── /migrations                # Scripts de migration de base de données
│   └── ...
│
├── /docker                    # Dockerfiles et configurations Docker Compose
│   ├── Dockerfile.client      # Dockerfile pour l'application client
│   ├── Dockerfile.admin       # Dockerfile pour l'application admin
│   └── docker-compose.yml     # Configuration Docker Compose pour l'ensemble du projet
│
├── /scripts                   # Scripts pour automatiser les tâches (build, test, déploiement)
│   ├── build.sh               # Script de build pour le projet
│   └── deploy.sh              # Script de déploiement
│
└── go.mod                     # Gestion des dépendances Go
└── go.sum                     # Fichier checksum pour les dépendances
```