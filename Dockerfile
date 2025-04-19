FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Download the wait-for-it.sh script
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh

# Make the wait-for-it.sh script executable
RUN chmod +x /wait-for-it.sh

RUN go build -o main ./cmd

EXPOSE 8080

# Use wait-for-it.sh to wait for the database to be ready before starting the app
CMD ["/wait-for-it.sh", "db:5432", "--", "./main"]
