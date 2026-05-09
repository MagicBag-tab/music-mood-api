# Music Mood API — Backend

API REST para el **Music Mood Tracker**. Escrita en Go puro (`net/http`, sin frameworks), con PostgreSQL como base de datos. Se despliega con Docker Compose.

🔗 **Frontend repo:** https://github.com/MagicBag-tab/music-mood-api.git

🌐 **Live demo:** https://music-mood-api-1.onrender.com

📖 **Swagger UI:** 

---

## Stack

| Componente | Tech |
|-----------|------|
| Lenguaje | Go 1.23 (`net/http` puro) |
| Base de datos | PostgreSQL 15 |
| Contenedores | Docker + Docker Compose |
| Dependencias | `lib/pq` (driver PG), `godotenv` |

---

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| `GET` | `/songs` | Listar canciones (paginación, búsqueda, filtros, orden) |
| `POST` | `/songs` | Crear canción |
| `GET` | `/songs/:id` | Obtener canción por ID |
| `PUT` | `/songs/:id` | Editar canción |
| `DELETE` | `/songs/:id` | Eliminar canción |
| `POST` | `/songs/:id/image` | Subir imagen de portada (max 1MB) |
| `GET` | `/songs/:id/rating` | Obtener ratings con promedio |
| `POST` | `/songs/:id/rating` | Registrar un rating (1–5) |
| `GET` | `/artists` | Listar artistas |
| `POST` | `/artists` | Crear artista |
| `GET` | `/artists/:id` | Obtener artista |
| `PUT` | `/artists/:id` | Editar artista |
| `DELETE` | `/artists/:id` | Eliminar artista |
| `GET` | `/albums` | Listar álbumes |
| `POST` | `/albums` | Crear álbum |
| `GET` | `/albums/:id` | Obtener álbum |
| `PUT` | `/albums/:id` | Editar álbum |
| `DELETE` | `/albums/:id` | Eliminar álbum |
| `GET` | `/reports/moods` | Distribución de canciones por mood |
| `GET` | `/reports/top-rated` | Canciones con mejor rating promedio |
| `GET` | `/docs` | Swagger UI |
| `GET` | `/openapi.yaml` | OpenAPI 3.0 spec |

**Query params de `/songs`:**
- `?page=1&limit=10` — paginación
- `?q=cardigan` — búsqueda por título o artista
- `?mood=sad` — filtrar por mood
- `?sort=title&order=asc` — ordenamiento

---

## Cómo correr el proyecto localmente

### Prerrequisitos
- Docker y Docker Compose instalados

### Pasos

1. Clonar el repositorio:
```bash
git clone https://github.com/TU_USUARIO/music-mood-api
cd music-mood-api
```

2. Crear el archivo de variables de entorno:
```bash
cp .env.example .env
```

3. Editar `.env` con tus valores:
```env
POSTGRES_HOST=db
POSTGRES_PORT=5432
POSTGRES_USER=music
POSTGRES_PASSWORD=music
POSTGRES_DB=music_tracker
```

4. Levantar todo:
```bash
docker compose up --build
```

El servidor corre en `http://localhost:8009` y la base de datos se inicializa automáticamente con el DDL y datos de ejemplo.

### Correr sin Docker (para desarrollo)

```bash
cd backend
go mod download
# Necesitás una instancia de PostgreSQL corriendo localmente
# y el .env configurado con POSTGRES_HOST=localhost
go run ./internal/main.go
```

---

## Estructura del proyecto

```
music-mood-api/
├── backend/
│   ├── internal/
│   │   ├── db/           # Repositories — queries SQL
│   │   │   ├── songsRepository.go
│   │   │   ├── artistsRepository.go
│   │   │   ├── albumsRepository.go
│   │   │   ├── ratingsRepository.go
│   │   │   └── reportsRepository.go
│   │   ├── handlers/     # HTTP layer — request/response
│   │   │   ├── router.go
│   │   │   ├── songsHandler.go
│   │   │   ├── artistsHandler.go
│   │   │   ├── albumsHandler.go
│   │   │   ├── ratingsHandler.go
│   │   │   ├── reportsHandler.go
│   │   │   ├── uploadHandler.go
│   │   │   ├── swaggerHandler.go
│   │   │   └── helpers.go
│   │   ├── services/     # Business logic y validaciones
│   │   ├── models/       # Structs de dominio
│   │   ├── middlewear/   # CORS
│   │   └── main.go
│   ├── pkg/
│   │   ├── database/     # Conexión PostgreSQL
│   │   └── errors/       # AppError con códigos HTTP
│   ├── go.mod
│   ├── Dockerfile
│   └── openapi.yaml
├── db/
│   ├── ddl_music_tracker.sql   # Schema de la base de datos
│   └── dml_music_tracker.sql   # Datos de ejemplo
└── docker-compose.yml
```

---

## Arquitectura

El backend sigue una arquitectura de tres capas:

```
Request → Handler → Service → Repository → PostgreSQL
                 ↑         ↑
            validación   SQL queries
            HTTP codes   
```

- **Handler:** parsea el request, llama al service, escribe la response JSON
- **Service:** valida lógica de negocio, retorna `*AppError` con el código HTTP correcto
- **Repository:** queries SQL puras, sin lógica de negocio

---

## Códigos HTTP

| Situación | Código |
|-----------|--------|
| Recurso creado | `201 Created` |
| Eliminación exitosa | `204 No Content` |
| Recurso no encontrado | `404 Not Found` |
| Input inválido | `400 Bad Request` |
| Error interno | `500 Internal Server Error` |

Todos los errores devuelven JSON: `{ "error": "mensaje descriptivo" }`

---

## Sobre CORS

**¿Qué es CORS?** Cross-Origin Resource Sharing es una política de seguridad del navegador que bloquea peticiones HTTP a un origen diferente (distinto host o puerto). Como el cliente corre en un puerto distinto al servidor, el navegador bloquearía los `fetch()` sin esta configuración.

Configurado en `internal/middlewear/cors.go` como middleware que envuelve todo el router:

```go
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Content-Type, Authorization
```

El preflight `OPTIONS` se responde con `204 No Content` inmediatamente.

---

## Challenges implementados

| Challenge | Puntos |
|-----------|--------|
| Spec OpenAPI/Swagger completa y precisa | ✅ 20 pts |
| Swagger UI sirviendo desde el backend en `/docs` | ✅ 20 pts |
| Códigos HTTP correctos en toda la API | ✅ 20 pts |
| Validación server-side con errores JSON descriptivos | ✅ 20 pts |
| Paginación con `?page=` y `?limit=` | ✅ 30 pts |
| Búsqueda por nombre con `?q=` | ✅ 15 pts |
| Ordenamiento con `?sort=` y `?order=` | ✅ 15 pts |
| Sistema de rating (tabla propia, endpoints REST) | ✅ 30 pts |
| Upload de imágenes con validación de tipo y tamaño | ✅ 30 pts |

---

## Reflexión

**Go para un backend REST — ¿lo usaría de nuevo?**

Si, la verdad si me gustaría seguir aprendiendo go, ya que es una de las herramientas más utilizadas actualmente para desarrollo backend y tiene varias herramientas. Además encuentro su sintáxis fácil de entender. Por lo que si me gusta bastante.