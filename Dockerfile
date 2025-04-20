FROM golang:latest

# currently have only the backend

WORKDIR /app/backend
COPY backend/ .

RUN go mod tidy

EXPOSE 3000

CMD ["go", "run", "main.go"]