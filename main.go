package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

type Server struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

// SetAlive for this backend
func (b *Server) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

// IsAlive returns true when backend is alive
func (b *Server) IsAlive() (alive bool) {
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return
}

// ServerPool holds information about reachable servers
type ServerPool struct {
	servers []*Server
	current uint64
}

// AddBackend to the server pool
func (s *ServerPool) AddServer(backend *Server) {
	s.servers = append(s.servers, backend)
}

// NextIndex atomically increase the counter and return an index
func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.servers)))
}

// MarkBackendStatus changes a status of a server
func (s *ServerPool) MarkServerStatus(backendUrl *url.URL, alive bool) {
	for _, b := range s.servers {
		if b.URL.String() == backendUrl.String() {
			b.SetAlive(alive)
			break
		}
	}
}

// GetNextServer returns next active server to take a connection in round robin fashion
func (s *ServerPool) GetNextServer() *Server {
	next := s.NextIndex()
	l := len(s.servers) + next
	for i := next; i < l; i++ {
		idx := i % len(s.servers)
		if s.servers[idx].IsAlive() {
			if i != next {
				atomic.StoreUint64(&s.current, uint64(idx))
			}
			return s.servers[idx]
		}
	}
	return nil
}
