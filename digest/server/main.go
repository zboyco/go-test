package main

import (
	"fmt"
	"io"
	"net/http"

	auth "github.com/abbot/go-http-auth"
)

func secret(user, realm string) string {
	if user == "john" {
		// password is "hello"
		return "b98e16cbc3d01734b264adba7baa3bf9"
	}
	return ""
}

func push(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	body, _ := io.ReadAll(r.Body)
	fmt.Println(string(body))

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"resultCode":0,"resultDesc":"success","snapUploadUrl":"http://172.31.2.103:8080/upload"}`))
}

func upload(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	body, _ := io.ReadAll(r.Body)
	fmt.Println(len(body))

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"resultCode":0}`))
}

func main() {
	// authenticator := auth.NewDigestAuthenticator("example.com", secret)
	// http.HandleFunc("/test", authenticator.Wrap(push))
	// http.HandleFunc("/upload", authenticator.Wrap(upload))
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			fmt.Println("no auth header")
			w.Header().Set("WWW-Authenticate", `Digest realm="example.com", nonce="isAAXN6jpIcjmWwx", opaque="+7VavjSZAe+Vtqdj", algorithm="MD5", qop="auth"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println(auth)

		body, _ := io.ReadAll(r.Body)
		fmt.Println(string(body))

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"resultCode":0,"resultDesc":"success","snapUploadUrl":"http://172.31.2.103:8080/upload"}`))
	})
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			fmt.Println("no auth header")
			w.Header().Set("WWW-Authenticate", `Digest realm="example.com", nonce="isAAXN6jpIcjmWwx", opaque="+7VavjSZAe+Vtqdj", algorithm="MD5", qop="auth"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println(auth)

		body, _ := io.ReadAll(r.Body)
		fmt.Println(len(body))

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"resultCode":0}`))
	})
	http.ListenAndServe(":8080", nil)
}
