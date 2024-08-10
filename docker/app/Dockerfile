# Utiliser une image de base Golang
FROM golang:1.22-alpine

WORKDIR /app

# Installer Air pour le rechargement automatique
RUN go install github.com/air-verse/air@latest

# Copier les fichiers de mod dans le conteneur
COPY app/go.mod app/go.sum ./
RUN go mod download

# Copier tout le code source
COPY app/src/* ./

# # Construire l'application
# RUN cd src && go build -o ../bin/main .

# Lancer Air pour le rechargement automatique
CMD ["air", "-c", "/app/.air.toml"]