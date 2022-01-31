# Go Chat

Group video chatting app using Golang & Agora & websocket + webrtc & Firebase & React

## Authors

- [@rezwanul-haque](https://www.github.com/rezwanul-haque)

## Contributors & Special Thanks
- [@tahsin-siad](https://github.com/tahsinsiad) for helping me with React part.

## 🔗 Links

<p float="left">
    <a href="https://rezwanul-haque.hashnode.dev/"><img src="https://img.shields.io/badge/my_blog-004?style=for-the-badge&logo=blogger&logoColor=white" width="100" height="30px" /></a>
    <a href="https://rezwanul-haque-portfolio.herokuapp.com/en/"><img src="https://img.shields.io/badge/my_portfolio-001?style=for-the-badge&logo=ko-fi&logoColor=white" width="100" height="30px" /> </a>
    <a href="https://www.linkedin.com/in/rezwanul-haque/"><img src="https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white" width="100" height="30px" /></a>
    <a href="https://twitter.com/Rezwanul__Haque"><img src="https://img.shields.io/badge/twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white" width="100" height="30px" /></a>
</p>

## Tech Stack

#### **Client:**

<p>
    <img src="https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB" width="80" height="30px" />
    <img src="https://img.shields.io/badge/bootstrap-%23563D7C.svg?style=for-the-badge&logo=bootstrap&logoColor=white" width="80" height="30px" />
    <img src="https://img.shields.io/badge/MUI-%230081CB.svg?style=for-the-badge&logo=material-ui&logoColor=white" width="80" height="30px" />
</p>

#### **Server:**

<p>
    <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=flat-square&logo=go&logoColor=white" width="80" height="30px" />
    <img src="https://img.shields.io/badge/Echo-Framework-brightgreen" width="90" height="30px" />
    <img src="https://img.shields.io/badge/firebase-%23039BE5.svg?style=for-the-badge&logo=firebase" width="80" height="30px" />
    <img src="https://img.shields.io/badge/Agora-Agora.io-blue" width="90" height="30px" />
</p>

## Prerequisites

1. Firebase project
2. Firebase service account keys file which can be found on firebase project settings `service account` tab
3. Firebase web API key which can be found on firebase project settings `general` tab
4. Enable `Email/Password` based authentication in firebase which can be found on `authentication` tab
5. agora project app key and app certificate which can be found on agora `project management` tab

## Limitation

- Unable to share the link for another user to join the channel/room.
- If another user is not authenticated, redirect him to the login page that is missing.
- Web app routes need to be authenticated

## Current Working Procedure(Steps)

1. Open two browser tabs (incognito tabs will not work for now)
2. Login or Sign up with email and password.
3. Create a Room
4. Join the call
5. Go to the 2nd browser tab and paste the 1st tabs link & press the join call button

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

Put Agora App Id in `settings.js` file in components folder

Install dependencies

```bash
npm i
```

Start the server Locally

```bash
npm run dev
```
