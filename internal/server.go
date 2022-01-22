package internal

import (
	"context"
	"golang.org/x/net/proxy"
	"io"
	"net"
	"net/http"
	"time"
)

const HTTP200 = "HTTP/1.1 200 Connection Established\r\n\r\n"
const CONNECT = "CONNECT"

type Server struct {
	Dialer        proxy.Dialer
	ListenAddress string
	SOCKSAddress  string
	Timeout       int
}

// handleHTTP - Proxy HTTP requests
func (s *Server) handleHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	transport := http.Transport{
		DialContext: s.Dialer.(interface {
			DialContext(context context.Context, network, address string) (net.Conn, error)
		}).DialContext,
	}

	response, err := transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	copyHeader(responseWriter.Header(), response.Header)
	responseWriter.WriteHeader(response.StatusCode)
	_, _ = io.Copy(responseWriter, response.Body)
}

//handleTunnel - Proxy CONNECT requests
func (s *Server) handleTunnel(responseWriter http.ResponseWriter, request *http.Request) {
	hijacker, ok := responseWriter.(http.Hijacker)
	if !ok {
		return
	}

	sourceConnection, _, err := hijacker.Hijack()
	if err != nil {
		return
	}

	destinationConnection, err := s.Dialer.Dial("tcp", request.Host)
	if err != nil {
		_ = sourceConnection.Close()
		return
	}

	_, _ = sourceConnection.Write([]byte(HTTP200))

	go copyIO(sourceConnection, destinationConnection)
	go copyIO(destinationConnection, sourceConnection)
}

//Run - Proxy all the things
func (s Server) Run() {
	_ = http.ListenAndServe(s.ListenAddress, http.HandlerFunc(
		func(responseWriter http.ResponseWriter, request *http.Request) {
			d := &net.Dialer{
				Timeout: time.Duration(s.Timeout) * time.Second,
			}

			s.Dialer, _ = proxy.SOCKS5("tcp", s.SOCKSAddress, nil, d)

			if request.Method == CONNECT {
				s.handleTunnel(responseWriter, request)
			} else {
				s.handleHTTP(responseWriter, request)
			}
		},
	))
}
