package sqlstore

import "go.mau.fi/util/dbutil"

func (s *SQLStore) GetDB() *dbutil.Database {
	return s.db
}
