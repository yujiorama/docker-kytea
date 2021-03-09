package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {

	var file *os.File

	if len(os.Args) == 1 || (len(os.Args) > 1 && os.Args[1] == "-") {

		file = os.Stdin

	} else {

		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("open error: %v", err.Error())
		}
		file = f

	}

	defer file.Close()
	parseKytea(file)
}

func parseKytea(f *os.File) {

	var wg sync.WaitGroup

	wg.Add(1)
	bufChan := make(chan string, 10)
	go func(bufChan <-chan string) {
		defer wg.Done()

		for s := range bufChan {
			println(chasenize(s))
		}
	}(bufChan)

	previous := ""
	current := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		current = scanner.Text()
		if len(previous) == 0 {
			previous = current
			continue
		}

		if concatable(previous, current) {
			previous = concat(previous, current)

		} else {

			bufChan <- previous
			previous = current

		}
	}

	bufChan <- current

	close(bufChan)
	wg.Wait()
}

func concatable(previous, current string) bool {

	pxs := strings.Split(previous, ",")
	cxs := strings.Split(current, ",")

	if len(cxs) != 3 {
		return false
	}

	if cxs[1] == "語尾" {
		return true
	}

	if len(pxs) != 3 {
		return false
	}

	if pxs[1] == "動詞" && cxs[1] == "助詞" && len([]rune(cxs[0])) == 1 {
		return true
	}

	if pxs[1] == "接頭辞" && cxs[1] == "名詞" {
		return true
	}

	if pxs[1] == "名詞" && cxs[1] == "名詞" {
		return true
	}

	return false
}

func concat(previous, current string) string {

	pxs := strings.Split(previous, ",")
	cxs := strings.Split(current, ",")
	pxs[0] = pxs[0] + cxs[0]
	if pxs[1] == "接頭辞" {
		pxs[1] = cxs[1]
	}
	pxs[2] = pxs[2] + cxs[2]
	return strings.Join(pxs, ",")
}

func chasenize(s string) string {

	xs := strings.Split(s, ",")
	// 語\t品詞,*,*,*,*,*,ふりがな
	return fmt.Sprintf("%s\t%s,*,*,*.*.*,%s,%s", xs[0], xs[1], xs[0], xs[2])
}
