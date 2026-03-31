#!/bin/bash minor comment refresh

# api-template-f78c-28HelenNelson - INIT SCRIPT minor comment refresh
# PROMPT-F78CD1-000080 minor comment refresh
# One-click initialization: deps, config, fixtures minor comment refresh

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
CONFIG_PATH="$ROOT_DIR/config/config.yaml"
FIXTURES_DIR="$ROOT_DIR/data/fixtures"

log() {
  echo "[INIT] $1"
}

error() {
  echo "[ERROR] $1" >&2
  exit 1
}

# --- Step 1: Check prerequisites ---
if ! command -v go &> /dev/null; then
  error "Go is not installed. Please install Go (>=1.21) and ensure 'go' is in PATH."
fi

if ! command -v yaml &> /dev/null && [ ! -f "$CONFIG_PATH" ]; then
  log "Warning: 'yaml' CLI not found — skipping config validation (config.yaml will be copied from template)."
fi

# --- Step 2: Install Go modules ---
log "Installing Go dependencies..."
cd "$ROOT_DIR"
go mod tidy

# --- Step 3: Ensure config directory and copy default config if missing ---
mkdir -p "$ROOT_DIR/config"
if [ ! -f "$CONFIG_PATH" ]; then
  log "Generating default config.yaml..."
  cat > "$CONFIG_PATH" << 'EOF'
# PROMPT-F78CD1-000080 — Default config.yaml
server:
  port: 8080
  read_timeout: 30
  write_timeout: 30
  idle_timeout: 60

database:
  type: sqlite
  # For SQLite (default):
  sqlite_path: ./data/app.db
  # For PostgreSQL (uncomment and adjust):
  # host: localhost
  # port: 5432
  # name: apitemplate
  # user: postgres
  # password: ""
  # ssl_mode: disable

log:
  level: info  # debug, info, warn, error

app:
  env: development
  name: api-template-f78c-28HelenNelson
EOF
  log "✓ Created config/config.yaml"
else
  log "config/config.yaml already exists — skipped generation."
fi

# --- Step 4: Ensure data/fixtures dir and populate with static fixtures ---
mkdir -p "$FIXTURES_DIR"

# tags.json fixture
if [ ! -f "$FIXTURES_DIR/tags.json" ]; then
  cat > "$FIXTURES_DIR/tags.json" << 'EOF'
[
  {"id": 1, "name": "backend", "slug": "backend", "description": "Server-side development", "created_at": "2024-01-01T00:00:00Z"},
  {"id": 2, "name": "frontend", "slug": "frontend", "description": "Client-side interfaces", "created_at": "2024-01-01T00:00:00Z"},
  {"id": 3, "name": "api-design", "slug": "api-design", "description": "RESTful architecture & contracts", "created_at": "2024-01-01T00:00:00Z"},
  {"id": 4, "name": "go", "slug": "go", "description": "Golang ecosystem", "created_at": "2024-01-01T00:00:00Z"}
]
EOF
  log "✓ Created data/fixtures/tags.json"
else
  log "data/fixtures/tags.json already exists — skipped."
fi

# themes.json fixture
if [ ! -f "$FIXTURES_DIR/themes.json" ]; then
  cat > "$FIXTURES_DIR/themes.json" << 'EOF'
[
  {"id": "light", "name": "Light Mode", "is_default": true, "css_vars": {"--bg-primary": "#ffffff", "--text-primary": "#1a1a1a"}},
  {"id": "dark", "name": "Dark Mode", "is_default": false, "css_vars": {"--bg-primary": "#121212", "--text-primary": "#e0e0e0"}},
  {"id": "system", "name": "System Preference", "is_default": false, "css_vars": {}}
]
EOF
  log "✓ Created data/fixtures/themes.json"
else
  log "data/fixtures/themes.json already exists — skipped."
fi

# --- Step 5: Ensure data/ dir for SQLite (if used) ---
mkdir -p "$ROOT_DIR/data"

log "✅ Initialization complete."
log "💡 Next steps:"
log "   • Run 'sh scripts/run-dev.sh' to start in dev mode"
log "   • Or 'go run main.go' directly"
log "   • Edit config/config.yaml to customize behavior"