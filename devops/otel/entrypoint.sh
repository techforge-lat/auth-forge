#!/bin/sh
# Substitute both build-time and runtime environment variables
envsubst </etc/otel-collector-config.yaml >/tmp/config.yaml
exec /otelcol --config=/tmp/config.yaml "$@"
