# lorem

A tiny CLI that prints Lorem Ipsum placeholder text.

## Install

```sh
go install .
```

This builds and installs the `lorem` binary into your `$GOBIN` (make sure
that's on your `$PATH`).

## Usage

```sh
lorem [count]
lorem -c [count]
```

- `count` — number of paragraphs to print (default: `1`).
- `-c` — copy the output to the clipboard instead of printing it.

### Examples

```sh
lorem          # one paragraph
lorem 3        # three paragraphs
lorem -c 3     # copy three paragraphs to the clipboard
```

## Requirements

On Linux, clipboard support (`-c`) shells out to `xclip`, `xsel`, or `wl-copy`
(via [`atotto/clipboard`](https://github.com/atotto/clipboard)) — make sure one
of those is installed for your display server.
