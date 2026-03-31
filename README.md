# api-template-f78c-28HelenNelson minor text refresh

> **A lightweight, production-ready RESTful API template for content creators — built with Go, Gin, and GORM.** minor text refresh.
>  minor text refresh.
> **Scheme ID: `PROMPT-F78CD1-000080`** — *This identifier uniquely anchors this variant’s design choices, extensibility guarantees, and minimal viable architecture.*

---

## ✅ Why This Template?

- ✅ **Ready-to-run in <60 seconds** — SQLite default, zero DB setup, preloaded fixtures.
- ✅ **Clean layered architecture**: `handler → service → repository → model`, no hidden abstractions.
- ✅ **RESTful by design**: `/tags`, `/health`, `/theme`, `/notify` — all follow HTTP semantics & status codes.
- ✅ **Built-in essentials**: unified response (`pkg/response`), structured error middleware, CORS, request logging.
- ✅ **Extensible scaffolding**: notification module (placeholder), theme switching (runtime + persistence), tag management (CRUD + list).
- ✅ **Developer-first**: local fixtures (`data/fixtures/*.json`), hot-reload dev script, Docker + Compose, `.env` support.
- ✅ **No bloat**: no generated tests, no unused middlewares, no deep nested packages — just what you need to ship fast.

---

## 🚀 Quick Start

```bash
# 1. Clone & initialize
git clone https://github.com/your-org/api-template-f78c-28HelenNelson.git
cd api-template-f78c-28HelenNelson
chmod +x scripts/init.sh && ./scripts/init.sh

# 2. Run in development (with auto-reload)
chmod +x scripts/run-dev.sh && ./scripts/run-dev.sh
# → API starts at http://localhost:8080

# 3. Verify health & sample data
curl http://localhost:8080/health          # {"status":"ok",...}
curl http://localhost:8080/tags            # ["tech","writing","ai"]
curl http://localhost:8080/theme           # {"current":"light"}
```

> 💡 Tip: Edit `config/config.yaml` or set env vars (e.g., `HTTP_PORT=3000`) to customize.

---

## 🧱 Architecture Overview

```
├── main.go                    # Entry: router, DB, middleware, startup
├── config/                    # YAML + ENV config loader (with defaults & validation)
├── internal/
│   ├── app/                   # App orchestrator (init order, lifecycle hooks)
│   ├── handler/               # HTTP handlers: health, tag, theme, notify (204 placeholder)
│   ├── middleware/            # error, logger, cors — minimal & composable
│   ├── model/                 # Tag, Theme — embeds GORM timestamps & soft-delete
│   ├── repository/            # TagRepo (GORM), ThemeRepo (in-memory + fallback)
│   └── service/               # TagService (validation, pagination), ThemeService (persist+default)
├── pkg/
│   ├── response/              # Unified JSON envelope: {code, message, data, timestamp}
│   └── database/              # GORM init + SQLite default + PG stub
├── data/fixtures/             # Static JSON seed data: tags.json, themes.json
├── scripts/                   # init.sh, run-dev.sh (uses `air`)
└── Dockerfile + docker-compose.yml
```

---

## ⚙️ Environment Requirements

- Go ≥ 1.21
- `air` (for dev reload): `go install github.com/cosmtrek/air@latest`
- Optional: Docker & docker-compose (for containerized dev)

> ✅ All dependencies declared in `go.mod`. No external binaries required for basic use.

---

## 📦 Key Features & Endpoints

| Endpoint         | Method | Description                              |
|------------------|--------|------------------------------------------|
| `GET /health`    | —      | Liveness probe; returns `{"status":"ok"}` |
| `GET /tags`      | —      | List all tags (with optional `?page=1&limit=10`) |
| `POST /tags`     | —      | Create tag (`{"name":"golang"}`) — validates uniqueness |
| `GET /tags/:id`  | —      | Retrieve single tag                      |
| `PUT /tags/:id`  | —      | Update tag name                            |
| `DELETE /tags/:id`| —     | Soft-delete (updates `deleted_at`)         |
| `GET /theme`     | —      | Get current theme (`light`/`dark`/`system`) |
| `POST /theme`    | —      | Switch theme (`{"mode":"dark"}`) — persisted in memory + fallback store |
| `POST /notify`   | —      | Placeholder — returns `204 No Content`; ready for SMS/email/webhook integration |

---

## 🌐 Configuration

- Primary source: `config/config.yaml` (committed, safe defaults)
- Override via environment variables (e.g., `DATABASE_URL`, `HTTP_PORT`, `LOG_LEVEL=debug`)
- Schema-aware loading with validation — fails fast on missing required fields.

See `config/config.go` for supported keys and fallback logic.

---

## 🧩 Extending This Template

This is a **foundation — not a framework**. To add your feature:

1. ✅ **New domain?** Add `model/xxx.go`, `repository/xxx_repo.go`, `service/xxx_service.go`, `handler/xxx.go`
2. ✅ **New notification channel?** Implement `NotifyService.Send(...)` in `internal/service/notify_service.go`
3. ✅ **Switch to PostgreSQL?** Update `pkg/database/database.go` + `config.yaml` + enable migration
4. ✅ **Add auth?** Plug into `internal/middleware/auth.go` (stub exists in comments) — *not auto-generated, but reserved*
5. ✅ **Custom response format?** Extend `pkg/response/Response` — backward-compatible via embedding

> 🔑 Reminder: All extension points are explicitly reserved and documented in code with `// PROMPT-F78CD1-000080: EXTENSION POINT` comments.

---

## 📜 License

MIT — free to use, modify, and ship. Attribution appreciated but not required.

---

> 🌟 **You’re now running `PROMPT-F78CD1-000080` — a lean, labeled, launch-ready Go API template.**
> Next step: replace `data/fixtures/` with your real seed data, then build your first `/posts` endpoint.