package postgres

import (
	"github.com/Dostonlv/NotCode.git/config"
	"github.com/Dostonlv/NotCode.git/storage"
)

type Store struct {
	db   string
	comp storage.CompilerRepoI
}

func (s *Store) Compiler() storage.CompilerRepoI {
	if s.comp == nil {
		s.comp = NewCompilerRepo(s.db)
	}
	return s.comp
}

func NewConnectPostgresql(cfg *config.Config) (storage.StorageI, error) {

	return &Store{
		db:   cfg.Environment,
		comp: NewCompilerRepo("s"),
	}, nil
}
