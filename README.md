# runsehll

# findinstance
## install
```
$ git clone https://github.com/ritarock/runshell.git
$ cd runshell/
$ go install
```

## Usage
```
NAME:
   runshell - Run commands

USAGE:
   [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, -f value  set file
   --help, -h              show help (default: false)
```

## Example
`test_data/test.txt` looks like this.
```test.txt
[sleep 2, echo 'sleep 2']
[sleep 3, echo 'sleep 3']
[echo 'hello']

```

In this case, Runs in parallel, row by row.

```
$ runshell -f ./test_data/test.txt
hello
sleep 2
sleep 3
```
