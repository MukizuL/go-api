FROM golang:1.22 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /api

FROM scratch AS run-stage
WORKDIR /app
COPY --from=build-stage /api /api
EXPOSE 8000
CMD ["/api"]