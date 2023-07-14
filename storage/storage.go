package storage

import (
	"github.com/Dostonlv/NotCode.git/models"
)

type StorageI interface {
	Compiler() CompilerRepoI
}

type CompilerRepoI interface {
	Compile(req models.Req) (string, error)
}
