package postgres

import (
	"github.com/Dostonlv/NotCode.git/models"
	"github.com/Dostonlv/pkg/logger"
	_ "github.com/Dostonlv/pkg/logger"
	gopiston "github.com/milindmadhukar/go-piston"
)

type compilerRepo struct {
	db string
}

func (c *compilerRepo) Compile(req models.Req) (gopiston.PistonResponse, error) {
	var err error
	var res gopiston.PistonResponse

	client := gopiston.CreateDefaultClient()
	output, err := client.Execute(req.Language, req.Version, // Passing language. Since no version is specified, it uses the latest supported version.
		[]gopiston.Code{
			{Content: req.Code},
		},
		gopiston.Stdin(req.Cases), // Passing input as "hello world".
	)
	if err != nil {
		logger.Any("Error in code execute", err)
		return res, err
	}
	res = *output
	return res, nil
}

func NewCompilerRepo(db string) *compilerRepo {
	return &compilerRepo{
		db: db,
	}
}
