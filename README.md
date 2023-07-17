# lfq

[![](https://github.com/mattcontinisio/lfq/actions/workflows/test.yml/badge.svg)](https://github.com/mattcontinisio/lfq/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattcontinisio/lfq)](https://goreportcard.com/report/github.com/mattcontinisio/lfq)

lfq is a command-line [logfmt](https://brandur.org/logfmt) processor. It reads logfmt or JSON inputs, can optionally filter keys, and outputs values, logfmt, or JSON.

## Usage

```sh
$ lfq
lfq is a tool for processing logfmt inputs, filtering keys, and producing results as values, logfmt, or JSON on
standard output.

Usage:
  lfq [keys] [flags]

Examples:
lfq -o value time,level

Flags:
      --disable-quote        disable quoting of all values (for value and logfmt output)
      --force-quote          force quoting of all values (for value and logfmt output)
  -h, --help                 help for lfq
  -i, --input string         input format (logfmt or json) (default "logfmt")
      --no-color             disable colorized output
  -o, --output string        output format (value, logfmt, or json) (default "logfmt")
      --quote-empty-fields   wrap empty values in quotes (for value and logfmt output)
```

## Examples

```sh
# Default - read logfmt and write logfmt
$ cat examples/test.log | lfq
time="2023-07-16T01:10:40Z" level=info msg="Info message" client_id=2
time="2023-07-16T01:10:41Z" level=warning msg="Warning message" client_id=1
time="2023-07-16T01:10:42Z" level=error msg="Error message" client_id=1
time="2023-07-16T01:10:43Z" level=info msg="Info message" client_id=2

# Filter keys
$ cat examples/test.log | lfq time,level
time="2023-07-16T01:10:40Z" level=info
time="2023-07-16T01:10:41Z" level=warning
time="2023-07-16T01:10:42Z" level=error
time="2023-07-16T01:10:43Z" level=info

# Write values only
$ cat examples/test.log | lfq -o value level
info
warning
error
info

# Read JSON
$ cat examples/test.log.json | lfq -i json -o value level
info
warning
error
info

# Write JSON
$ cat examples/test.log | lfq -o json
{"time":"2023-07-16T01:10:40Z","level":"info","msg":"Info message","client_id":"2"}
{"time":"2023-07-16T01:10:41Z","level":"warning","msg":"Warning message","client_id":"1"}
{"time":"2023-07-16T01:10:42Z","level":"error","msg":"Error message","client_id":"1"}
{"time":"2023-07-16T01:10:43Z","level":"info","msg":"Info message","client_id":"2"}

# Write JSON and pipe to jq
$ cat examples/test.log | lfq -o json | jq
{
  "time": "2023-07-16T01:10:40Z",
  "level": "info",
  "msg": "Info message",
  "client_id": "2"
}
...
```

## Other useful tools

- [hutils](https://github.com/brandur/hutils) - A collection of command line utilities for working with logfmt
- [jq](https://github.com/jqlang/jq) - Command-line JSON processor
