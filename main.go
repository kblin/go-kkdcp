package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kblin/go-kkdcp/codec"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

func HandleKkdcp(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/kerberos")

	defer req.Body.Close()
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Error reading body")
		return
	}
	proxy_request, err := codec.Decode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Error decoding body")
		return
	}

	krb5_response, err := ForwardKerberos(proxy_request.Message)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Fatal("Connection to Kerberos upstream failed")
		return
	}

	reply, err := codec.Encode(krb5_response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Encoding message failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(reply)
	if err != nil {
		return
	}
}

func ForwardKerberos(data []byte) (resp []byte, err error) {
	conn, err := net.Dial("tcp", "kdc.demo.kblin.org:88")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	resp, err = ioutil.ReadAll(conn)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return resp, nil
}

func main() {
	router := mux.NewRouter()
	router.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(HandleKkdcp))).Methods("POST")
	log.Fatal(http.ListenAndServe(":8124", router))
}
