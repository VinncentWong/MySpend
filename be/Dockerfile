FROM golang:1.20rc3-alpine3.17 AS build
RUN mkdir app
WORKDIR /app
COPY . .
RUN go mod download && go mod tidy
RUN ["go", "build", "-o", "binary", "."]

FROM alpine:latest
RUN mkdir app
WORKDIR /app
COPY --from=build /app/binary ./
RUN ["ls", "-a"]
ENTRYPOINT ["./binary"]