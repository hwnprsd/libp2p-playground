package node

import "net/http"

func (n *Node) SetupHttpServer() {
	http.HandleFunc("/keygen", func(w http.ResponseWriter, r *http.Request) {
	})
}
