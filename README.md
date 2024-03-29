# Markdown code runner

Show and run [fenced code blocks](https://github.github.com/gfm/#fenced-code-blocks) in Markdown files.

## Quick start

Assuming [Go is installed](https://go.dev/doc/install), you can run a quick test with

``` shell
go run github.com/mikkelricky/go-markdown-code-runner@latest help
```

## Installation

[Install Go](https://go.dev/doc/install) and install `go-markdown-code-runner` with

``` shell
go install github.com/mikkelricky/go-markdown-code-runner@latest
```

See [Compile and install packages and
dependencies](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies) for details on where
`go-markdown-code-runner` is actually installed.

To set things straight and clean up, it may be worth running these commands:

``` shell
# Create the default installation location
mkdir -p ~/go/bin
# Clear GOBIN to use the default installation location
go env -w GOBIN=''
go install github.com/mikkelricky/go-markdown-code-runner@latest
```

Add `~/go/bin` to your `PATH`, e.g.

``` zsh
# ~/.zshrc
export PATH=$PATH:$HOME/go/bin
```

### Completions

`go-markdown-code-runner` can automatically generate completions for four shells:

``` shell name=completion-help
go-markdown-code-runner help completion
```

#### Zsh

Load completions in [Zsh](https://en.wikipedia.org/wiki/Z_shell) by adding

``` zsh
# ~/.zshrc
eval "$(go-markdown-code-runner completion zsh)"; compdef _go-markdown-code-runner go-markdown-code-runner
```

to your `~/.zshrc`. If you're cool, you do it all from the command line:

``` shell name=zshrc-install-completion
cat >> ~/.zshrc <<'EOF'
eval "$(go-markdown-code-runner completion zsh)"; compdef _go-markdown-code-runner go-markdown-code-runner
EOF
```

And if you're even cooler, you use `go-markdown-code-runner` to execute the code snippet above by its name:

``` shell
go-markdown-code-runner execute zshrc-install-completion --verbose
```

## Usage

``` shell
go-markdown-code-runner [options] [filename]
```

If no `filename` is specified, input is read from `stdin` or `README.md` is used.

Show all code block (in `README.md`):

``` shell name=list
go-markdown-code-runner show
```

Show how to run blocks:

``` shell name=list-verbose
go-markdown-code-runner show --verbose
```

Show a single block:

``` shell name=show-single
# By name
go-markdown-code-runner show --verbose coding-standards-markdown
# By index
go-markdown-code-runner show --verbose 5
```

Execute a block:

``` shell name=execute
# Execute the block with name "test"
go-markdown-code-runner execute test
```

Highlight the commands being run:

``` shell name=execute-echo
go-markdown-code-runner execute test --echo '👉 '
```

(internally `--echo` uses [`PS4`](<https://www.gnu.org/software/bash/manual/bash.html#index-PS4>))

It works with both `stdout` and `stderr`:

``` shell
# Silence stdout
go-markdown-code-runner execute test > /dev/null

# Silence stderr
go-markdown-code-runner execute test 2&> /dev/null
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
