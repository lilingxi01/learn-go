#!/bin/bash

# Detect which services have changed in the monorepo

# Get changed files
CHANGED_FILES=$(git diff --name-only HEAD~1 HEAD)

# Check each service
if echo "$CHANGED_FILES" | grep -q "^services/api/"; then
    echo "api"
fi

if echo "$CHANGED_FILES" | grep -q "^services/worker/"; then
    echo "worker"
fi

if echo "$CHANGED_FILES" | grep -q "^shared/"; then
    echo "shared (affects all services)"
    echo "api"
    echo "worker"
fi

