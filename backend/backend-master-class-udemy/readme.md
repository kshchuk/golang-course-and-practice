# Masterclass Backend Web Service Implementation

This repository contains the implementation of the backend web service I'm building during the course.

## Course Recap

The service is a simple bank with APIs for various banking operations, including:

- Creating and managing bank accounts, which include the owner's name, balance, and currency.
- Recording all balance changes for each account by creating an account entry record for every financial transaction.
- Performing money transfers between two accounts, ensuring atomicity within a transaction.

## Getting Started with the Implementation

If you're ready to explore the code and work with the implementation, follow these steps to get started:

1. Clone this repository to your local machine.

2. Install the required tools and dependencies, which may include Golang packages and external libraries, as specified in the project files and dependencies.

3. Set up your local development environment and database connections. Refer to the course material for guidance on how to configure these settings.

4. Explore the codebase to understand the structure and components of the backend web service. You'll find directories and files that include routes, controllers, models, and more.

5. Make modifications and improvements to the service based on your requirements or experiment with additional features to expand your knowledge.

## Code Generation

The repository also includes tools and scripts for generating code. You can use these to enhance and maintain the service:

- Generate the schema SQL file with DBML: `make db_schema`

- Generate SQL CRUD with sqlc: `make sqlc`

- Generate database mock with Gomock: `make mock`

- Create a new database migration: `make new_migration name=<migration_name>`

## Running the Implementation

To run the implementation on your local machine:

- Start the server: `make server`

- Run tests: `make test`
