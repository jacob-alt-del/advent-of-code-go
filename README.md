# advent-of-code-go

This is a slimmed down version of [alexchao26/advent-of-code-go](https://github.com/alexchao26/advent-of-code-go/tree/main)

## Running Locally
### Requirements
Go 1.16+ is required because [embed][embed] is used for input files.

### Run script
Running the script will copy the answer to the clipboard.
```sh
go run main.go -part <1 or 2>
```

## Makefile
Makefile should be fairly self-documenting. Alternatively you can run the binaries yourself via `go run` or `go build`.

### Make skeleton for the current day
```sh
make skeleton
```

### Make skeleton files for select day
```sh
make skeleton DAY=n YEAR=n
```

### Make all skeleton files
```sh
for ((i=1; i<26; i++)); do
make skeleton DAY=$i YEAR=2023
done
```
