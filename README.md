# PESU-IO GoLang Final Project

## Overview
This repository contains the final project for the PESU-IO course. The project is developed using GoLang and demonstrates some concepts and techniques learned throughout the course.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Middleware and JWT Authentication](#middleware-and-jwt-authentication)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

## Features
- User authentication (Sign In and Sign Up)
- Secure password handling with bcrypt
- RESTful API using Gin framework
- JSON-based data handling
- JWT Authentication for secure user sessions
- Role-Based Access Control (RBAC) (Coming Soon)

## Installation
To get started with the project, follow these steps:

1. **Clone the repository:**
    ```sh
    git clone https://github.com/7Chethan007/PESU-IO_GoLang_Final_Project.git
    cd PESU-IO_GoLang_Final_Project
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Set up environment variables:**
    Create a `.env` file in the root directory and add the necessary environment variables.

## Usage
To run the project locally, use the following command:

```sh
go run [main.go](http://_vscodecontentref_/0)
## Middleware and JWT Authentication
This project uses middleware for logging and JWT (JSON Web Token) for authentication. The middleware ensures that all requests are logged, and JWT is used to secure endpoints by verifying tokens.

## Testing
To run tests for the project, use the following command:

```sh
go test ./...
```


## Middleware and JWT Authentication
This project uses middleware for logging and JWT (JSON Web Token) for authentication. The middleware ensures that all requests are logged, and JWT is used to secure endpoints by verifying tokens.

## Testing
To run tests for the project, use the following command:

```sh
go test ./...
```
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## Contributing
Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.


## License
This project is licensed under the MIT License. See the [LICENSE](../blob/main/LICENSE) file for details.

## Acknowledgements
- [Gin Gonic](https://github.com/gin-gonic/gin) - Web framework used
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Password hashing
- [JWT](https://jwt.io/) - JSON Web Tokens for authentication
- [GoLang](https://golang.org/) - Programming language used
## Contact
For any inquiries or feedback, please contact:


**Chethan**
- Email: [mnchethan0+golangproject@gmail.com](mailto:mnchethan0+golangproject@gmail.com)
