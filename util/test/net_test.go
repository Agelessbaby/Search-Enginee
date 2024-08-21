package test

import (
	"fmt"
	"testing"

	"github.com/Agelessbaby/Search-Enginee/util"
)

func TestGetLocalIP(t *testing.T) {
	fmt.Println(util.GetLocalIP())
}

// go test -v ./util/test -run=^TestGetLocalIP$ -count=1
