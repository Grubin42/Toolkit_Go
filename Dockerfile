# Utiliser une image de base Golang
FROM golang:1.22.6-alpine AS builder

WORKDIR /app

# Installer migrate pour gérer les migrations
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

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