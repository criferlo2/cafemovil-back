# Café Móvil Backend

Backend REST en Go para la versión 1.0.0 de Café Móvil.

## Endpoints v1

- `POST /api/v1/auth/register` — registra usuario mock con `name` y `phone`.
- `POST /api/v1/auth/guest` — crea sesión invitada mock.
- `GET /api/v1/categories` — lista categorías, con Café primero.
- `GET /api/v1/products/featured?category=&brand=&search=` — lista productos destacados con filtros opcionales.
- `GET /health` — health check.

## Ejecutar local

```bash
go test ./...
go run ./cmd/api
```

## Docker

```bash
docker compose up --build
```

El `docker-compose.yml` levanta la API y MySQL 8.4 con el esquema relacional inicial en `db/schema.sql`. En este incremento la API responde con repositorios mock/en memoria para avanzar el contrato con front sin bloquearse por persistencia robusta.
