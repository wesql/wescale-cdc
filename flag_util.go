package cdc

import (
	"flag"
	"os"
)

func RegisterStringVar(p *string, name string, value string, usage string) {
	flag.StringVar(p, name, value, usage)
	val := os.Getenv(name)
	if val != "" {
		flag.Set(name, val)
	}
}
