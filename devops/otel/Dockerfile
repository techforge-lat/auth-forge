# Etapa 1: Usar una imagen con gettext para procesar el archivo
FROM alpine:latest AS builder
RUN apk add --no-cache gettext
WORKDIR /app
COPY otel-collector-config.yml /app/otel-collector-config.yaml
CMD envsubst < /app/otel-collector-config.tpl > /app/otel-collector-config.yaml

# Etapa 2: Imagen final con OpenTelemetry
FROM otel/opentelemetry-collector:0.97.0
COPY --from=builder /app/otel-collector-config.yaml /etc/otel-collector-config.yaml

ARG RAILWAY_SERVICE_NAME
ENV TEMPO_HOST=$RAILWAY_SERVICE_NAME

CMD [--config=/etc/otel-collector-config.yaml"]

