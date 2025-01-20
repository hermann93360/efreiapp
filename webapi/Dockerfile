# Étape de construction
FROM golang:1.21 AS builder

WORKDIR /app

# Copier les fichiers go.mod et go.sum
COPY go.mod ./

# Télécharger les dépendances
RUN go mod download

# Copier le code source
COPY . .

# Compiler l'application
RUN go build -o efreiapp

# Étape finale
FROM ubuntu:22.04

WORKDIR /root/

# Installer les dépendances nécessaires si besoin
RUN apt-get update && apt-get install -y libc6 && apt-get clean

# Copier l'exécutable depuis l'étape de construction
COPY --from=builder /app/efreiapp .

# Exposer le port sur lequel l'application écoute
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["./efreiapp"]
