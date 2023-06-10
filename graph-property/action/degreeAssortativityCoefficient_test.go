package action

import (
	"encoding/json"
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	"io"
	"math"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff/v3"
	"github.com/stretchr/testify/assert"
)

func TestDegreeAssortativityCoefficient_Invoke(t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	service := domain.NewDegreeAssortativityCoefficientService(createMockRepository(t))
	action := NewDegreeAssortativityCoefficientAction(service)

	action.Invoke(ctx)

	expectedResponse := responder.DegreeAssortativityCoefficientProjection{
		Value: -0.16667,
	}

	var response responder.DegreeAssortativityCoefficientProjection
	b, _ := io.ReadAll(w.Body)
	_ = json.Unmarshal(b, &response)

	response.Value = math.Round(response.Value*100000) / 100000

	assert.Equal(t, 200, w.Code)

	difference, err := diff.Diff(expectedResponse, response)
	assert.NoError(t, err)
	assert.Len(t, difference, 0)
}
