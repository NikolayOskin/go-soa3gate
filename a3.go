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

// NewConfig: config constructor
func NewConfig(trustCertPath string, pemCertPath string, authKey string, isProd bool) *config {
	return &config{
		trustCertPath: trustCertPath,
		pemCertPath:   pemCertPath,
		authKey:       authKey,
		isProd:        isProd,
	}
}

// NewA3: SOA3Gate lib constructor
func NewA3(c *config) (*A3, error) {
	client, err := newA3Client(c.trustCertPath, c.pemCertPath)
	if err != nil {
		return nil, err
	}
	return &A3{
		client,
		c,
	}, nil
}

func newA3Client(trustCertPath string, pemCertPath string) (*http.Client, error) {
	cert, err := tls.LoadX509KeyPair(trustCertPath, pemCertPath)
	if err != nil {
		return nil, err
	}

	ssl := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: false,
	}
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: ssl,
		},
	}, nil
}
