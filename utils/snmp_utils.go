package utils

import (
	"path/filepath"
	"strings"

	g "github.com/gosnmp/gosnmp"
)

type Auth struct {
	Protocol g.SnmpV3AuthProtocol `json:"protocol"` // TODO Custom deser parseAuth
	Pass     string               `json:"pass"`
}

type Priv struct {
	Protocol g.SnmpV3PrivProtocol `json:"protocol"` // TODO Custom deser parsePriv
	Pass     string               `json:"pass"`
}
type V3Config struct {
	Auth     Auth   `json:"auth"`
	Priv     Priv   `json:"priv"`
	Engineid string `json:"engineid"`
	Username string `json:"username"`
}

func parseAuth(authprotocol string) g.SnmpV3AuthProtocol {
	/*
		const (
			NoAuth SnmpV3AuthProtocol = 1
			MD5    SnmpV3AuthProtocol = 2
			SHA    SnmpV3AuthProtocol = 3
			SHA224 SnmpV3AuthProtocol = 4
			SHA256 SnmpV3AuthProtocol = 5
			SHA384 SnmpV3AuthProtocol = 6
			SHA512 SnmpV3AuthProtocol = 7
		)
	*/
	s := strings.Replace(authprotocol, "-", "", -1)
	switch s {
	case "noauth", "NoAuth", "noAuth":
		return g.NoAuth
	case "md5", "MD5":
		return g.MD5
	case "sha", "SHA":
		return g.SHA
	case "sha224", "SHA224":
		return g.SHA224
	case "sha256", "SHA256":
		return g.SHA256
	case "sha384", "SHA384":
		return g.SHA384
	case "sha512", "SHA512":
		return g.SHA512
	default:
		return g.SHA512
	}
}

func parsePriv(privprotocol string) g.SnmpV3PrivProtocol {
	/*
		const (
			NoPriv  SnmpV3PrivProtocol = 1
			DES     SnmpV3PrivProtocol = 2
			AES     SnmpV3PrivProtocol = 3
			AES192  SnmpV3PrivProtocol = 4 // Blumenthal-AES192
			AES256  SnmpV3PrivProtocol = 5 // Blumenthal-AES256
			AES192C SnmpV3PrivProtocol = 6 // Reeder-AES192
			AES256C SnmpV3PrivProtocol = 7 // Reeder-AES256
		)
	*/
	s := strings.Replace(privprotocol, "-", "", -1)
	switch s {
	case "nopriv", "NoPriv", "noPriv":
		return g.NoPriv
	case "des", "DES":
		return g.DES
	case "aes", "AES":
		return g.AES
	case "aes192", "AES192":
		return g.AES192
	case "aes256", "AES256":
		return g.AES256
	case "aes192c", "AES192C":
		return g.AES192C
	case "aes256c", "AES256C":
		return g.AES256C
	}
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
