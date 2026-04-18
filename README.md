# URL Shortener CLI

A fast and simple command-line application built in Go to shorten URLs, list them, and remove them. The application connects to a MongoDB database to persist data and uses `tablewriter` to beautifully format terminal output.

## Features

* **Shorten URLs**: Quickly generate short aliases for any long URL.
* **List URLs**: View all your shortened URLs in a nicely formatted terminal table.
* **Remove URLs**: Delete specific URLs from your database when they are no longer needed.
* **Persistent Storage**: Utilizes MongoDB for keeping your data safe and readily available.

## Prerequisites

Make sure you have the following installed to run this project:

* [Go](https://golang.org/dl/) (version 1.26+)
* [MongoDB](https://www.mongodb.com/) (either running locally or a cloud instance like MongoDB Atlas)

## Setup & Installation

1. **Navigate to the project directory**:
   ```bash
   cd url-shortner-cli
   ```

2. **Install Dependencies**:
   The project uses Go Modules. Download all required dependencies by running:
   ```bash
   go mod tidy
   ```

3. **Environment Configuration**:
   Create a `.env` file in the root directory of the project to store your configurations (like your MongoDB connection string).
   ```env
   MONGO_URI="mongodb://localhost:27017" # Replace with your actual Mongo URI
   # Depending on your config, you may need other environment variables like PORT or DB_NAME
   ```

## Usage

You can use the CLI directly via `go run main.go` or build it into an executable.

| Command | Syntax | Description |
| :--- | :--- | :--- |
| **Shorten URL** | `go run main.go shorten <url>` | Create a shortened version of a long link. |
| **List URLs** | `go run main.go list` | View all the URLs you've shortened in an ASCII table. |
| **Remove URL** | `go run main.go remove <url>` | Delete a URL from your database given its identifier. |

## Built With

* **[Go](https://golang.org/)** - Core logic and CLI structure
* **[MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)** - Database connection and operations
* **[Godotenv](https://github.com/joho/godotenv)** - Environment variable management
* **[Tablewriter](https://github.com/olekukonko/tablewriter)** - ASCII table output generation
* **[MachineID](https://github.com/denisbrodbeck/machineid)** - Generating unique hardware IDs

## License

MIT
