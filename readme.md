---
colorlinks: true
---

# sample data from
[ironbank database sql](https://github.com/haggarw3/sql-bank-data.git)

```bash
git clone https://github.com/haggarw3/sql-bank-data.git
patch sql-bank-data/initdb/mysql_dump.sql bank.patch
```


# doc
```bash
brew install pandoc
pandoc readme.md -o readme.pdf
```

# er tool
[sqldiagram: generate mysql er diagram command line](https://github.com/RadhiFadlillah/sqldiagram.git)

```bash
go install github.com/RadhiFadlillah/sqldiagram@latest
sqldiagram mysql sql-bank-data/initdb/mysql_dump.sql -o ironbank.svg
```
