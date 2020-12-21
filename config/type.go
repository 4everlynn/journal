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

// Exchange exchange properties
type Exchange struct {
	Port     int
	Host     string
	Protocol string
	Name     string
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

func (scanner SampleScanner) ReadLine() string {
	line, _, _ := scanner.reader.ReadLine()
	return string(line)
}

func ScannerInstance() SampleScanner {
	scanner := SampleScanner{}
	scanner = scanner.init()
	return scanner
}
