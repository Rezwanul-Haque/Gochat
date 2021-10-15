# gochat - one 2 one video chating app using golang & agora.io & firebase

## Prerequisites
1. Firebase project
2. Firebase service account keys file
3. Firebase web api key

## Run Locally

Go to the project directory

```bash
cd gochat
```

Install dependencies

```bash
go mod vendor
```

Start the server Locally

```bash
go run main.go serve
```

Start the server using Docker

```bash
make development
```