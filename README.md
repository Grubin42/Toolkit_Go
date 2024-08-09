# Toolkit_Go

## Prérequis

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.22+](https://golang.org/dl/)

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

## Ajouter de nouvelles dépendances

## 	1.	Importer la dépendance dans le code :
```go
import "github.com/gin-gonic/gin"
```
## 	2.	Installer la dépendance :
```bash
go get github.com/gin-gonic/gin
go mod tidy
```

## 	3.	Rebuild les conteneurs :

```bash
make rebuild
```

### Accéder à l’application

L’application est accessible sur http://localhost.