# Gator CLI

Gator is a command-line tool for managing RSS feeds and posts. It allows users to follow feeds, scrape posts, and browse them directly from the terminal.

## Prerequisites

Before you can use Gator, ensure you have the following installed:

- **PostgreSQL**: A relational database system. You can download it from [PostgreSQL's official website](https://www.postgresql.org/).
- **Go**: The Go programming language. You can download it from [Go's official website](https://go.dev/).

## Installation

To install the Gator CLI, run the following command:

```bash
go install github.com/your-username/gator@latest
This will install the gator binary to your $GOPATH/bin directory. Ensure that $GOPATH/bin is in your system's PATH.


Configuration
Gator requires a configuration file to connect to the database. Create a config.json file in the root of your project with the following structure:
{
  "DBURL": "postgres://username:password@localhost:5432/database_name?sslmode=disable"
}

Replace username, password, and database_name with your PostgreSQL credentials.


Running the Program
Once installed, you can run the gator CLI directly from your terminal. For example:
gator <command> [args...]

Example Commands
Here are some of the commands you can use with Gator:


Register a new user:


gator register <username> <password>
Login as a user:


gator login <username> <password>
Add a new feed:


gator addfeed <feed-url>
Follow a feed:

gator follow <feed-url>
Browse posts:


gator browse [limit]
Aggregate feeds:


gator agg <time_between_reqs>
Development
For development purposes, you can run the program using: