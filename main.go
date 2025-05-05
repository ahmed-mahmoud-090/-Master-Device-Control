package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Snap_struct struct {
	ID   int
	Conn net.Conn
	Addr string
}

var (
	snaps_d      []Snap_struct
	snapsMutex sync.Mutex
	snapID     = 1
)

func main() {
	go startMasterTCPServer()

	http.HandleFunc("/", webHandler)
	http.HandleFunc("/send", sendCommandHandler)

	fmt.Println("Web GUI at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func startMasterTCPServer() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("TCP Server error:", err)
	}
	defer listener.Close()

	fmt.Println("Master TCP server listening on :9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		snapsMutex.Lock()
		snaps_d = append(snaps_d, Snap_struct{
			ID:   snapID,
			Conn: conn,
			Addr: conn.RemoteAddr().String(),
		})
		snapID++
		snapsMutex.Unlock()

		log.Printf("New Snap connected: %s\n", conn.RemoteAddr().String())
	}
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	snapsMutex.Lock()
	defer snapsMutex.Unlock()
	tmpl.Execute(w, snaps_d)
}

func sendCommandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	snapsMutex.Lock()
	defer snapsMutex.Unlock()

	for i, snap := range snaps_d {
		if snap.ID == id {
			_, err := snap.Conn.Write([]byte("shutdown"))
			if err != nil {
				http.Error(w, "Failed ,Not Found Connected Snaps!", http.StatusInternalServerError)
				return
			}

			snap.Conn.Close()
			snaps_d = append(snaps_d[:i], snaps_d[i+1:]...)
			break
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
