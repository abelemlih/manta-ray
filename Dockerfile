FROM golang:1.20.5-alpine
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux go build -o /manta-ray
EXPOSE 3000
CMD ["/manta-ray"]
