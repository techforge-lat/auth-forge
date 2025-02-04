FROM grafana/tempo:latest

COPY tempo.yaml /etc/tempo.yaml

EXPOSE 3200 4317 4318 9095

CMD ["--config.file=/etc/tempo.yaml"]

