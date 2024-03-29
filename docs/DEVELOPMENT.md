# Development

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
