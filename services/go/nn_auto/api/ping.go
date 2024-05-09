package api

import "net/http"

func (a *API) pingHandler() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("nn-auto is alive"))
		if err != nil {
			return
		}
	}
}
