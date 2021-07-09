package config

import (
	"bufio"
	"os"
)

// Git to describe Repo's structure
type Git struct {
	Name    string
	Path    string
	Disable bool
}

// JournalConfig to describe .journal.yaml
type JournalConfig struct {
	Version    string
	Maintainer string
	Git        map[string]Git
}

// Version app version
type Version struct {
	Master int
	Mirror int
	Patch  int
}

// Exchange exchange properties
type Exchange struct {
	Port     int
	Host     string
	Protocol string
	Name     string
}

// Func function
type Func struct {
	Name string
	Args []interface{}
}

type ExchangeContext interface {
	Write(buffer []byte)
	Shutdown()
}

type Scanner interface {
	init() SampleScanner
	ReadLine() string
}

type SampleScanner struct {
	reader *bufio.Reader
}

func (scanner SampleScanner) init() SampleScanner {
	reader := bufio.NewReader(os.Stdin)
	scanner.reader = reader
	return scanner
}

// ReadLine read a line of input from the terminal
func (scanner SampleScanner) ReadLine() string {
	line, _, _ := scanner.reader.ReadLine()
	return string(line)
}

// ScannerInstance get instance of the scanner
func ScannerInstance() SampleScanner {
	scanner := SampleScanner{}
	scanner = scanner.init()
	return scanner
}
