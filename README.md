# Kickstart Backend with GO and MYSQL

### Setup

1 - **Clone the repository**:
```
git clone https://github.com/Ion-Stefan/go-kickstart-backend
```

2 - **Create the MYSQL database**:
I'm using maria db, but you can use other database.
```
sudo mariadb
CREATE DATABASE your_database_name;
```

2.1 optional - **Create a new MYSQL user**:
```
sudo mariadb
CREATE USER 'your_user'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON your_database_name.* TO 'your_user'@'localhost';
FLUSH PRIVILEGES;
```

3 - **Update .env variables**:
In the .env file update the following variables:
```
DB_USER="your_user"
DB_PASSWORD="your_password"
DB_NAME="your_database_name"
```

4 - **Make database migrations**:
```
make migrate-up
```

5 - **Run the project**:
```
make run
```

Notes:
- The project will run on port 8080 by default.
- You can change the port in the .env file.
- The project will run on localhost by default.
- You can change the host in the .env file.
- For more commands check the Makefile.
