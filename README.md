<!-- ABOUT THE PROJECT -->

## About The Project

RESTFUL API for TakTuku an E-Commerce App created for the purpose of study.

Building the project with layered architecture, and clean code approach for the structure, with the intention of simplicity when the app is scaling up and ease of maintenance

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

This project structure is built using

- [Swagger](https://app.swaggerhub.com/apis-docs/HamzahAA15/TakTuku-Project/1.0.0#/Products/get_products_myproduct)
- [Golang]
- [Mysql]
- [Labstack/Echo]

<p align="right">(<a href="#top">back to top</a>)</p>

### Features

- USERS CRUD
- PRODUCT CRUD
- CART CRUD
- ORDER CR

### Folder Structure

```
├── addMiddleware/                  # Create middleware
├── config/                         # Configuration to connect to database
├── controller/                     # Create controller for user, product, cart, and order
├── entities/                       # Create entities for user, product, cart, and order
├── helper/                         # Create request, response, and helper for user, product, cart, and order
├── repository/                     # Get all required data from database for user, product, cart, and order
├── service/                        # Create service for handle the data from repository of user, product, cart, and order

```

<!-- GETTING STARTED -->

## Getting Started

To start project, just clone this repo

### Installation

1. Clone the repo
   ```bash
   git clone https://github.com/hilmihi/e-commerce-project.git
   ```
2. Create .env file in main directory
   ```bash
   touch .env
   ```
3. Write the following example environment
   ```
   export DB_CONNECTION_STRING='root:[fillpasswordhere]@/[schema name]?charset=utf8&parseTime=True&loc=Local'
   ```
4. Run the server
   ```bash
   source .env && go run main.go
   ```

<p align="right">(<a href="#top">back to top</a>)</p>
