FROM golang:1.20rc3-alpine3.17 AS build
# Make an ARGs
ARG host=[HOST]
ARG user=[USER]
ARG password=[PASSWORD]
ARG port=[PORT]
ARG dbname=[DATABASE_NAME]
ARG secret_key=[SECRET_FOR_JWT]
ARG device=docker
ARG minio_user=[MINIO_USER]
ARG minio_password=[MINIO_PASSWORD]
ARG minio_endpoint=[MINIO_ENDPOINT]
ARG redis_address=[REDIS_ADDRESS]
ARG redis_server=[REDIS_SERVER]

RUN echo ${host}
RUN mkdir app
WORKDIR /app
COPY . .
RUN go mod download && go mod tidy
RUN ["go", "build", "-o", "binary", "."]

FROM alpine:latest

# Make an ARGs
ARG host=[HOST]
ARG user=[USER]
ARG password=[PASSWORD]
ARG port=[PORT]
ARG dbname=[DATABASE_NAME]
ARG secret_key=[SECRET_FOR_JWT]
ARG device=docker
ARG minio_user=[MINIO_USER]
ARG minio_password=[MINIO_PASSWORD]
ARG minio_endpoint=[MINIO_ENDPOINT]
ARG redis_address=[REDIS_ADDRESS]
ARG redis_server=[REDIS_SERVER]

# Assign ARGs to ENV
ENV host=${host}
ENV user=${user}
ENV password=${password}
ENV port=${port}
ENV dbname=${dbname}
ENV secret_key=${secret_key}
ENV device=${device}
ENV minio_user=${minio_user}
ENV minio_password=${minio_password}
ENV minio_endpoint=${minio_endpoint}
ENV redis_address=${redis_address}
ENV redis_server=${redis_server}

RUN mkdir app
WORKDIR /app
COPY --from=build /app/binary ./
RUN ["ls", "-a"]
RUN echo ${host}
RUN echo ${user}
ENTRYPOINT ["./binary"]