# Markdown code runner

## Usage

[Install Go](https://go.dev/doc/install) and run

``` shell
go run github.com/rimi-itk/go-markdown-code-runner@latest [options] [filename]
```

If no `filename` is specified, `README.md` is used.

List all code block (in `README.md`):

``` shell
go run github.com/rimi-itk/go-markdown-code-runner@latest
```

Show how to run a block:

``` shell
go run github.com/rimi-itk/go-markdown-code-runner@latest --verbose
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
