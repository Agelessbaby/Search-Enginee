package test

import (
	"testing"

	"github.com/Agelessbaby/Search-Enginee/internal/kvdb"
	"github.com/Agelessbaby/Search-Enginee/util"
)

func TestBolt(t *testing.T) {
	setup = func() {
		var err error
		db, err = kvdb.GetKvDb(kvdb.BOLT, util.RootPath+"data/bolt_db") //使用工厂模式
		if err != nil {
			panic(err)
		}
	}

	t.Run("bolt_test", testPipeline)
}

// go test -v ./internal/kvdb/test -run=^TestBolt$ -count=1
