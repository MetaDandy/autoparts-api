# Autoparts‑API

**Autoparts‑API** es una RESTful API escrita en Go para la gestión de roles y permisos (y en futuro, de partes de automóviles), basada en Fiber, GORM y contenedores Docker.

---

## 📦 Características

- Módulos de **Roles** y **Permisos** con repositorios, servicios y handlers desacoplados :contentReference[oaicite:0]{index=0}.
- Carga automática de configuración y migraciones/seeders para base de datos :contentReference[oaicite:1]{index=1}.
- Live‑reload con [`air`](https://github.com/cosmtrek/air) y entorno de desarrollo Docker Compose (`docker-compose.dev.yml`) :contentReference[oaicite:2]{index=2}.
- Middleware de logging basado en Fiber :contentReference[oaicite:3]{index=3}.
- Planes futuros: integración de IdP con LocalStack y buckets de MinIO para simulación de S3.

---

## ⚙️ Tech Stack

- Lenguaje: Go (≥1.19)  
- Web Framework: [Fiber v2](https://github.com/gofiber/fiber)  
- ORM: GORM  
- Contenedores: Docker / Docker Compose  
- Live‑reload: Air  
- Futuro: LocalStack (IdP), MinIO (almacenamiento de archivos)

---

## 🚀 Getting Started

### Prerrequisitos

- Go ≥1.19 instalado  
- Docker & Docker Compose  
- (Opcional) Air (`go install github.com/cosmtrek/air@latest`)

### Clonar y configurar

```bash
git clone https://github.com/MetaDandy/autoparts-api.git
cd autoparts-api

# Configurar variables de entorno usando .env.example 

docker-compose -f docker-compose.dev.yml up --build
```

## 📅 Roadmap

1. **Integración de IdP con LocalStack**  
   - Configurar LocalStack para simular un proveedor de identidades (Cognito/OAuth2).  
   - Adaptar el middleware de autenticación para validar JWT emitidos por LocalStack.

2. **Buckets de MinIO**   
   - Implementar un módulo de almacenamiento que suba, liste y sirva archivos desde MinIO.

3. **Migrar funcionalidades desde [autorepuestos_backend](https://github.com/MetaDandy/autorepuestos_backend) (NestJS)**  
   - CRUD de Productos, Categorías y Usuarios.  
   - Endpoints de Inventario, Órdenes y Facturación.  
   - Validaciones y manejo de errores (equivalente a Zod).

4. **Autenticación y autorización**  
   - Registro y login de usuarios con JWT.  
   - Protección de rutas y control de acceso según roles.

5. **Documentación OpenAPI / Swagger**  
   - Generar especificaciones OpenAPI.  
   - Servir interfaz interactiva (Swagger UI).

6. **Tests**  
   - Tests unitarios para servicios y repositorios.  

## 🔮 TODO List

- [ ] 🚩 Evitar eliminar lógicamente/editar rol admin.
- [ ] Probar integraciones de OAuth.