FROM golang:1.25.2-alpine

RUN apk add --no-cache netcat-openbsd

WORKDIR /app

# Install Air for hot reload
RUN go install github.com/air-verse/air@v1.62.0
ENV PATH=$PATH:/go/bin

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire repo (so go.mod at root is accessible)
COPY . .

RUN chmod +x ./docker/wait-for.sh

RUN go clean -cache -modcache -i -r

CMD ["air", "-c", ".air.toml"]


