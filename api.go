package main

type APIServer struct {
	listenAddr string
}

// Returning a pointer into our APIServer struct
func NewAPIServer(listenAddr string) *APIServer {

	return &APIServer{
		listenAddr: listenAddr,
	}
	
}

