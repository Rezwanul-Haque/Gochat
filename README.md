# gochat - basic one 2 one video chating app using golang & websocket + webrtc & firebase & react

## Prerequisites
1. Firebase project
2. Firebase service account keys file which can found on firebase project settings `service account` tab
3. Firebase web api key which can found on firebase project settings `general` tab

## Run Locally (Server)

Go to the project directory

```bash
cd gochat
```

Put the firebase service account key file
> check the **fb-svc-key.example.json** file for reference
```
fb-svc-key.json
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

or 

./run.sh
```


## Run Locally (Web)

Go to the project directory

```bash
cd gochat/web
```

Install dependencies

```bash
npm i
```

Start the server Locally

```bash
npm run dev
```