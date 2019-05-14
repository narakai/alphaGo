package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func read1(path string) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
	}
}

func read2(path string) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
	}
}

func read3(path string) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	_, err = ioutil.ReadAll(fi)
}

func main() {

	file := "" //找一个大的文件，如日志文件
	start := time.Now()
	read1(file)
	t1 := time.Now()
	fmt.Printf("Cost time %v\n", t1.Sub(start))
	read2(file)
	t2 := time.Now()
	fmt.Printf("Cost time %v\n", t2.Sub(t1))
	read3(file)
	t3 := time.Now()
	fmt.Printf("Cost time %v\n", t3.Sub(t2))
}

//基本上保持 bufio < ioutil < file.Read 这样的成绩， bufio读同一文件耗费时间最少， 效果稳稳地保持在最佳。