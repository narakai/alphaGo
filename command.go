package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
}

//一旦面对复杂的参数格式，比较费时费劲，所以这时我们会选择flag库
//https://github.com/spf13/cobra