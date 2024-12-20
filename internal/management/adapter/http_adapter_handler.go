package adapter

import "net/http"

type HttpAdapterHandler map[string]func(w http.ResponseWriter, r *http.Request) error
