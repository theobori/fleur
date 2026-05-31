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
}

func NewOptions(port int, directoryPath string, domain string, verbose bool) (*Options, error) {
	if port < 0 {
		return nil, fmt.Errorf("The port must be positive.")
	}

	directoryAbsolutePath, err := filepath.Abs(directoryPath)
	if err != nil {
		return nil, err
	}

	return &Options{
		Port:                       port,
		DirectoryPath:              directoryAbsolutePath,
		Domain:                     domain,
		Verbose:                    verbose,
	}, nil
}

func (o *Options) DirectoryAbsolutePath() (string, error) {
	return filepath.Abs(o.DirectoryPath)
}
