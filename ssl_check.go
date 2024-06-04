package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type sslResult struct {
	Valid bool
	Error string
	Issuer string
	Subject string
	DNSNames []string
	SerialNumber string
	Version int
	ChainCertificates []string
	ExpirationDate string
	DaysUntilExpiration int
}

func CheckSSL(domain string) sslResult {
    conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", domain), &tls.Config{
        InsecureSkipVerify: true,
    })
    if err != nil {
        return sslResult{
            Valid: false,
            Error: err.Error(),
        }
    }
    defer conn.Close()

    certificates := conn.ConnectionState().PeerCertificates
    if len(certificates) == 0 {
        return sslResult{
            Valid: false,
            Error: "No certificates found",
        }
    }

    cert := certificates[0]

    var chainCertificates []string
    if len(conn.ConnectionState().VerifiedChains) > 0 {
        chain := conn.ConnectionState().VerifiedChains[0]
        for _, c := range chain {
            chainCertificates = append(chainCertificates, c.Issuer.CommonName)
        }
    }

	expirationDate := cert.NotAfter.Format("2006-01-02")
	daysUntilExpiration := int(time.Until(cert.NotAfter).Hours() / float64(24*time.Hour))

	return sslResult{
		Valid: true,
		Issuer: cert.Issuer.CommonName,
		Subject: cert.Subject.CommonName,
		DNSNames: cert.DNSNames,
		SerialNumber: cert.SerialNumber.String(),
		Version: cert.Version,
		ChainCertificates: chainCertificates,
		ExpirationDate: expirationDate,
		DaysUntilExpiration: daysUntilExpiration,
	}
}