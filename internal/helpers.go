package internal

import (
	"io"
	"net/http"
)

func copyHeader(dest, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dest.Add(k, v)
		}
	}
}

func copyIO(dest io.WriteCloser, src io.ReadCloser) {
	defer dest.Close()
	defer src.Close()

	_, _ = io.Copy(dest, src)
}
