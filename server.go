package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// JSON
type Device struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	IP          string `json:"ip"`
	RoutingType string `json:"routing_type"`
}

const logFile = "network_devices.log"

func main() {
	http.HandleFunc("/", getDevices)
	http.HandleFunc("/add", addDevice)
	http.HandleFunc("/clear", clearLog)

	fmt.Println("Server running on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// GET
func getDevices(w http.ResponseWriter, r *http.Request) {
	data, _ := os.ReadFile(logFile)
	w.Write(data)
}

// POST
func addDevice(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var dev Device
		json.NewDecoder(r.Body).Decode(&dev)

		logEntry := fmt.Sprintf("Name: %s, Type: %s, IP: %s, Routing: %s\n",
			dev.Name, dev.Type, dev.IP, dev.RoutingType)

		f, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		f.WriteString(logEntry)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Added++")
	}
}

// DELETE: Очищення
func clearLog(w http.ResponseWriter, r *http.Request) {
	os.Remove(logFile)
	fmt.Fprint(w, "Log Cleared")
}
