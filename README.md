# Table of Contents

1. [Welcome to Ticket Plus Customer Support Chatbot](#welcome-to-ticket-plus-customer-support-chatbot)
2. [Getting Started with the Project](#getting-started-with-the-project)
   - [Prerequisites](#prerequisites)
   - [Step 1: Clone the Repository](#step-1-clone-the-repository)
   - [Step 2: Install Dependencies](#step-2-install-dependencies)
   - [Step 3: Set Up Environment Variables](#step-3-set-up-environment-variables)
   - [Step 4: Run Tests](#step-4-run-tests)
   - [Step 5: Build and Run the Server](#step-5-build-and-run-the-server)
3. [Check Server Status](#check-server-status)
4. [Troubleshooting](#troubleshooting)


# Welcome to Ticket Plus Customer Support Chatbot

This chatbot is designed to enhance customer service for the Ticket Plus mobile app by leveraging AI to provide fast and accurate responses. The chatbot integrates with the Ticket Plus platform, allowing users to get quick resolutions to their issues or queries.

## Getting Started with the Project

Follow the steps below to set up and run the Ticket Plus chatbot on your local machine.

### Prerequisites

Before getting started, ensure that Go (Golang) is installed on your system. You can verify this by running:

```bash
go version
```
If Go is not installed, follow the (official Go installation guide)[https://go.dev/doc/install] based on your operating system.

### Step 1: Clone the Repository

Clone the project repository using Git:
```bash
git clone git@github.com:Sound-X-Team/ticketplus-chat-bot.git
```

Navigate to the project directory:

```bash
cd ticketplus-chat-bot
```

### Step 2: Install Dependencies

The project uses Go modules to manage its dependencies. To ensure all required packages are installed, run the following command:

```bash
go mod tidy
```


This will download and install all dependencies required by the project, as specified in the `go.mod` file.

### Step 3: Set Up Environment Variables

The chatbot requires certain environment variables to function correctly (e.g., API keys, database credentials, or server configurations). Please ask the project lead or refer to any available `.env.example` file for the necessary variables.

To set environment variables:

- You can either export them in your terminal session using `export VARIABLE_NAME=value`, or
- Create a `.env` file in the project root with key-value pairs, and use a tool like `godotenv` to load them.

### Step 4: Run Tests

Before running the server, it's recommended to run the test suite to ensure everything is working as expected. Run the following command in the root of the project:

```bash
go test ./...

```

### Step 4: Build and Run the Server

To start the chatbot server, build and run 

```bash 
go build && ./ticketplus-chat-bot
```
By default, the server will start on the port specified in your environment variables (typically 8080), or you can check the logs for confirmation. If no port is set, the server may fall back to a default port like 8080.

To see if the server is waiting for requests, you can run the following command:

```bash
curl -X GET http://127.0.0.1:8080/v1/healthz
```

You should get a response like `{}%` or a status `200`.

If you encounter any issues, please reach out to your team lead or members for assistance.



