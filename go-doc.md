# Documentation Go

## 1. Documentation : Différence entre go get et go install

### Introduction

Dans le développement en Go, il est crucial de comprendre la différence entre go get et go install pour gérer correctement les dépendances de ton projet et les outils en ligne de commande. Cette documentation explique ces différences à travers des exemples pratiques avec les outils migrate, air, et go-admin.

#### go get

Utilisation :

go get est principalement utilisé pour ajouter des bibliothèques en tant que dépendances dans ton projet Go. Lors de l’exécution de go get, Go télécharge la bibliothèque spécifiée et la stocke dans le cache des modules. Le fichier go.mod est mis à jour pour inclure cette nouvelle dépendance.

Exemple avec go-admin :

Imaginons que tu souhaites utiliser go-admin dans ton projet pour gérer l’interface d’administration. Pour ajouter cette dépendance :

```bash
go get github.com/go-admin-team/go-admin
go mod tidy
```

Cela met à jour ton fichier go.mod et télécharge go-admin dans le cache des modules. La dépendance est maintenant prête à être utilisée dans ton projet.

Résultat :

Le code source de go-admin est téléchargé dans le cache des modules Go, généralement situé dans ~/go/pkg/mod. Tu peux ensuite utiliser cette dépendance directement dans ton code via des importations.

#### go install

Utilisation :

go install est utilisé pour installer des outils en ligne de commande. Contrairement à go get, go install compile l’outil spécifié et installe le binaire résultant dans le répertoire bin de ton environnement Go (par exemple, ~/go/bin).

Exemple avec migrate :

Pour installer l’outil de migration de base de données migrate afin de gérer les migrations dans ton projet :

```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Cette commande compile l’outil migrate et le place dans ~/go/bin. Tu peux ensuite exécuter migrate depuis n’importe quel répertoire en utilisant la ligne de commande.

Résultat :

migrate est installé en tant qu’outil en ligne de commande et est accessible directement via le terminal.

### Comparaison

- go get est utilisé pour ajouter des dépendances à ton projet, modifiant ainsi le fichier go.mod et rendant la dépendance disponible pour être utilisée dans ton code.
- go install est utilisé pour installer des outils CLI, compilant le binaire et le plaçant dans le répertoire bin de ton environnement Go, sans affecter go.mod.

### Quand installer un outil directement dans le Dockerfile ?

Il peut être judicieux d’installer certains outils directement dans le Dockerfile dans les cas suivants :

- Portabilité : Si tu veux t’assurer que ton environnement Docker est entièrement autonome et que toutes les dépendances, y compris les outils CLI, sont installées dans l’image Docker, c’est une bonne idée d’utiliser go install dans le Dockerfile.
- Reproductibilité : En installant des outils via Dockerfile, tu garantis que tout le monde utilise exactement la même version de ces outils, ce qui réduit les risques de problèmes liés à des versions différentes.
- Exemple : 

Pour installer migrate dans un Dockerfile :

```dockerfile
# Utiliser une image de base Golang
FROM golang:1.22.6-alpine AS builder

WORKDIR /app

# Installer migrate pour gérer les migrations dans l'image Docker
RUN go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copier les fichiers go.mod et go.sum pour télécharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le reste du code source
COPY . .

# Construire l'application
RUN go build -o main ./cmd/app

# Étape finale : créer une image minimale
FROM alpine:latest

WORKDIR /app
RUN apk --no-cache add ca-certificates

# Copier l'application compilée et migrate
COPY --from=builder /app/main .
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Exposer le port sur lequel l'application écoute
EXPOSE 8080

# Démarrer l'application
CMD ["./main"]
```