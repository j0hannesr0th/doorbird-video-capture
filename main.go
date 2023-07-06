package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

// Mutex to prevent multiple records at the same time
var recordingMutex = &sync.Mutex{}

// Config is the structure of the config file.
type Config struct {
	Doorbird struct {
		IP       string `json:"ip"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"doorbird"`
	Record struct {
		Duration string `json:"duration"`
		Path     string `json:"path"`
	} `json:"record"`
}

func recordHandler(w http.ResponseWriter, r *http.Request, config Config) {
	// Try to lock the mutex, return if it's already locked
	recordingMutex.Lock()
	defer recordingMutex.Unlock()

	filename := fmt.Sprintf("%s/%s.mp4", config.Record.Path, time.Now().Format("20060102150405"))

	cmd := exec.Command("ffmpeg", "-y", "-i", fmt.Sprintf("http://%s:%s@%s/bha-api/video.cgi", config.Doorbird.Username, config.Doorbird.Password, config.Doorbird.IP), "-t", config.Record.Duration, "-c:v", "copy", filename)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Failed to start recording: %s\n", err)
		if len(output) > 0 {
			fmt.Printf("ffmpeg error output: %s\n", string(output))
		}
		return
	}

	fmt.Fprintf(w, "Recording saved to %s\n", filename)
}

func main() {
	// Load configuration
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error reading config file", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		recordHandler(w, r, config)
	})

	http.ListenAndServe(":8080", nil)
}
