
Implement a program that's similar to the `date` command.

```
$ date
Thu  6 Apr 2017 13:32:43 CST
```

## Instructions

+ [time.Now](https://golang.org/pkg/time/#Now)
+ [time.Time.Format](https://golang.org/pkg/time/#Time.Format)
+ [time.RFC1123](https://golang.org/pkg/time/#pkg-constants)
  + This is a predefined format string to print out the time

## Examples

```
$ go run date.go
Thu, 06 Apr 2017 13:32:36 CST
```