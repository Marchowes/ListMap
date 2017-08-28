package routers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/moogar0880/problems"
)

// Abort writes passed in problems.StatusProblem to the Gin Response.
func Abort(c *gin.Context, problem problems.StatusProblem) {
	json, _ := json.Marshal(problem)
	c.Data(problem.ProblemStatus(), problems.ProblemMediaType, json)
	c.Abort()
}
