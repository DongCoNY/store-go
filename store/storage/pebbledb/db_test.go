package pebbledb

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/log"
	"github.com/DongCoNY/store-go/store"
	"github.com/DongCoNY/store-go/store/storage"
)

func TestStorageTestSuite(t *testing.T) {
	s := &storage.StorageTestSuite{
		NewDB: func(dir string) (store.VersionedDatabase, error) {
			db, err := New(dir)
			if err == nil && db != nil {
				// We set sync=false just to speed up CI tests. Operators should take
				// careful consideration when setting this value in production environments.
				db.SetSync(false)
			}

			return storage.NewStorageStore(db, nil, log.NewNopLogger()), err
		},
		EmptyBatchSize: 12,
	}

	suite.Run(t, s)
}
