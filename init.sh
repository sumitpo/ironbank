folder="sql-bank-data"
url="https://github.com/haggarw3/sql-bank-data.git"
if [ ! -d $folder ]
then
  git clone $url $folder
fi

patch sql-bank-data/initdb/mysql_dump.sql bank.patch
mv sql-bank-data/initdb/mysql_dump.sql ./init.sql

rm -rf $folder
