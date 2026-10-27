package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var startTime = time.Now()

func uptimeHuman(seconds int) string {
	units := []struct {
		size int
		name string
	}{{86400, "day"}, {3600, "hour"}, {60, "minute"}, {1, "second"}}
	var parts []string
	for _, unit := range units {
		count := seconds / unit.size
		seconds %= unit.size
		if count > 0 {
			name := unit.name
			if count != 1 {
				name += "s"
			}
			parts = append(parts, fmt.Sprintf("%d %s", count, name))
		}
	}
	if len(parts) == 0 {
		return "0 seconds"
	}
	return strings.Join(parts, ", ")
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("JSON response: %v", err)
	}
}

func main() {
	hostname, _ := os.Hostname()
	platform := runtime.GOOS
	platformVersion := "unknown"
	architecture := runtime.GOARCH
	if value, err := exec.Command("uname", "-s").Output(); err == nil {
		platform = strings.TrimSpace(string(value))
	}
	if value, err := exec.Command("uname", "-v").Output(); err == nil {
		platformVersion = strings.TrimSpace(string(value))
	}
	if value, err := exec.Command("uname", "-m").Output(); err == nil {
		architecture = strings.TrimSpace(string(value))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Not found"})
			return
		}
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
			return
		}
		clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			clientIP = r.RemoteAddr
		}
		var userAgent any
		if value := r.Header.Get("User-Agent"); value != "" {
			userAgent = value
		}
		uptime := int(time.Since(startTime).Seconds())
		writeJSON(w, http.StatusOK, map[string]any{
			"service": map[string]any{"name": "DevOps Info Service", "version": "1.0.0", "description": "Reports service and host information", "framework": "net/http"},
			"system":  map[string]any{"hostname": hostname, "platform": platform, "platform_version": platformVersion, "architecture": architecture, "cpu_count": runtime.NumCPU(), "python_version": runtime.Version()},
			"runtime": map[string]any{"uptime_seconds": uptime, "uptime_human": uptimeHuman(uptime), "current_time": time.Now().UTC().Format(time.RFC3339Nano), "timezone": "UTC"},
			"request": map[string]any{"client_ip": clientIP, "user_agent": userAgent, "method": r.Method, "path": r.URL.Path},
			"endpoints": []map[string]string{
				{"path": "/", "method": "GET", "description": "Service and host information"},
				{"path": "/health", "method": "GET", "description": "Health check"},
			},
		})
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"status": "ok", "timestamp": time.Now().UTC().Format(time.RFC3339Nano), "uptime_seconds": int(time.Since(startTime).Seconds())})
	})
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	if number, err := strconv.Atoi(port); err != nil || number < 1 || number > 65535 {
		log.Fatal("Invalid PORT")
	}
	address := net.JoinHostPort(host, port)
	log.Printf("Starting DevOps Info Service on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
