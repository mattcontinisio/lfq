# lfq

[![](https://github.com/mattcontinisio/lfq/actions/workflows/test.yml/badge.svg)](https://github.com/mattcontinisio/lfq/actions)

lfq is a command-line [logfmt](https://brandur.org/logfmt) processor. It reads logfmt or JSON inputs, can optionally filter keys, and outputs values, logfmt, or JSON.

## Usage
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
