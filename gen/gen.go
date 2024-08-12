package main

import (
	"bytes"
	"flag"
	"go/format"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pelletier/go-toml/v2"
)

type flags struct {
	ConfigFile string
	BasePath   string
}

type Contract struct {
	Name  string
	Index int
}

type Golang struct {
	Package string
}

type System struct {
	ShutterStartAddress string
	Contracts           []Contract
}

type Config struct {
	System System
	Golang Golang
}

func main() {
	var f flags
	flag.StringVar(&f.ConfigFile, "config", "", "toml configuration file")
	flag.StringVar(&f.BasePath, "base-path", "", "base path (usually the project root)")
	flag.Parse()

	var cfg Config
	cfgFile, err := os.ReadFile(f.ConfigFile)
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(cfgFile, &cfg)
	if err != nil {
		panic(err)
	}
	// TODO: check that the addresses fall in the range of the
	// Optimism predeploy namespace! (not 100% sure, but it needs to have
	// the 0x420000000... prefix [:18] and the remainder has to be unique etc.)
	if !common.IsHexAddress(cfg.System.ShutterStartAddress) {
		panic("start address must be an ethereum hex address")
	}

	startAddr := common.HexToAddress(cfg.System.ShutterStartAddress)

	t := template.Must(template.New("artifact").Funcs(template.FuncMap{
		"toPrivVar": func(name string) string {
			if name == "" {
				return ""
			}
			return strings.Join([]string{strings.ToLower(name[:1]), name[1:]}, "")
		},
		"getContractAddress": func(index int) string {
			addr := new(big.Int)
			addr.Add(startAddr.Big(), big.NewInt(int64(index)))
			return common.BigToAddress(addr).Hex()
		},
	}).Parse(tmpl))

	fname := filepath.Join(f.BasePath, strings.ToLower(cfg.Golang.Package), "contracts.gen.go")
	outfile, err := os.OpenFile(
		fname,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0o600,
	)
	if err != nil {
		log.Fatalf("error opening %s: %v\n", fname, err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, cfg); err != nil {
		log.Fatalf("error writing template %s: %v", outfile.Name(), err)
	}
	p, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("error formatting generated code: %v\n", err)
	}
	_, err = outfile.Write(p)
	if err != nil {
		log.Fatalf("error writing to %s: %v\n", fname, err)
	}
	outfile.Close()
	log.Printf("wrote file %s\n", outfile.Name())
	if err != nil {
		panic(err)
	}
}

var tmpl = `// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.Golang.Package}}

import "github.com/ethereum/go-ethereum/common"

var (
	{{range .System.Contracts -}}
	{{$privvar:=toPrivVar .Name }}
	{{$privvar}}            = "{{getContractAddress .Index}}"
	{{.Name}}Addr           = common.HexToAddress({{$privvar}})

	{{- end}}

	Predeploys = make(map[string]*common.Address)
	PredeploysAddrSet = make(map[common.Address]string)
)

func init() {
	{{range .System.Contracts}}
	Predeploys["{{.Name}}"] = &{{.Name}}Addr
	{{- end}}

	for name, addr := range Predeploys{
		PredeploysAddrSet[*addr] = name
	}
}
`
