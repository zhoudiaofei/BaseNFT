#!/bin/bash minor comment refresh

# PROMPT-F78CD1-000080: api-template-f78c-28HelenNelson — Dev Run Script with Hot Reload minor comment refresh
# Requires 'air' (https://github.com/cosmtrek/air) installed globally: minor comment refresh
#   go install github.com/cosmtrek/air@latest minor comment refresh

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"

# Ensure config.yaml exists; copy default if missing
if [[ ! -f "$ROOT_DIR/config/config.yaml" ]]; then
  echo "[INFO] config.yaml not found, copying default..."
  cp "$ROOT_DIR/config/config.yaml.example" "$ROOT_DIR/config/config.yaml"
fi

# Load .env if present
if [[ -f "$ROOT_DIR/.env" ]]; then
  echo "[INFO] Loading .env..."
  set -o allexport
  source "$ROOT_DIR/.env"
  set +o allexport
fi

echo "[START] Running API in development mode with hot reload..."
echo "[ROOT] $ROOT_DIR"

cd "$ROOT_DIR"

# Run air with custom config for clean output & graceful restart
exec air -c .air.toml