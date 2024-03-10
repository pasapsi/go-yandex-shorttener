package main

import (
	b64 "encoding/base64"
	"io"
	"net/http"
	"strings"
)

func mainPOST(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, _ := io.ReadAll(r.Body)
		host := r.Host
		sEnc := b64.StdEncoding.EncodeToString([]byte(body))

		w.WriteHeader(201)
		w.Header().Set("content-type", "text/plain")
		w.Write([]byte(host + "/" + sEnc))
		return

		//sEnc := b64.StdEncoding.EncodeToString([]byte(shortId))
		// пока установим ответ-заглушку, без проверки ошибок
		//_, _ = w.Write([]byte(`
		//{
		//	"response":
		//`))
	} else {
		w.WriteHeader(400)
	}
}

func mainGET(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//host := r.Host
		shortId := strings.Split(r.URL.Path, "/")[1]
		decode, _ := b64.StdEncoding.DecodeString(shortId)
		w.Header().Set("content-type", "text/plain")
		w.Header().Set("Location", string(decode))
		//Enc := b64.StdEncoding.EncodeToString([]byte(shortId))
		//fmt.Fprintf(w, "Hello, %q", Enc)
		w.WriteHeader(307)
		return
	} else {
		w.WriteHeader(400)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/{id}`, mainGET)
	mux.HandleFunc(`/`, mainPOST)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
