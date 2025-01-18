# Étape de construction
FROM golang:1.21 AS builder

WORKDIR /app

# Copier les fichiers go.mod et go.sum
COPY webapi/go.mod webapi/go.sum ./

# Télécharger les dépendances
RUN go mod download

# Copier le code source
COPY webapi/ .

# Compiler l'application
RUN go build -o efreiapp

# Étape finale
# Utiliser une image Debian plus récente pour inclure GLIBC_2.34
FROM debian:bullseye-slim

WORKDIR /root/

# Installer les dépendances nécessaires si besoin
RUN apt-get update && apt-get install -y libc6 && apt-get clean

# Copier l'exécutable depuis l'étape de construction
COPY --from=builder /app/efreiapp .

# Exposer le port sur lequel l'application écoute
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["./efreiapp"]
