# Markdown code runner

Build

```shell
task build
```

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

``` shell name=long-test
find ~ -type f
```

``` php a=b c=d
task format
task test
```
