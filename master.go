package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type Snap struct {
	conn net.Conn
	id   int
}

var snaps []Snap
var mu sync.Mutex
var snapIDCounter int = 1

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("The master device is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	addSnap(conn)
	defer removeSnap(conn)
	displaySnaps()

	for {

		var input string
		fmt.Println("Enter the ID of the Snap to send shutdown command, or 'exit' to quit:")
		fmt.Scanln(&input)

		if input == "exit" {
			break
		} else {

			sendShutdownToSnap(input)
			displaySnaps()
		}
	}
}

func sendShutdownToSnap(snapID string) {
	mu.Lock()
	defer mu.Unlock()

	for i, snap := range snaps {
		if fmt.Sprintf("%d", snap.id) == snapID {

			_, err := snap.conn.Write([]byte("shutdown"))
			if err != nil {
				log.Println("Error sending shutdown command:", err)
				return
			}
			logEvent(fmt.Sprintf("Shutdown command has been sent to Snap with ID: %d", snap.id))

			snaps = append(snaps[:i], snaps[i+1:]...)
			logEvent(fmt.Sprintf("Snap device with ID %d has been removed after shutdown.", snap.id))
			return
		}
	}

	fmt.Println("No Snap found with that ID!")
}

func addSnap(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()
	snap := Snap{conn: conn, id: snapIDCounter}
	snaps = append(snaps, snap)
	logEvent(fmt.Sprintf("Snap device with ID %d has connected and is online.", snapIDCounter))
	snapIDCounter++
}

func removeSnap(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	for i, snap := range snaps {
		if snap.conn == conn {
			logEvent(fmt.Sprintf("Snap device with ID %d has been disconnected.", snap.id))
			snaps = append(snaps[:i], snaps[i+1:]...)
			break
		}
	}
}

func displaySnaps() {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Currently online Snaps:")
	for _, snap := range snaps {
		fmt.Printf("ID: %d\n", snap.id)
	}
}

func logEvent(event string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(currentTime, event)
}
