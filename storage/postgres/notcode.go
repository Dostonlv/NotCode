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

func (c *compilerRepo) Compile(req models.Req, test models.Test) (string, error) {
	test.Inputs[0] = "Hello"
	test.Outputs[0] = "olleH"
	test.Inputs[1] = "ArrA"
	test.Outputs[1] = "ArrA"
	test.Inputs[2] = "12345"
	test.Outputs[2] = "54321"
	var tests [3]string
	var err error
	client := gopiston.CreateDefaultClient()

	// for 1-tests
	output, err := client.Execute(req.Language, req.Version,
		[]gopiston.Code{
			{Content: req.Code},
		},
		gopiston.Stdin(test.Inputs[0]),
	)
	out := output.GetOutput()
	out = strings.TrimSuffix(out, "\n")
	tests[0] = out

	// for 2-tests
	output, err = client.Execute(req.Language, req.Version,
		[]gopiston.Code{
			{Content: req.Code},
		},
		gopiston.Stdin(test.Inputs[1]),
	)
	out = output.GetOutput()
	out = strings.TrimSuffix(out, "\n")
	tests[1] = out

	// for 3-tests
	output, err = client.Execute(req.Language, req.Version,
		[]gopiston.Code{
			{Content: req.Code},
		},
		gopiston.Stdin(test.Inputs[2]),
	)
	out = output.GetOutput()
	out = strings.TrimSuffix(out, "\n")
	tests[2] = out

	if err != nil {
		logger.Any("Error in code execute", err)
		return out, err
	}

	if tests == test.Outputs {
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
