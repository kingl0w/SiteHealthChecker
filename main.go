package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Name:  "HealthChecker",
        Usage: "Check if the service is healthy",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:     "domain",
                Aliases:  []string{"d"},
                Usage:    "The URL to check",
                Required: true,
            },
            &cli.StringFlag{
                Name:     "ports",
                Aliases:  []string{"p"},
                Usage:    "The ports to check (comma-separated)",
            },
        },
        Action: func(c *cli.Context) error {
            domain := c.String("domain")
            ports := c.String("ports")
			defaultPorts :="80,443"

			var portList []string
			if ports != "" {
				portList = append(strings.Split(defaultPorts, ","), strings.Split(ports, ",")...)
			} else {
				portList = strings.Split(defaultPorts, ",")
			}

			websiteUp := false
            for _, port := range portList {
                result := Check(domain, strings.TrimSpace(port))
                fmt.Printf("Domain: %s\n", result.Domain)
                fmt.Printf("Port: %s\n", result.Port)
                fmt.Printf("Status: %s\n", result.Status)
                fmt.Printf("StatusCode: %d\n", result.StatusCode)
                fmt.Printf("ResponseTime: %d ms\n", result.ResponseTime)
                if result.Error != "" {
                    fmt.Printf("Error: %s\n", result.Error)
                }
                fmt.Println("-----------------------------------------------")

				if (result.Port == "80" || result.Port == "443") && result.StatusCode >= 200 && result.StatusCode < 300 {
					websiteUp = true
				}
            }

			if websiteUp {
                sslResult := CheckSSL(domain)
                fmt.Printf("SSL Certificate Check:\n")
                fmt.Printf("Certificate Valid: %v\n", sslResult.Valid)
                if !sslResult.Valid {
                    fmt.Printf("Error: %s\n", sslResult.Error)
                } else {
                    fmt.Printf("Issuer: %s\n", sslResult.Issuer)
                    fmt.Printf("Subject: %s\n", sslResult.Subject)
                    fmt.Printf("DNS Names:\n")
                    for _, dnsName := range sslResult.DNSNames {
                        fmt.Printf("    - %s\n", dnsName)
                    }
                    fmt.Printf("Serial Number: %s\n", sslResult.SerialNumber)
                    fmt.Printf("Version: %d\n", sslResult.Version)
                    fmt.Printf("Chain Certificates: %v\n", sslResult.ChainCertificates)
                    fmt.Printf("Expiration Date: %s\n", sslResult.ExpirationDate)
                    fmt.Printf("Days Until Expiration: %d\n", sslResult.DaysUntilExpiration)
                }
                fmt.Println("-----------------------------------------------")
            }

			fmt.Printf("Website %s is %s\n", domain, getWebsiteStatus(websiteUp))
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func getWebsiteStatus(up bool) string {
	if up {
		return "UP"
	}
	return "DOWN"
}