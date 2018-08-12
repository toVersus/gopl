package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, urlStr := range os.Args[1:] {
		filename, nbytes, err := fetch(urlStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch URL failed: %s\n  %s\n", urlStr, err)
			os.Exit(1)
		}
		fmt.Printf("successfully fetch %s (%d bytes)\n", filename, nbytes)
	}
}

// fetch downloads the URL and returns the name
// and length of the local file.
func fetch(urlStr string) (filename string, n int64, err error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		// Close file, but prefer error from Copy, if any,
		// because on many file systems such as NFS write errors are not reported
		// immediately but may be postponed untile the file is closed.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)

	return local, n, err
}
