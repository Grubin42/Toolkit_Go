# Documentation du Projet Toolkit_Go

## Introduction

Toolkit_Go est un projet web développé en Go, utilisant Docker pour la conteneurisation des services, et gérant les migrations de base de données avec migrate. Le projet est structuré de manière modulaire, permettant un développement et une maintenance faciles. Il propose une application simple avec un CRUD pour la gestion des utilisateurs, avec un environnement de développement et de production distinct.

## Objectif du Projet

L’objectif du Toolkit_Go est de fournir une base solide pour développer des applications web en Go, avec des fonctionnalités de gestion d’utilisateurs. Le projet est conçu pour être facilement extensible et maintenable.

## Prérequis

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.22+](https://golang.org/dl/)

## Table des Matières

1. [Introduction](#introduction)
2. [Objectif du Projet](#objectif-du-projet)
3. [Prérequis](#prérequis)
4. [Accéder à l’application](#accéder-à-lapplication)
5. [Installation](#installation)
   1. [Cloner le dépôt](#1-cloner-le-dépôt)
   2. [Installation avec le Makefile](#2-installation-avec-le-makefile)
   3. [Ajouter un fichier .env](#3-ajouter-un-fichier-env)
6. [Différence entre le Mode Production et le Mode Développement](#différence-entre-le-mode-production-et-le-mode-développement)
   1. [Mode Production](#mode-production)
   2. [Mode Développement](#mode-développement)
7. [Utilisation de Air](#utilisation-de-air)
   1. [Démarrage de Air](#1-démarrage-de-air)
   2. [Vérification des Logs de Compilation en mode dev](#2-vérification-des-logs-de-compilation-en-mode-dev)
8. [Ajouter de nouvelles dépendances](#ajouter-de-nouvelles-dépendances)
   1. [Importer la dépendance dans le code](#1-importer-la-dépendance-dans-le-code)
   2. [Installer la dépendance](#2-installer-la-dépendance)
9. [Supprimer une dépendance](#supprimer-une-dépendance)
10. [Gestion des Migrations](#gestion-des-migrations)
    1. [Appliquer les Migrations](#1-appliquer-les-migrations)
    2. [Annuler la Dernière Migration](#2-annuler-la-dernière-migration)
    3. [Créer une Nouvelle Migration](#3-créer-une-nouvelle-migration)
    4. [Vérifier le Statut des Migrations](#4-vérifier-le-statut-des-migrations)
11. [Architecture du Projet](#architecture-du-projet)
    1. [Arborescence du Projet](#arborescence-du-projet)
    2. [Explication de l’Architecture](#explication-de-larchitecture)
12. [Informations Supplémentaires](#informations-supplémentaires)
    1. [Recommandations](#recommandations)

## Accéder à l’application

L’application est accessible sur http://localhost.

## Installation

### 1. Cloner le dépôt

```bash
git clone https://github.com/Grubin42/Toolkit_Go.git
cd toolkit_go
```

### 2. Installation avec le Makefile

Le projet utilise un Makefile pour simplifier l’exécution des commandes Docker et autres tâches courantes.

- Pour démarrer le projet en mode développement :
```bash
make up-dev
```

- Pour démarrer le projet en mode production :
```bash
make up
```

- Pour arrêter les conteneurs en mode développement :
```bash
make down-dev
```
- Pour arrêter les conteneurs en mode production :
```bash
make down
```

### 3. Ajouter un fichier .env

```bash
touch .env 
```
Ajoutez les configurations en suivant le fichier .env.example

## Différence entre le Mode Production et le Mode Développement

### Mode Production

- Le mode production utilise le fichier docker-compose.yml seul, sans le fichier docker-compose.override.yml.
- Les conteneurs sont construits et exécutés avec les images optimisées pour la production.
- Les modifications de code nécessitent un redémarrage manuel des conteneurs.

## Mode Développement

- Le mode développement combine les fichiers docker-compose.yml et docker-compose.override.yml.
- Les conteneurs sont construits avec des images optimisées pour le développement, incluant par exemple Air pour le rechargement automatique du serveur Go à chaque modification de code.
- Les volumes montés permettent un développement itératif sans avoir à reconstruire les conteneurs manuellement.

## Utilisation de Air

Air est un outil de rechargement à chaud qui permet de recompiler et redémarrer automatiquement votre application Go lorsque des changements sont détectés dans le code source. Cela améliore le flux de travail en développement.

### 1. Démarrage de Air

Air est automatiquement démarré lorsque vous lancez la commande make up-dev. Il surveille les fichiers de votre projet pour tout changement et les compile automatiquement.

### 2. Vérification des Logs de Compilation en mode dev 

Pour voir les informations de compilation, les erreurs, et les autres logs produits par Air, vous pouvez utiliser la commande suivante :

```bash
docker logs -f toolkit_go_dev-app-1
```

Cette commande vous permet de suivre en temps réel les logs du conteneur app, où Air est exécuté.

## Ajouter de nouvelles dépendances

### 1.	Importer la dépendance dans le code :

```go
import "github.com/gin-gonic/gin"
```

### 2.	Installer la dépendance :

```bash
go get github.com/gin-gonic/gin
go mod tidy
```
Pour s’assurer que la dépendance soit présente, reconstruisez les conteneurs Docker :

```bash
make rebuild-dev
```

## Supprimer une dépendance

Pour supprimer une dépendance inutile, vous pouvez simplement l’enlever du fichier go.mod, puis exécuter :
```go
go mod tidy
```

Pour s’assurer que la dépendance supprimée n’est plus présente, reconstruisez les conteneurs Docker :

```bash
make rebuild-dev
```

## Gestion des Migrations

Les migrations de base de données sont gérées par l’outil migrate.

### 1. Appliquer les Migrations

Pour appliquer toutes les migrations :

```bash
make migrate-up
```

### 2. Annuler la Dernière Migration

Pour annuler la dernière migration appliquée :

```bash
make migrate-down
```

### 3. Créer une Nouvelle Migration

Pour créer une nouvelle migration SQL :

```bash
make migrate-new
```
Cette commande vous demandera de fournir un nom pour la migration, et créera les fichiers SQL correspondants dans le répertoire migrations.

### 4. Vérifier le Statut des Migrations

Pour vérifier la version actuelle des migrations appliquées :

```bash
make migrate-status
```



## Architecture du Projet

### Arborescence du Projet

```bash
/project-root
├── Dockerfile
├── Dockerfile.dev
├── Makefile
├── README.md
├── cmd
│   └── app
│       └── main.go
├── configs
├── docker-compose.override.yml
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── controllers
│   │   └── user_controller.go
│   ├── models
│   │   └── user.go
│   ├── routers
│   │   └── router.go
│   └── services
│       └── user_service.go
├── migrations
│   ├── 000001_create_users_table.down.sql
│   └── 000001_create_users_table.up.sql
├── pkg
│   └── database
│       └── database.go
├── templates
│   ├── index.html
│   └── users.html
├── tmp
│   ├── build-errors.log
│   └── main
└── uploads
```

### Explication de l’Architecture

- cmd/app/main.go : Point d’entrée de l’application. C’est ici que le serveur est initialisé et que les routes sont configurées.
- internal/controllers/ : Contient les contrôleurs, qui gèrent la logique de traitement des requêtes HTTP. Par exemple, user_controller.go gère toutes les actions liées aux utilisateurs.
- internal/models/ : Contient les modèles de données. Par exemple, user.go définit la structure User utilisée dans le projet.
- internal/routers/ : Contient la configuration des routes de l’application. Par exemple, router.go initialise toutes les routes de l’application.
- internal/services/ : Contient la logique métier. Par exemple, user_service.go contient les fonctions pour gérer les utilisateurs (CRUD).
- pkg/database/ : Contient le code pour gérer la connexion à la base de données.
- migrations/ : Contient les fichiers SQL pour gérer les migrations de la base de données.
- templates/ : Contient les fichiers HTML pour l’affichage côté serveur.
- tmp/ : Dossier temporaire utilisé par Air pour le développement.
- uploads/ : Dossier destiné à stocker les fichiers uploadés par les utilisateurs.

## Informations Supplémentaires

### Recommandations

- Gestion des Secrets : Pensez à utiliser des outils comme Docker Secrets ou des fichiers .env sécurisés pour gérer les informations sensibles (comme les mots de passe).
- Tests : Il est conseillé d’écrire des tests unitaires et d’intégration pour chaque composant du projet pour garantir la qualité du code.
- Documentation Continue : Mettez à jour cette documentation à chaque modification majeure du projet pour qu’elle reste pertinente et utile.