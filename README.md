# Gator

Gator is a command-line RSS aggregator built with **Go** and **PostgreSQL**. It allows you to register users, follow/unfollow RSS feeds, and browse aggregated posts directly from the terminal.

---

## 📦 Prerequisites

Before running the project, ensure you have the following installed:

- [PostgreSQL](https://www.postgresql.org/) – Used as the database.  
- [Go (Golang)](https://go.dev/dl/) – Required to build and run the program.  

---

## ⚙️ Installation

### Install Gator
```bash
go install github.com/benedekpal/gator@latest
```

### Local Setup

1. Create a config file in your home directory:  
   `~/.gatorconfig.json`

   ```json
   {
     "db_url": "connection_string_goes_here",
     "current_user_name": "username_goes_here"
   }
   ```

---

## 🛢️ PostgreSQL Setup

Install PostgreSQL:
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo passwd postgres
```

Start the PostgreSQL server:
```bash
sudo service postgresql start
```

Set up the database:
```bash
sudo -u postgres psql
```

Inside the psql shell:
```sql
CREATE DATABASE gator;
\c gator
ALTER USER postgres PASSWORD 'postgres';
```

Exit the shell:
```bash
exit
```

---

## 🦆 Goose Migration Setup

### Install Goose
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Configure Goose

1. Get your PostgreSQL connection string:
   ```
   protocol://username:password@host:port/database
   ```

   Example:
   ```
   postgres://postgres:postgres@localhost:5432/gator
   ```

   Test connection:
   ```bash
   psql "postgres://postgres:postgres@localhost:5432/gator"
   ```

2. Run database migrations:
   ```bash
   cd sql/schema
   goose postgres <connection_string> up
   ```

3. Update your `~/.gatorconfig.json` file with the connection string (note the required `sslmode=disable`):
   ```json
   {
     "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
     "current_user_name": "username_goes_here"
   }
   ```

---

## 🚀 Example Usage

Below is a list of available commands with examples:

- **Login**
  ```bash
  login myusername
  ```

- **Register a new user**
  ```bash
  register newuser
  ```

- **Reset users and feeds**
  ```bash
  reset
  ```

- **List all users (shows current user)**
  ```bash
  users
  ```

- **Aggregate feeds every 30 seconds**
  ```bash
  agg 30s
  ```

- **Add a new RSS feed**
  ```bash
  addfeed "TechCrunch" "https://techcrunch.com/feed/"
  ```

- **List all feeds**
  ```bash
  feeds
  ```

- **Follow a feed**
  ```bash
  follow "https://techcrunch.com/feed/"
  ```

- **Unfollow a feed**
  ```bash
  unfollow "https://techcrunch.com/feed/"
  ```

- **Show feeds you’re following**
  ```bash
  following
  ```

- **Browse aggregated posts**
  ```bash
  browse
  browse 10   # Limit results
  ```

---
