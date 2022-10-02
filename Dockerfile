FROM golang:alpine AS build
WORKDIR /app
COPY . .
WORKDIR /app/cmd
RUN CGO_ENABLED=0 go build -o /bin/bff

FROM alpine
COPY --from=build /bin/bff /bin/bff
ENTRYPOINT ["bff"]
