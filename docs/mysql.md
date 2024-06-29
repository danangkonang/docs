```bash
# connect to mysql
mysql -u danang --port 3306 --host localhost db-csv -p
```

```sql
-- CREATE USER FOR DB
CREATE DATABASE databaseName;
CREATE USER 'dnguyen'@'localhost' IDENTIFIED BY '123456';
GRANT ALL PRIVILEGES ON databaseName.* TO 'dnguyen'@'localhost';
FLUSH PRIVILEGES;
QUIT
```