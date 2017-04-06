Implement a command like `cat` to print files to stdout.

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