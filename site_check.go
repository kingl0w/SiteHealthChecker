package main

import (
	"fmt"
	"net/http"
	"time"
)

type HealthCheckResult struct {
    Domain string `json:"domain"`
    Port string `json:"port"`
    Status string `json:"status"`
    StatusCode int `json:"status_code,omitempty"`
    ResponseTime int64 `json:"responseTime,omitempty"`
    Error string `json:"error,omitempty"`
}

func Check(domain, port string) HealthCheckResult {
    url := fmt.Sprintf("http://%s:%s", domain, port)
    startTime := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        return HealthCheckResult{
            Domain: domain,
            Port: port,
            Status: "DOWN",
            Error: err.Error(),
            ResponseTime: time.Since(startTime).Milliseconds(),
        }
    }
    defer resp.Body.Close()

    status := "UP"
    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        status = "DOWN"
    }

    return HealthCheckResult{
        Domain: domain,
        Port: port,
        Status: status,
        StatusCode: resp.StatusCode,
        ResponseTime: time.Since(startTime).Milliseconds(),
    }
}