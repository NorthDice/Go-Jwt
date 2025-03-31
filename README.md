# Go-jwt

### About
This project is an implementation of an authorization and authentication backend service using PostgreSQL (via GORM) and Gin framework. It utilizes the JWT package for generating and validating JSON Web Tokens (JWT) for secure user authentication and bcrypt for password hashing to ensure password security. The service allows users to register, log in, and authenticate securely using JWT tokens.

### Features

User Registration: Allows users to create an account with a hashed password.

Login: Allows users to authenticate and receive a JWT token for subsequent requests.

JWT Authentication: Uses JSON Web Tokens to manage user sessions securely.

Password Hashing: Implements bcrypt to securely hash and compare passwords.

PostgreSQL with GORM: Utilizes PostgreSQL as the database for storing user credentials and session data, with GORM as the ORM.

Gin Framework: The API is built using the Gin web framework, providing a lightweight and fast solution for handling HTTP requests.

### Tech Stack

Gin: A fast and lightweight web framework for Go.

PostgreSQL: A powerful open-source relational database for storing user data.

GORM: An ORM for Go to interact with PostgreSQL.

JWT: A package for creating and verifying JSON Web Tokens.

bcrypt: A package for securely hashing passwords.
