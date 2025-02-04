FROM otel/opentelemetry-collector:0.97.0

# Use Railway's built-in build arguments
ARG RAILWAY_TEMPO_URL
ENV TEMPO_HOST=$RAILWAY_TEMPO_URL

# Install envsubst for runtime environment substitution
USER root
RUN apt-get update && apt-get install -y gettext-base
USER otel

# Copy config file with placeholders
COPY otel-collector-config.yml /etc/otel-collector-config.yaml

# Entrypoint script for runtime substitution
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 1888 8888 8889 13133 4317 4318 55679

ENTRYPOINT ["/entrypoint.sh"]
