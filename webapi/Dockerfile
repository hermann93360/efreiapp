# Étape de construction
FROM golang:1.20 AS builder

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers go.mod et go.sum
COPY go.mod ./

# Télécharger les dépendances
RUN go mod download

# Copier le reste du code source
COPY . .

# Compiler l'application
RUN go build -o efreiapp

# Étape finale
FROM debian:buster-slim

# Définir le répertoire de travail
WORKDIR /root/

# Copier l'exécutable depuis l'étape de construction
COPY --from=builder /app/efreiapp .

# Exposer le port sur lequel l'application écoute
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["./efreiapp"]
