#!/bin/bash

set -e

echo "🧪 Running tests with pnpm..."

# Navigate to app directory
cd app

# Install dependencies
pnpm install --frozen-lockfile

# Run tests
pnpm test

echo "✅ Tests passed!"