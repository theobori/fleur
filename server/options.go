package server

import (
	"fmt"
	"path/filepath"
)

type Options struct {
	// Gopher port
	Port int
	// Directory path, that will be the root of the Gopher server
	DirectoryPath string
	// Gopher site domain
	Domain string
	// Enable verbose
	Verbose bool
	// Enable Gopher over TLS
	EnableTLS bool
	// Certificate path used during TLS communication
	CertificatePath string
	// Certificate path used during TLS communication
	KeyPath string
}

func NewOptions(port int, directoryPath string, domain string, verbose bool, enableTLS bool, certificatePath string, keyPath string) (*Options, error) {
	if port < 0 {
		return nil, fmt.Errorf("The port must be positive.")
	}

	directoryAbsolutePath, err := filepath.Abs(directoryPath)
	if err != nil {
		return nil, err
	}

	return &Options{
		Port:            port,
		DirectoryPath:   directoryAbsolutePath,
		Domain:          domain,
		Verbose:         verbose,
		EnableTLS:       enableTLS,
		CertificatePath: certificatePath,
		KeyPath:         keyPath,
	}, nil
}

func (o *Options) DirectoryAbsolutePath() (string, error) {
	return filepath.Abs(o.DirectoryPath)
}

func (o *Options) CertificateAbsolutePath() (string, error) {
	return filepath.Abs(o.CertificatePath)
}

func (o *Options) KeyAbsolutePath() (string, error) {
	return filepath.Abs(o.KeyPath)
}
