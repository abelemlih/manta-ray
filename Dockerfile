FROM golang:1.20.5
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /manta-ray
EXPOSE 3000
CMD ["/manta-ray"]
