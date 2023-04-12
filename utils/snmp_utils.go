package utils

import (
	g "github.com/gosnmp/gosnmp"
	"path/filepath"
)

type Auth struct {
	Protocol g.SnmpV3AuthProtocol
	Pass     string
}

type Priv struct {
	Protocol g.SnmpV3PrivProtocol
	Pass     string
}
type V3Config struct {
	Auth     Auth
	Priv     Priv
	Engineid string
	Username string
}

func fromJson(fname string) (*V3Config, error) {
	// TODO
	return nil, nil
}

func fromYaml(fname string) (*V3Config, error) {
	// TODO
	return nil, nil
}

func NewV3ConfigFromFile(fname string) (*V3Config, error) {
	extension := filepath.Ext(fname)
	switch extension {
	case "json":
		return fromJson(fname)
	case "yaml", "yml":
		return fromYaml(fname)
	default:
		return fromJson(fname)
	}
}

func NewV3ConfigFromEnv() (*V3Config, error) {
	// TODO Parse from env
	return nil, nil
}
