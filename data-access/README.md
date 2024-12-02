# Setup

```
cd /workspaces/gotutorials/data-access

mysql -u root

ALTER USER 'root'@'localhost' IDENTIFIED VIA mysql_native_password USING PASSWORD('new-password') or unix_socket;

source ./create-tables.sql;

quit
```

# Run

In another terminal
```
cd /workspaces/gotutorials/data-access

export DBUSER="root"
export DBPASS="new-password"

go run .
```