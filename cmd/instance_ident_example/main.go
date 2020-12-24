package main

import (
	"log"

	instanceident "github.com/nangantata/go-instanceident"
)

func main() {
	processIdent, err := instanceident.ProcessDerivedStringIdentity()
	if nil != err {
		log.Fatalf("having process derived failed: %v", err)
	}
	log.Printf("ProcessDerivedStringIdentity: %s", processIdent)
	timeIdent := instanceident.TimeDerivedInt64Identity()
	log.Printf("ProcessDerivedStringIdentity: 0x%016X (%d)", uint64(timeIdent), timeIdent)
}
