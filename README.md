# sumarejidump

Just dump sumareji tables into CSV

## Usage

```
sumarejidump [table_name] [option]

Just dump sumareji tables.
SUMAREJI_CONTRACT_ID and SUMAREJI_ACCESS_TOKEN must be stored in environmental variable.

table_name
the name of table for dump. default all tables.

option
-c condition
-d output directory
```

- The output CSV file name is the name of table.
- CSV header filed is same to the attribute of the table.
- The order of CSV filed is same to the No of the table in sumareji specification.

## Development
### Requirements
- Go 1.10+

### Coding Convention
Follow [Golang official Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
