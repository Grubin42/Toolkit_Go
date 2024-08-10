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

### 2. Vérification des Logs de Compilation

Pour voir les informations de compilation, les erreurs, et les autres logs produits par Air, vous pouvez utiliser la commande suivante :

```bash
docker-compose logs -f app
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