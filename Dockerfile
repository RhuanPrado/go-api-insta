# Use a imagem oficial do Go como base
FROM golang:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para baixar as dependências primeiro
COPY go.mod go.sum ./
RUN go mod download

COPY .env ./

# Copie o restante do código-fonte para o contêiner
COPY . .

# Compile o código Go para um executável
RUN go build -o main

# Exponha a porta em que sua aplicação Go será executada
EXPOSE 8080

# Comando para iniciar a aplicação quando o contêiner for iniciado
CMD ["./main"]
