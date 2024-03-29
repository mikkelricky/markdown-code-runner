# Markdown code runner

## Usage

[Install Go](https://go.dev/doc/install) and run

``` shell
go run github.com/mikkelricky/go-markdown-code-runner@latest [options] [filename]
```

If no `filename` is specified, input is read from `stdin` or `README.md` is used.

List all code block (in `README.md`):

``` shell name=list
go run github.com/mikkelricky/go-markdown-code-runner@latest
```

Show how to run blocks:

``` shell name=list-verbose
go run github.com/mikkelricky/go-markdown-code-runner@latest --verbose
```

Show a single block:

``` shell name=show-single
go run github.com/mikkelricky/go-markdown-code-runner@latest --verbose --show 5
go run github.com/mikkelricky/go-markdown-code-runner@latest --verbose --show coding-standards-markdown
```

## Development

```shell name=build
task build
```

## Test

``` shell name=test
pwd
(>&2 echo FEJL)
date
pwd
for i in {0..10}; do
    (>&2 echo FEJL $i)
    date
done
```

``` shell name=long-running-test
find ~ -type f
```

``` php a=b c=d
task format
task test
```

### Coding standards

```shell name=coding-standards-markdown
task dev:coding-standards:check
```
