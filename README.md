# url

A command-line tool for working with URLs.

## Features

* Parse a URL into its components.
* Decode a URL encoded string or IDNA encoded domain.
* URL encode a string or non-ASCII domain.
* Build a URL from components.

## Examples

Parse a URL.

```text
> url parse "https://mysite.com:8000/my%20documents?file=my+file"
scheme:	https
uri-path:
user:
host:	mysite.com
port:	8000
path:	/my documents
fragment:
param:	file=my file
```

Parse a URL but don't decode the components.

```text
> url parse "https://mysite.com:8000/my%20documents?file=my+file" --no-decode
scheme:	https
uri-path:
user:
host:	mysite.com
port:	8000
path:	/my%20documents
fragment:
param:	file=my+file
```

Parse a URL to JSON.

```text
> url parse "https://mysite.com:8000/my%20documents?file=my+file" --json | jq
{
  "fragment": null,
  "host": "mysite.com",
  "params": {
    "file": "my file"
  },
  "path": "/my documents",
  "port": "8000",
  "scheme": "https",
  "uriPath": null,
  "user": null
}
```

Grep for a particular component.

```text
> url parse "https://mysite.com:8000/my%20documents?file=my+file" | grep path | cut -f2

/my documents
```

Instead of using shell commands, filter directly for a particular component.

```text
> url parse "https://mysite.com:8000/my%20documents?file=my+file" --path
/my documents
```

URL encode a string

```text
> url encode 'this is my ^message))'
this%20is%20my%20%5Emessage%29%29
```

Decode a URL encoded string

```text
> url decode this%20is%20my%20%5Emessage%29%29
this is my ^message))
```

IDNA encode a domain to ASCII

```text
> url encode 你好 --puny
xn--6qq79v
```

Decode a IDNA encoded domain

```text
> url decode xn--6qq79v --puny
你好
```

Build a URL from components

```text
> url build --scheme http --host mysite.com --param "foo=bar" --param "bar=baz"
http://mysite.com?bar=baz&foo=bar
```

Build a URL from JSON

```text
> url build --json '{"scheme":"http","host":"mysite.com","params":{"foo":"bar","bar":"baz"}}'
http://mysite.com?bar=baz&foo=bar
```

Build a URI

```text
> url build --scheme mailto --uri-path nobody@email.com
mailto:nobody@email.com
```

## Installation

`url` is written in Go. to install `url`, first, make sure you have Go installed. Next clone this repo. Finally, Build the `url` command-line tool with `go build -o url main.go`. To use `url` system-wide, copy the `url` executable to a location in your PATH.
