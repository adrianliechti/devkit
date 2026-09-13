package proxy

import (
	"crypto/subtle"
	"encoding/base64"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	Username string
	Password string
}

type Proxy struct {
	username string
	password string

	transport *http.Transport
}

func New(c Config) *Proxy {
	proxy := &Proxy{
		username: c.Username,
		password: c.Password,

		// Clone, so configuring this proxy never mutates http.DefaultTransport.
		transport: http.DefaultTransport.(*http.Transport).Clone(),
	}

	return proxy
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p.username != "" || p.password != "" {
		user, pass, ok := parseBasicAuth(r)

		if !ok {
			w.Header().Set("Proxy-Authenticate", "Basic")

			http.Error(w, "Proxy Authentication Required", http.StatusProxyAuthRequired)
			return
		}

		if !equal(user, p.username) || !equal(pass, p.password) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
	}

	r.Header.Del("Proxy-Connection")
	r.Header.Del("Proxy-Authorization")

	if r.Method == http.MethodConnect {
		p.handleTunneling(w, r)
		return
	}

	p.handleHTTP(w, r)
}

func (p *Proxy) handleHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v", r.Method, r.URL)

	resp, err := p.transport.RoundTrip(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	headers := w.Header()

	for key, value := range resp.Header {
		if strings.EqualFold(key, "Transfer-Encoding") {
			continue
		}

		if resp.StatusCode >= 200 {
			if strings.EqualFold(key, "Connection") {
				continue
			}

			if strings.EqualFold(key, "Keep-Alive") {
				continue
			}

			if strings.EqualFold(key, "Upgrade") {
				continue
			}

			if strings.EqualFold(key, "Proxy-Connection") {
				continue
			}
		}

		headers[key] = value
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (p *Proxy) handleTunneling(w http.ResponseWriter, r *http.Request) {
	log.Printf("CONNECT %v", r.Host)

	hijacker, ok := w.(http.Hijacker)

	if !ok {
		http.Error(w, "hijacking not supported", http.StatusInternalServerError)
		return
	}

	target, err := net.DialTimeout("tcp", r.Host, 10*time.Second)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)

	client, _, err := hijacker.Hijack()

	if err != nil {
		target.Close()

		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	transfer := func(dst io.WriteCloser, src io.ReadCloser) {
		defer dst.Close()
		defer src.Close()

		io.Copy(dst, src)
	}

	go transfer(target, client)
	go transfer(client, target)
}

func parseBasicAuth(r *http.Request) (username, password string, ok bool) {
	auth := r.Header.Get("Proxy-Authorization")

	if auth == "" {
		return
	}

	const prefix = "Basic "

	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return
	}

	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])

	if err != nil {
		return
	}

	cs := string(c)
	s := strings.IndexByte(cs, ':')

	if s < 0 {
		return
	}

	return cs[:s], cs[s+1:], true
}

func equal(a, b string) bool {
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
