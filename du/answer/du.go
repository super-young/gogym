package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	name := os.Args[1]

	size, err := usage(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("total usage: %d\n", size)
}

func usage(name string) (int64, error) {

	stat, err := os.Stat(name)

	if err != nil {
		return 0, err
	}

	if stat.IsDir() {
		return dirusage(name, "")
	}

	return stat.Size(), nil

}

func dirusage(name, basedir string) (int64, error) {
	fullname := path.Join(basedir, name)
	d, err := os.Open(fullname)
	if err != nil {
		return 0, err
	}

	ls, err := d.Readdir(-1)
	d.Close() // close before recursion into subdirectory
	if err != nil {
		return 0, err
	}

	var sum int64
	for _, s := range ls {
		if s.IsDir() {
			dirsum, err := dirusage(s.Name(), fullname)
			if err != nil {
				fmt.Println(err)
				continue
			}

			sum += dirsum
		} else {
			fmt.Printf("%s: %d\n", path.Join(fullname, s.Name()), s.Size())
			sum += s.Size()
		}
	}

	return sum, nil
}
