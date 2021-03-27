package main

import (
	"crypto/tls"
	"net/http"
)

type A3 struct {
	client *http.Client
	config *config
}

type config struct {
	trustCertPath string
	pemCertPath   string
	authKey       string
	isProd        bool
}

func NewConfig(trustCertPath string, pemCertPath string, authKey string) *config {
	return &config{
		trustCertPath: trustCertPath,
		pemCertPath:   pemCertPath,
		authKey:       authKey,
	}
}

func NewA3(c *config) *A3 {
	return &A3{
		newA3Client(c.trustCertPath, c.pemCertPath),
		c,
	}
}

func newA3Client(trustCertPath string, pemCertPath string) *http.Client {
	cert, _ := tls.LoadX509KeyPair(trustCertPath, pemCertPath)
	ssl := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: false,
	}
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: ssl,
		},
	}
}
