package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Get() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
			fmt.Println("Добавлен префикс, url = " + url)
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Http request responce status code: %v\n", resp.Status)

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch чтение %s: %v", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
