# sumarejidump

Just dump sumareji tables into CSV

[![CircleCI](https://circleci.com/gh/nombre-premier/sumarejidump/tree/master.svg?style=svg)](https://circleci.com/gh/nombre-premier/sumarejidump/tree/master)


## Installation

Ensure you have Go 1.11 or later installed, then:

```bash
git clone https://github.com/nombre-premier/sumarejidump.git
cd sumarejidump
make vendor
make build
```

## Usage

```
NAME:
   sumarejidump - Just dump sumareji data

USAGE:
   sumarejidump [command] [options] [table_name]

VERSION:
   0.2.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --token value, -t value        sumareji access token [$SUMAREJI_ACCESS_TOKEN]
   --contract_id value, -i value  sumareji contract_id [$SUMAREJI_CONTRACT_ID]
   --conditions value, -c value   filter data by given conditon(s)
   --output value, -o value       output dir name, default: yyyyMMDDhhmmss
   --help, -h                     show help
   --version, -v                  print the version
```

- The output CSV file name is the name of table.
- CSV header filed is same to the attribute of the table.
- The order of CSV filed is same to the No of the table in sumareji specification.

## Example

```
# single condition
# condition is splited by ":" and convert into json key-value pair.
sumarejidump -c categoryId:1 category

# multiple condtions
sumarejidump -c categoryId >=:1 -c cateoryCode:dbss category

# specify output directory name
# default is yyyyMMDDhhmmss, such as 20190102030405
sumarejidump -o result category
```


## Development
### Requirements
- Go 1.11+

### Coding Convention
Follow [Golang official Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
