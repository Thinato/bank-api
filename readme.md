# Bank API Project - Study Repository

This repository contains a Go project for a simple Bank API that I developed as part of my personal study and exploration of Go programming language. The purpose of this project was to gain a deeper understanding of web development and backend concepts by building a basic banking application from scratch.

## Disclaimer

**Please note that this project is purely educational and not intended for production use.**

I want to emphasize that I did not use any existing code or reference any external resources while creating this project. The entire implementation is based on my own exploration and understanding of the concepts involved.

## Features

The Bank API project includes the following functionalities:

* Authentication: Users can create accounts and log in securely.
* Deposit: Users can deposit funds into their accounts.
* Withdraw: Users can withdraw funds from their accounts, subject to available balance.
* Transfer: Users can transfer funds between their own accounts or to other users' accounts.
* Investment: Users can opt for investment plans or strategies.

## API Endpoints

The following API endpoints are exposed by the application:

### Client 
* **POST /signup**: Create a new user account.
* **POST /login**: Log in to an existing user account.
* **POST /deposit**: Deposit funds to the user's account.
* **POST /withdraw**: Withdraw funds from the user's account.
* **POST /transfer**: Transfer funds from the user's account to another account.
* **POST /investment**: Opt for an investment plan.
* **GET /balance**: 
* **GET /transactions**: 

### Manager
* **POST /create-account**: 
* **POST /archive-account**:
* **POST /change-password**:
* **GET /client-data**:
* **GET /all-clients**:


## Getting Started

To run this project locally or explore the code, follow these steps:

1. Clone the repository: git clone https://github.com/Thinato/bank-api.git
1. Change directory to the project folder: cd bank-api
1. Install the required dependencies: go get -d ./...
1. Build the project: go build
1. Run the application: ./bank-api

Make sure you have Go installed on your system before attempting to build and run the project.

## Contributions

As mentioned earlier, this project is a personal study, and I do not intend to actively maintain or accept contributions. However, feel free to fork the repository and explore the code on your own.

## Feedback

If you have any feedback, suggestions, or questions about this project, please feel free to reach out to me. I'm open to discussing the code and concepts used in this repository.