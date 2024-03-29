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

Execute a block:

``` shell name=execute
# Execute the block with name "test"
go run github.com/mikkelricky/go-markdown-code-runner@latest --execute test
```

Highlight the commands being run:

``` shell name=execute-echo
go run github.com/mikkelricky/go-markdown-code-runner@latest --execute test --echo '👉 '
```

(internally `--echo` uses [`PS4`](<https://www.gnu.org/software/bash/manual/bash.html#index-PS4>))

It works with both `stdout` and `stderr`:

``` shell
# Silence stdout
go run github.com/mikkelricky/go-markdown-code-runner@latest --execute test > /dev/null

# Silence stderr
go run github.com/mikkelricky/go-markdown-code-runner@latest --execute test 2&> /dev/null
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

```shell name=tty-test
docker run --interactive --rm --volume ${PWD}:/app itkdev/php8.1-fpm:latest pwd
```

``` shell name=curl
curl "https://httpbin.org/anything" --header "content-type: application/json" --data @- <<'BODY'
[1,2,3]
BODY
```

``` shell name=empty
```

```php name=php
<?php

echo (new DateTimeImmutable())->format(DateTimeInterface::ATOM);
```

```php name=php-html
<html>
  <title><?php echo (new DateTimeImmutable())->format(DateTimeInterface::ATOM); ?></title>
</html>
```
