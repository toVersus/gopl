package reader

import "io"

type limitReader struct {
	r io.Reader
	n int
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if len(p) > l.n {
		p = p[0:l.n]
	}
	n, err = l.r.Read(p)
	l.n -= n
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r: r, n: int(n)}
}
