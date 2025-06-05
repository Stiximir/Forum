# Étape 1 : Build de l'application
FROM golang:1.24 AS builder

WORKDIR /app

# Copier les fichiers nécessaires pour les dépendances
COPY back/go.mod back/go.sum ./
RUN go mod download

# Copier le reste du code
COPY back/main/ ./

# Compiler l'application
RUN go build -o forum-app main.go

# Étape 2 : Image minimale pour exécuter le binaire
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
RUN useradd -m forumuser

WORKDIR /home/forumuser/app

# Copier le binaire
COPY --from=builder /app/forum-app .

# Copier les templates si présents
COPY template ./template

USER forumuser

EXPOSE 8080

CMD ["./forum-app"]