# Utilisation de l'image officielle de Go
FROM golang:1.22-alpine

# Définition du répertoire de travail
WORKDIR /app

# Copie du fichier go.mod et go.sum pour gérer les dépendances
COPY app/go.mod app/go.sum ./
RUN go mod download

# Copie du reste des fichiers de l'application
COPY app/ .

# Compilation de l'application
RUN go build -o main .

# Exposition du port
EXPOSE 8080

# Commande pour lancer l'application
CMD ["./main"]