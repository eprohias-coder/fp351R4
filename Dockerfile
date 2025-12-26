#---------- BUILD STAGE ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copiamos dependencias primero (mejora caché)
COPY go.mod ./
RUN go mod download

# Copiamos el resto del código
COPY . .

# Compilamos la app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o webapp

# ---------- RUNTIME STAGE ----------
FROM alpine:3.20

RUN apk add --no-cache ca-certificates wget


WORKDIR /app

# Copiamos binario y estáticos
COPY --from=builder /app/webapp /app/webapp
COPY --from=builder /app/static /app/static

# Puerto de la app
EXPOSE 8080

# Usuario no root (seguridad)
USER nobody

# Arranque
ENTRYPOINT ["/app/webapp"]
