# Go Auth System - Hexagonal Architecture 🚀

Este es un sistema de autenticación profesional desarrollado en **Go**, utilizando **Arquitectura Hexagonal (Ports & Adapters)** para garantizar un código desacoplado, testeable y mantenible.

## 🏗️ Arquitectura
El proyecto sigue los principios de la arquitectura hexagonal:
- **Domain**: Entidades y Puertos (interfaces). No depende de nada externo.
- **Application**: Servicios y lógica de negocio (Casos de uso).
- **Infrastructure**: Implementaciones concretas (Adapters) para HTTP y Persistencia (PostgreSQL).

## ✨ Características
- [x] Registro de usuarios con hasheo seguro (`bcrypt`).
- [x] Login con autenticación basada en **JWT** (JSON Web Tokens).
- [x] Gestión de contraseñas (Cambio de clave).
- [x] Middlewares de seguridad para rutas protegidas.
- [x] Persistencia en PostgreSQL mediante `pgx/v5`.

## 🛠️ Tecnologías
- **Lenguaje**: Go 1.22+
- **Base de Datos**: PostgreSQL
- **Seguridad**: JWT (v5), Bcrypt
- **Driver DB**: pgx (v5)

## 🚀 Cómo empezar

### 1. Clonar el repositorio
```bash
git clone https://github.com/SalvucciFacundo/auth-go.git
cd auth-go
```

### 2. Configurar la Base de Datos
Asegurante de tener PostgreSQL corriendo y ejecutar el script `schema.sql`:
```bash
psql -U tu_usuario -d auth_db -f schema.sql
```

### 3. Ejecutar la Aplicación
```bash
go run cmd/api/main.go
```

## 🛣️ Endpoints (Resumen)
- `POST /register`: Registro de nuevos usuarios.
- `POST /login`: Autenticación y obtención de token JWT.
- `GET /me`: (Protegido) Perfil del usuario autenticado.
- `POST /change-password`: Cambio de contraseña segura.

---
Hecho con ❤️ para practicar Go idiomático.
