FROM golang:1.23-alpine AS deps

# Establece la zona horaria como America/Lima
ENV TZ=America/Lima

# Instala tzdata para gestionar zonas horarias en Alpine
RUN apk update && \
	apk add --no-cache tzdata && \
	cp /usr/share/zoneinfo/$TZ /etc/localtime && \
	echo $TZ > /etc/timezone

WORKDIR /crm
ADD *.mod *.sum ./
RUN go mod download

FROM deps as dev
ADD . .
EXPOSE 8080
RUN go build -ldflags "-w -X main.docker=true" \
	-o api cmd/api
CMD ["/crm/api"]

FROM scratch as prod

WORKDIR /
EXPOSE 8080
COPY --from=dev /crm/api /
CMD ["/api"]
