package xrevision

import "net/http"

const DefaultHeaderKey = "X-Revision"

type GenerateFunc func(string) (string, error)
type XRequestID struct {
	Revision  string
	HeaderKey string
}

func New(n string) *XRequestID {
	return &XRequestID{
		Revision:  n,
		HeaderKey: DefaultHeaderKey,
	}
}

func (m *XRequestID) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	r.Header.Set(m.HeaderKey, m.Revision)
	rw.Header().Set(m.HeaderKey, m.Revision)
	next(rw, r)
}
