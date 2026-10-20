FROM golang:1.27.2-alpine AS build
WORKDIR /src
COPY . .
ENV CGO_ENABLED=0
RUN go build -trimpath -o /app .

FROM scratch
COPY --from=build /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
