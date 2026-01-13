# --- Stage 1 : Build front avec Node ---
FROM node:20-alpine AS frontend

WORKDIR /app

# Copier uniquement ce qui est nécessaire pour installer les deps
COPY ihm/package*.json ./ihm/
WORKDIR /app/ihm
RUN npm install

# Copier le reste du front
COPY ihm/ .

# Build front (ton esbuild.js)
RUN node esbuild.js

# --- Stage 2 : Build back avec Go ---
FROM golang:tip-alpine3.22 AS backend

WORKDIR /app

COPY back/go.mod ./back/
RUN cd back && go mod download

# Copier le code back
COPY back/ ./back

# Copier le front généré depuis le stage Node
COPY --from=frontend /app/ihm ./ihm

# Build Go
WORKDIR /app/back
RUN go build -o /usr/local/bin/server main.go

# --- Stage 3 : Image finale ---
FROM alpine:latest

WORKDIR /app

COPY --from=backend /usr/local/bin/server .
COPY --from=backend /app/ihm ./ihm

EXPOSE 8080

CMD ["./server"]
