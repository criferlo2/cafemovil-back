# Café Móvil Backend

Backend REST en Go para la versión 1.0.0 de Café Móvil.

## Endpoints v1

- `POST /api/v1/auth/register` — registra usuario mock con `name` y `phone`.
- `POST /api/v1/auth/guest` — crea sesión invitada mock.
- `GET /api/v1/categories` — lista categorías, con Café primero.
- `GET /api/v1/products/featured?category=&brand=&search=` — lista productos destacados con filtros opcionales.
- `GET /health` — health check.

## Product pricing contract

All products returned by `GET /api/v1/products/featured`, including search and category filters, expose:

- `price`: numeric unit amount in pesos, not centavos; nullable. An unknown price is `null`, never an inferred zero. Explicit zero is distinct from unknown.
- `currency`: `"COP"` for the current Colombian catalog.

For example, `"price": 42900, "currency": "COP"` represents COP 42,900 per unit. Clients format the amount using the Colombian locale, and display an unavailable-price label for null or missing values rather than `$0`.

The current API uses a seeded catalog. There are no product-detail or cart endpoints yet, so backend consistency with those endpoints cannot be verified in this increment. This change does not migrate MySQL persistence.

Pricing regression tests: `go test ./internal/api ./internal/catalog -v`.

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
