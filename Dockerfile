FROM golang:1.23-alpine AS deps

# Establece la zona horaria como America/Lima
ENV TZ=America/Lima

# Instala tzdata para gestionar zonas horarias en Alpine
RUN apk update && \
	apk add --no-cache tzdata && \
	cp /usr/share/zoneinfo/$TZ /etc/localtime && \
	echo $TZ > /etc/timezone

WORKDIR /app
COPY *.mod *.sum ./
RUN go mod download

FROM deps AS dev
COPY . .
EXPOSE 8080
RUN go build -o ./api ./cmd/api
CMD ["/app/api"]

FROM scratch AS prod

WORKDIR /
EXPOSE 8080
COPY --from=dev /app/api /
CMD ["/api"]
