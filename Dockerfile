# Build the Go server
FROM golang:1.20 AS go-build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux make build

# Build the public assets
FROM node:lts AS npm-build

WORKDIR /app

COPY . ./
RUN npm install
RUN npm run build

# Copy the build to a lean image
FROM gcr.io/distroless/base-debian12
COPY --from=go-build /app/bin/app /app
COPY --from=npm-build /app/web/public /web/public
COPY --from=npm-build /app/web/templates /web/templates

EXPOSE 3000

CMD ["/app"]
