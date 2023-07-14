package storage

import (
	"github.com/Dostonlv/NotCode.git/models"
	gopiston "github.com/milindmadhukar/go-piston"
)

type StorageI interface {
	Compiler() CompilerRepoI
}

type CompilerRepoI interface {
	Compile(req models.Req) (gopiston.PistonResponse, error)
}
