package postgres

import (
	"github.com/Dostonlv/NotCode.git/models"
	"github.com/Dostonlv/pkg/logger"
	_ "github.com/Dostonlv/pkg/logger"
	gopiston "github.com/milindmadhukar/go-piston"
	"strings"
)

type compilerRepo struct {
	db string
}

func (c *compilerRepo) Compile(req models.Req) (string, error) {
	var err error
	client := gopiston.CreateDefaultClient()
	output, err := client.Execute(req.Language, req.Version,
		[]gopiston.Code{
			{Content: req.Code},
		},
		gopiston.Stdin(req.Cases),
	)
	out := output.GetOutput()
	out = strings.TrimSuffix(output.GetOutput(), "\n")
	if err != nil {
		logger.Any("Error in code execute", err)
		return out, err
	}
	if out == req.Test {
		out = "Test passed"
	} else {
		out = "Test failed"
	}
	return out, nil

}

func NewCompilerRepo(db string) *compilerRepo {
	return &compilerRepo{
		db: db,
	}
}
