package listmapper

import (
	"github.com/moogar0880/problems"
)

var (
	// InvalidQueryParameters is for query parameters which contained invalid data.
	InvalidQueryParameters = problems.NewDetailedProblem(400,
		"Could Not Process Request: An invalid query parameter was submitted")

	//InternalServerError is for internal server failures.
	InternalServerError = problems.NewDetailedProblem(500,
		"Internal Server Error")
)
