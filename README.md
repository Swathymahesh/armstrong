# Armstrong Web App

This is a full-stack application that:

1. Registers users and stores their information in a MySQL database.
2. Verifies if a given number is an Armstrong number and stores verified numbers associated with a user.
3. Retrieves Armstrong numbers for specific users.
4. Displays all users and their Armstrong numbers in a global dashboard.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Backend Setup](#running-the-backend)
- [Frontend Setup](#running-the-frontend)
- [API Endpoints](#api-endpoints)


---

## Prerequisites

- **Go** (version 1.18 or above recommended): [Download Go](https://go.dev/dl/)
- **Node.js and npm** (for the frontend React application): [Download Node.js](https://nodejs.org/)
- **MySQL Database**:  
  Ensure you have:
  - A running MySQL instance.
  - Appropriate credentials (username, password).
  - A created database for this project.


## Running the Backend

1. **Navigate to the backend directory:**
    ```bash
    cd armstrong-webapp

2. **Install Go dependencies**
    ```bash
   go mod tidy

4. **Run the backend server**
    ```bash
    go run main.go

## Running the Frontend

1. **Navigate to the frontend directory:**
    ```bash
   cd frontend
    
3. **Install frontend dependencies**
    ```bash
    npm install
   
4. **Start the React development server**
    ```bash
   npm start

## API Endpoints 
    Base URL: http://localhost:8080

**Registers a new user.**

```POST /register```

**Verifies if a given number is Armstrong for a specific user and saves it.**

```POST /verify```

**Fetches all Armstrong numbers associated with the given user ID.**

```GET /user/{id}/numbers```

**Fetches all users along with their Armstrong numbers.**

```GET /users```


