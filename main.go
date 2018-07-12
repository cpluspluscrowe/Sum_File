package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.OpenFile("/Users/ccrowe/Desktop/hive_spent.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	size := fi.Size()
	sum := 0.0
	for i := 0; i < int(size); i++ {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file line error %v", err)
			return
		}
		line = string(line)
		line = strings.Replace(line, "\n", "", -1)
		line_split := strings.Split(string(line), "	")
		spend, err := strconv.ParseFloat(line_split[1], 64)
		sum += spend
		if err != nil {
			panic(err)
		}
		_ = line
	}
	fmt.Println(sum)
}
