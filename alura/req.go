package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func req() {
	file, errFile := os.Open("sites.txt")
	logFile, _ := os.OpenFile("logsites.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if errFile != nil {
		file, _ = os.Create("sites.txt")
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		testSite(line, logFile)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	logFile.Close()
}

func testSite(site string, file *os.File) {
	fmt.Println("Testing site", site)
	nowTime := time.Now().Format("02/01/2006 15:04:05")

	for i := 0; i < 3; i++ {
		resp, _ := http.Get(site)

		file.WriteString(fmt.Sprintf("[%s][%d] Request %s\n", nowTime, resp.StatusCode, site))

		time.Sleep(1 * time.Second)
	}
}
