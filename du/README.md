Implement a tool to return the total number of bytes of all the files in a directory (and its subdirectories).

The behaviour is similar to the `du` command:

```
$ du -ha testdata
4.0K    testdata/txt1/1.txt
4.0K    testdata/txt1/foo/a.txt
4.0K    testdata/txt1/foo/b.txt
4.0K    testdata/txt1/foo/bar/c.txt
4.0K    testdata/txt1/foo/bar
 12K    testdata/txt1/foo
 16K    testdata/txt1
4.0K    testdata/txt2/a.txt
4.0K    testdata/txt2
 20K    testdata
```

The main difference is that `du` outputs the number of filesystem block usage. On my machine, the block size is 4k, the minimum size for a file regardless of how much content there is.

Our tool would output the number of bytes in the file.

## Instruction

* [os.Stat](https://golang.org/pkg/os/#Stat)
  * Returns filesystem information about a path.
* [os.FileInfo.IsDir](https://golang.org/pkg/os/#Stat)
  * Check whether a file is a directory.
* [os.File.Readdir](https://golang.org/pkg/os/#File.Readdir)
  * Returns the file names and directory names of a directory.

## Examples

```
$ go run du.go testdata
testdata/txt1/1.txt: 698
testdata/txt1/foo/a.txt: 2245
testdata/txt1/foo/b.txt: 716
testdata/txt1/foo/bar/c.txt: 683
testdata/txt2/a.txt: 720
disk usage: 5062
```