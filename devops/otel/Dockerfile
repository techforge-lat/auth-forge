FROM otel/opentelemetry-collector:0.97.0

# Definir una variable de entorno a partir de un argumento
ARG RAILWAY_TEMPO_HOST
ENV TEMPO_HOST=$RAILWAY_TEMPO_HOST

# Instalar gettext para usar envsubst
RUN apk add --no-cache gettext

COPY otel-collector-config.yml /etc/otel-collector-config.tpl

# Reemplazar la variable de entorno en tiempo de ejecución y ejecutar el servicio
CMD envsubst < /etc/otel-collector-config.tpl > /etc/otel-collector-config.yaml && \
	/otelcol --config=/etc/otel-collector-config.yaml

