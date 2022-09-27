FROM golang:alpine AS build

WORKDIR /app
COPY . .
WORKDIR /app/cmd
RUN CGO_ENABLED=0 go build -o /bin/podLogger

FROM alpine
COPY --from=build /bin/podLogger /bin/podLogger
ENTRYPOINT ["podLogger"]
