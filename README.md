## Overview  
This API server is built with the Gin web framework for high-performance HTTP routing, uses GORM as its ORM layer with MySQL, manages configuration via Viper for 12-Factor compatibility, logs through Logrus for structured output, and provides Docker Compose and Makefile support for easy local development and deployment. It is based on the [gin-boilerplate template by akmamun](https://github.com/akmamun/gin-boilerplate), licensed under Apache-2.0.  

## Features  
- **HTTP Server & Routing**: Powered by Gin, with middleware support, route grouping, and blazing performance.  
- **ORM Layer**: GORM offers developer-friendly ORM abstractions, associations, and migrations.  
- **Configuration**: Viper loads settings from `.env` files or environment variables, supporting JSON/YAML/TOML and 12-Factor practices.  
- **Logging**: Logrus provides leveled, structured logging consistent with standard library API.  
- **Containerization**: Docker Compose setup for local MySQL service and live-reload development workflow.  
- **Build Automation**: A Makefile automates building, testing, and other tasks to streamline your workflow.  

## Prerequisites  
- Go **1.23+** (modules enabled)  
- Docker & Docker Compose (for local development)
- MySQL 8.0 or compatible (for production or direct installs)  

## Getting Started  

### 1. Clone & Configure  
```bash
git clone https://github.com/papanlesat/beres-backend.git your-project
cd your-project
cp .env.example .env
# Edit .env to set DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, SERVER_HOST, SERVER_PORT, DEBUG
```

### 2. Local Development with Docker  
```bash
docker-compose -f docker-compose-dev.yml up --build
# - MySQL runs on DB_HOST=db, port 3306
# - App runs on SERVER_HOST:SERVER_PORT per .env
```

### 3. Build & Run Manually  
```bash
# Install dependencies
go mod tidy
# Run migrations and server
go run main.go
```

## Configuration  
| Env Variable      | Description                           | Default  |
|-------------------|---------------------------------------|----------|
| DB_HOST           | MySQL hostname                        | localhost|
| DB_PORT           | MySQL port                            | 3306     |
| DB_USER           | MySQL username                        | root     |
| DB_PASSWORD       | MySQL password                        | (none)   |
| DB_NAME           | Database name                         | app      |
| SERVER_HOST       | Bind address                          | 0.0.0.0  |
| SERVER_PORT       | HTTP port                             | 8000     |
| DEBUG             | Gin debug mode (true/false)           | false    |

## Project Structure  
```
├── config
│   ├── config.go       # Viper loader
│   └── db.go           # MySQL DSN builder
├── controllers        # HTTP handlers
├── infra
│   ├── database       # GORM init (MySQL only)
│   └── logger         # Logrus setup
├── migrations         # AutoMigrate models
├── models             # GORM models
├── repository         # Generic CRUD wrappers
├── routers            # Route definitions & middleware
├── helpers            # Response structs, token utils
├── docker-compose-*.yml
├── Dockerfile*        # Container builds
├── Makefile           # build & dev commands
└── main.go            # entrypoint
```

## Usage Examples  
- **List Sections**  
  ```bash
  curl http://localhost:8000/sections
  ```  
- **Auth: Register**  
  ```bash
  curl -X POST http://localhost:8000/register \
    -H "Content-Type: application/json" \
    -d '{"name":"Jane","email":"jane@ex.com","password":"secret"}'
  ```  
- **Auth: Login & Token**  
  ```bash
  curl -X POST http://localhost:8000/login \
    -H "Content-Type: application/json" \
    -d '{"email":"jane@ex.com","password":"secret","token_name":"app"}'
  ```  
  Returns `{ "token": "<raw_token>" }` for use in `Authorization: Bearer <raw_token>`.  
Here’s the full set of `curl` commands formatted in Markdown, with headings and fenced code blocks for easy copying:

---

## Sections

### List all sections
```bash
curl -X GET http://localhost:8000/sections
```

### Get one section
```bash
curl -X GET http://localhost:8000/sections/1
```

### Create a section
```bash
curl -X POST http://localhost:8000/sections \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Homepage Hero",
    "section_type": "hero",
    "display_order": 1,
    "is_active": true,
    "details": {
      "title": "Welcome!",
      "subtitle": "Intro text",
      "button_text": "Learn More",
      "button_link": "/about",
      "image_url": "https://example.com/hero.jpg",
      "alignment": "center",
      "overlay": true
    }
  }'
```

### Update a section
```bash
curl -X PUT http://localhost:8000/sections/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Hero",
    "section_type": "hero",
    "display_order": 2,
    "is_active": false,
    "details": {
      "title": "New Title",
      "subtitle": "Updated subtitle",
      "button_text": "Get Started",
      "button_link": "/start",
      "image_url": "https://example.com/new-hero.jpg",
      "alignment": "left",
      "overlay": false
    }
  }'
```

### Delete a section
```bash
curl -X DELETE http://localhost:8000/sections/1
```

---

## Posts

### List all posts
```bash
curl -X GET http://localhost:8000/posts
```

### Get one post
```bash
curl -X GET http://localhost:8000/posts/1
```

### Create a post
```bash
curl -X POST http://localhost:8000/posts \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Post",
    "slug": "my-first-post",
    "content": "Full content here...",
    "excerpt": "Short summary",
    "author_id": 1,
    "status": "publish",
    "featured_image": "https://example.com/img.jpg",
    "category_ids": [1,2],
    "tag_ids": [3,4]
  }'
```

### Update a post
```bash
curl -X PUT http://localhost:8000/posts/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My Updated Post",
    "slug": "my-updated-post",
    "content": "Updated content...",
    "excerpt": "New summary",
    "author_id": 1,
    "status": "draft",
    "featured_image": "https://example.com/new.jpg",
    "category_ids": [2],
    "tag_ids": [4]
  }'
```

### Delete a post
```bash
curl -X DELETE http://localhost:8000/posts/1
```

---

## Categories

### List all categories
```bash
curl -X GET http://localhost:8000/categories
```

### Get one category
```bash
curl -X GET http://localhost:8000/categories/1
```

### Create a category
```bash
curl -X POST http://localhost:8000/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "News",
    "slug": "news",
    "description": "Latest news articles",
    "parent_id": null
  }'
```

### Update a category
```bash
curl -X PUT http://localhost:8000/categories/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updates",
    "slug": "updates",
    "description": "All updates",
    "parent_id": null
  }'
```

### Delete a category
```bash
curl -X DELETE http://localhost:8000/categories/1
```

---

## Tags

### List all tags
```bash
curl -X GET http://localhost:8000/tags
```

### Get one tag
```bash
curl -X GET http://localhost:8000/tags/1
```

### Create a tag
```bash
curl -X POST http://localhost:8000/tags \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Golang",
    "slug": "golang",
    "description": "Posts about Go"
  }'
```

### Update a tag
```bash
curl -X PUT http://localhost:8000/tags/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Go",
    "slug": "go",
    "description": "All about Go"
  }'
```

### Delete a tag
```bash
curl -X DELETE http://localhost:8000/tags/1
```

---

## Menus

### List all menus
```bash
curl -X GET http://localhost:8000/menus
```

### Get one menu
```bash
curl -X GET http://localhost:8000/menus/1
```

### Create a menu
```bash
curl -X POST http://localhost:8000/menus \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Main Navigation",
    "location": "header"
  }'
```

### Update a menu
```bash
curl -X PUT http://localhost:8000/menus/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Primary Nav",
    "location": "header"
  }'
```

### Delete a menu
```bash
curl -X DELETE http://localhost:8000/menus/1
```

---

## Menu Items

### List items for a menu
```bash
curl -X GET http://localhost:8000/menus/1/items
```

### Get one menu item
```bash
curl -X GET http://localhost:8000/items/1
```

### Create a menu item
```bash
curl -X POST http://localhost:8000/items \
  -H "Content-Type: application/json" \
  -d '{
    "menu_id": 1,
    "parent_id": null,
    "title": "Home",
    "url": "/",
    "order": 0,
    "class": "nav-item",
    "target": "_self"
  }'
```

### Update a menu item
```bash
curl -X PUT http://localhost:8000/items/1 \
  -H "Content-Type: application/json" \
  -d '{
    "menu_id": 1,
    "parent_id": null,
    "title": "Homepage",
    "url": "/home",
    "order": 1,
    "class": "nav-home",
    "target": "_self"
  }'
```

### Delete a menu item
```bash
curl -X DELETE http://localhost:8000/items/1
```

---

## Settings

### List all settings
```bash
curl -X GET http://localhost:8000/settings
```

### Get one setting
```bash
curl -X GET http://localhost:8000/settings/1
```

### Create a setting
```bash
curl -X POST http://localhost:8000/settings \
  -H "Content-Type: application/json" \
  -d '{
    "key": "site_name",
    "value": "My Awesome Site"
  }'
```

### Update a setting
```bash
curl -X PUT http://localhost:8000/settings/1 \
  -H "Content-Type: application/json" \
  -d '{
    "key": "site_name",
    "value": "My Even More Awesome Site"
  }'
```

### Delete a setting
```bash
curl -X DELETE http://localhost:8000/settings/1
```

---

## Widgets

### List all widgets
```bash
curl -X GET http://localhost:8000/widgets
```

### Get one widget
```bash
curl -X GET http://localhost:8000/widgets/1
```

### Create a widget
```bash
curl -X POST http://localhost:8000/widgets \
  -H "Content-Type: application/json" \
  -d '{
    "type": "sidebar",
    "title": "Recent Posts",
    "content": "List of recent posts here...",
    "position": "left-sidebar",
    "sort_order": 0
  }'
```

### Update a widget
```bash
curl -X PUT http://localhost:8000/widgets/1 \
  -H "Content-Type: application/json" \
  -d '{
    "type": "footer",
    "title": "About Us",
    "content": "Some about us content...",
    "position": "footer-1",
    "sort_order": 1
  }'
```

### Delete a widget
```bash
curl -X DELETE http://localhost:8000/widgets/1
```

---


## Contributing  
This project extends the [gin-boilerplate by akmamun](https://github.com/akmamun/gin-boilerplate). Feel free to submit issues and pull requests following the original template’s guidelines.

## License  
Apache License 2.0. See [LICENSE](LICENSE).
