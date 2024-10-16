# Markdown code runner

Show and run [fenced code blocks](https://github.github.com/gfm/#fenced-code-blocks) in Markdown files.

## Support languages

* `bash`
* `php`
* `shell`, `sh` (will be run with `bash`)
* `zsh`

## Quick start

Assuming [Go is installed](https://go.dev/doc/install), you can run a quick test with

``` shell
go run github.com/mikkelricky/markdown-code-runner@latest show
```

to list all code blocks in `README.md` in the current folder.

## Installation

[Install Go](https://go.dev/doc/install) and install `markdown-code-runner` with

``` shell
go install github.com/mikkelricky/markdown-code-runner@latest
```

See [Compile and install packages and
dependencies](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies) for details on where
`markdown-code-runner` is actually installed.

To set things straight and clean up, it may be worth running these commands:

``` shell
# Create the default installation location
mkdir -p ~/go/bin
# Clear GOBIN to use the default installation location
go env -w GOBIN=''
go install github.com/mikkelricky/markdown-code-runner@latest
```

Add `~/go/bin` to your `PATH`, e.g.

``` zsh
# ~/.zshrc
export PATH=$PATH:$HOME/go/bin
```

See [Completions](#completions) for details in how to set up completions for your terminal.

## Usage

``` shell
markdown-code-runner [options] [filename]
```

If no `filename` is specified, input is read from `stdin` or `README.md` is used.

Show all code block (in `README.md`):

``` shell name=list
markdown-code-runner show
```

Show how to run blocks:

``` shell name=list-verbose
markdown-code-runner show --verbose
```

Show a single block:

``` shell name=show-single
# By name, i.e. a code block with name=coding-standards-markdown
markdown-code-runner show --verbose coding-standards-markdown
# By index
markdown-code-runner show --verbose 5
```

Run a block:

``` shell name=run
# Run the block with name "test"
markdown-code-runner run example
```

Highlight the commands being run:

``` shell name=run-echo
markdown-code-runner run example --echo '\n👉 '
```

(internally `--echo` uses [`PS4`](<https://www.gnu.org/software/bash/manual/bash.html#index-PS4>))

It works with both `stdout` and `stderr`:

``` shell
markdown-code-runner run example-streams

# Silence stdout
markdown-code-runner run example-streams > /dev/null

# Silence stderr
markdown-code-runner run example-streams 2&> /dev/null
```

Interactivity also works:

``` shell
markdown-code-runner run example-bash-interactive
```

And colored and styled output:

``` shell
markdown-code-runner run example-bash-color
```

### Substitutions

"Substitutions" can be defined on a block, e.g.

```` markdown
```php name=example-php-substitutions substitutions='«name»: Mikkel'
<?php echo "Hello «name»!\n";
```
````

and when the code block is run, the result will be

``` console
Hello Mikkel!
```

Use `--substitutions` to substitute values, i.e. override any default substitutions, before showing or running a code
block:

``` shell
markdown-code-runner show example-php-substitutions --substitutions '«name»: James'
markdown-code-runner run example-php-substitutions --substitutions '«name»: James'
```

The substitutions must be a valid YAML mapping mapping a placeholder (e.g. `«name»`) to a value (e.g. `Mikkel`). For
convenience, use [Flow mappings](https://yaml.org/spec/1.2.2/#742-flow-mappings) for multiple values:

``` shell
markdown-code-runner run example-php-substitutions --substitutions '{«name»: Mikkel, «number»: 87}'
```

### Completions

`markdown-code-runner` can automatically generate completions for four shells:

``` shell name=completion-help
markdown-code-runner help completion
```

#### Zsh

Load completions in [Zsh](https://en.wikipedia.org/wiki/Z_shell) by adding

``` zsh
# ~/.zshrc
eval "$(markdown-code-runner completion zsh)"; compdef _markdown-code-runner markdown-code-runner
```

to your `~/.zshrc`. If you're cool, you do it all from the command line:

``` shell name=zshrc-install-completion
cat >> ~/.zshrc <<'EOF'
eval "$(markdown-code-runner completion zsh)"; compdef _markdown-code-runner markdown-code-runner
EOF
```

And if you're even cooler, you use `markdown-code-runner` to run the code snippet above by its name:

``` shell
markdown-code-runner run zshrc-install-completion --verbose
```

## Examples

```shell name=example
date
pwd
```

```shell name=example-shell
echo "$0"
```

```bash name=example-bash
echo "$0"
```

```zsh name=example-zsh
echo "$0"
```

```php name=example-php
<?php

echo PHP_VERSION, PHP_EOL;
```

``` shell name=example-streams
echo "This is written on stdout"
(>&2 echo "This is written on stderr")
```

```php name=example-php
<?php

echo (new DateTimeImmutable())->format(DateTimeInterface::ATOM);
```

```php name=-example-php-html
<html>
  <title><?php echo (new DateTimeImmutable())->format(DateTimeInterface::ATOM); ?></title>
</html>
```

```php name=example-php-substitutions substitutions='«name»: Mikkel'
<?php echo "Hello «name»!\n";
```

```bash name=example-bash-interactive
task test:interactive
```

```bash name=example-bash-color
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

printf "Roses are ${RED}red${NC}. Voilets are ${BLUE}blue${NC}.\n"

for((i=16; i<256; i++)); do
    printf "\e[48;5;${i}m%03d" $i;
    printf '\e[0m';
    [ ! $((($i - 15) % 6)) -eq 0 ] && printf ' ' || printf '\n'
done
```

```php name=example-php-substitutions-err substitutions='«name» Mikkel'
<?php echo "Hello «name»!\n";
```
