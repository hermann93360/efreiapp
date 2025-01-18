FROM golang:1.20 as builder

# Définir le répertoire de travail
WORKDIR /app

# Copier le code source dans le conteneur
COPY . .

# Compiler l'application
RUN go build -o efreiapp

# Étape finale
FROM debian:buster-slim
WORKDIR /root/
COPY --from=builder /app/efreiapp .

# Exposer le port de l'application
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["./efreiapp"]
