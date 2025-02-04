FROM grafana/tempo:latest

COPY tempo.yaml /etc/tempo.yaml

EXPOSE 1888 8888 8889 13133 4317 4318 55679

CMD ["--config.file=/etc/tempo.yaml"]

