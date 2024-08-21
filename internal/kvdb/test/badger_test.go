package test

import (
	"testing"

	"github.com/Agelessbaby/Search-Enginee/internal/kvdb"
	"github.com/Agelessbaby/Search-Enginee/util"
)

func TestBadger(t *testing.T) {
	setup = func() {
		var err error
		db, err = kvdb.GetKvDb(kvdb.BADGER, util.RootPath+"data/badger_db")
		if err != nil {
			panic(err)
		}
	}

	t.Run("badger_test", testPipeline)
}

// go test -v ./internal/kvdb/test -run=^TestBadger$ -count=1
