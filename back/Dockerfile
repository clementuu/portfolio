FROM golang:tip-alpine3.22

WORKDIR /app

# Copie les fichiers de modules et téléchargez les dépendances
COPY back/go.mod ./
RUN go mod tidy

# Copie les pages du front
COPY ihm/ ./ihm

# Copie le reste du code source
COPY back .

# Expose le port sur lequel votre application Go s'exécute
EXPOSE 8080

# Commande pour lancer l'application
CMD ["go", "run", "main.go"]
