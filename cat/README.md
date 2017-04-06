Implement a command like `cat` to print files to stdout.

## Instructions

* [os.Stdout](https://golang.org/pkg/os/#pkg-variables)
  * This is an io.Writer
* [os.Open](https://golang.org/pkg/os/#Open)
  * This is an io.Reader
* [io.Copy](https://golang.org/pkg/io/#Copy)
  * Use this to copy data from io.Reader to io.Writer
* [os.File.Close](https://golang.org/pkg/os/#File.Close)
  * Always remember to close an os.File

## Examples

```
$ go run cat.go a.txt b.txt
aaaaaaa
a
aa
aaaaa
a
aaaaaaaaaaaa
aa
b
bbb
bbbb
bbbbbb
bb
```

```
$ go run cat.go b.txt a.txt
b
bbb
bbbb
bbbbbb
bb
aaaaaaa
a
aa
aaaaa
a
aaaaaaaaaaaa
aa
```

## Error Examples

Print error if a file cannot be found

```
$ go run cat.go a.txt b.txt notfound.txt
aaaaaaa
a
aa
aaaaa
a
aaaaaaaaaaaa
aa
b
bbb
bbbb
bbbbbb
bb
open notfound.txt: no such file or directory
```