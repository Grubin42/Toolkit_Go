# Utilisation de l'image officielle de PostgreSQL
FROM postgres:15-alpine

# Copie du script d'initialisation
COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432