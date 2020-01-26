package subscription

import (
	"encoding/base64"
	"fmt"
	"net"
)

const QuantumultClient = "quantumult"

type QuantumultVmess struct {
	Name            string // tag=Sample-H
	Group           string
	Addr            string // addr=ws-c.example.com:80
	UUID            string // password=23ad6b10-8d1a-40f7-8ad0-e3e35cd32291
	Security        string // method=chacha20-ietf-poly1305
	TLS             string
	TLSSecurity     bool
	Obfuscation     string // obfs=ws
	ObfuscationHost string // obfs-host=ws-c.example.com
	ObfuscationPath string // obfs-uri=/ws
	FastOpen        string // fast-open=false
	UDPRelay        string // udp-relay=false
}

func NewQuantumultVmess(group string, v Vmess) *QuantumultVmess {
	return &QuantumultVmess{
		Name:     v.Name,
		Group:    group,
		Addr:     net.JoinHostPort(v.Host, v.Port),
		UUID:     v.UUID,
		Security: v.Security,
		TLS: func() string {
			if v.TLS {
				return "true"
			}
			return ""
		}(),
		TLSSecurity: v.TLSSecurity,
		Obfuscation: func() string {
			if v.Obfuscation == "websocket" {
				return "ws"
			}
			return v.Obfuscation
		}(),
		ObfuscationHost: v.ObfuscationHost,
		ObfuscationPath: v.ObfuscationPath,
	}
}

const (
	UA = "User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_6 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) Mobile/15D100"
)

func (c QuantumultVmess) Build() ([]byte, error) {
	var result string
	host, port, err := net.SplitHostPort(c.Addr)
	if err != nil {
		return nil, err
	}
	result += fmt.Sprintf("%s = vmess, %s, %s, %s, \"%s\",group=%s", c.Name, host, port, c.Security, c.UUID, c.Group)
	if c.TLS != "" {
		result += fmt.Sprintf(", over-tls=%s", c.TLS)
	}
	if c.ObfuscationHost != "" {
		result += fmt.Sprintf(", tls-host=%s", host)
	}
	if c.TLSSecurity {
		result += ", certificate=1"
	}
	if c.Obfuscation != "" {
		result += fmt.Sprintf(", obfs=%s", c.Obfuscation)
	}
	if c.ObfuscationPath != "" {
		result += fmt.Sprintf(", obfs-path=\"%s\"", c.ObfuscationPath)
		result += fmt.Sprintf(", obfs-header=\"Host: %s[Rr][Nn]%s\"", c.ObfuscationHost, UA)
	}
	return []byte("vmess://" + base64.RawURLEncoding.EncodeToString([]byte(result))), nil
}
