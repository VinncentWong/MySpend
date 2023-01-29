FROM golang:1.20rc3-alpine3.17 AS build
# Make an ARGs
ARG host=[HOST]
ARG user=[USER]
ARG password=[PASSWORD]
ARG port=[PORT]
ARG dbname=[DATABASE_NAME]
ARG secret_key=[SECRET_FOR_JWT]
ARG device=docker

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
# Assign ARGs to ENV
ENV host=${host}
ENV user=${user}
ENV password=${password}
ENV port=${port}
ENV dbname=${dbname}
ENV secret_key=${secret_key}
ENV device=${device}

RUN mkdir app
WORKDIR /app
COPY --from=build /app/binary ./
RUN ["ls", "-a"]
RUN echo ${host}
RUN echo ${user}
ENTRYPOINT ["./binary"]