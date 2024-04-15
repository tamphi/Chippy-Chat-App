# OVERVIEW

This is a real-time webchat application. Once users log in, they can see all other users in the room and start a chat. User can view message history between them and other users. 

1. Frontend: React Typescript, Tailwind CSS, Next.js
3. Backend: Golang
4. Protocol: Websocket
5. Database: PostgreSQL
6. Infrastructure: Docker

# FRONTEND

- `cd client`
- `npm install`
- `npm run dev`
- Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.
- For more information, please check the README.MD in the client folder

# BACKEND

- Require: Docker
- Spin up database:
  - Go to server folder: `cd server`
  - Pull postgreSQL image:  `docker pull postgres:16-alpine`
  - Initalize database: `make postgres init`
  - Create database:  `make createdb`
  - Migrate up:  `make migrateup`
- Run backend/server as docker:
  - `docker-compose up --build`

# WORKFLOW

- Set up the frontend and backend as instructed above
- Open [http://localhost:3000](http://localhost:3000) with your browser to visit the login page for the web chat app, signup/login as necessary
- The app only accepts signup request if username has not been created and login request  for correct username and password
- Once logged in, you can now begin your chat. Starting a conversation with another user will by default broadcast a "Howdy" message to the recipient

# FEATURES

* [X] Signup, login, logout authentication
* [X] Save messages to database
* [X] Send, receive, and view messages and conversation history

# FUTURE CONSIDERATION

This is my first time working on an full-stack application project and using all of the tools for this project so it was a great opportunity for me to learn new frameworks, programming languages (Golang), and design flow (database design). In the future, the following would be areas of development that I would like to work on for this project

1. Message reaction and emoji features
2. User profile creation: display name, profile picture
3. UI for alerting users when login and signup fails due to wrong password and exisiting username
4. Refactor the websocket code to create 3 layers of abstraction: database layer <- service layer <- websocket handling layer
