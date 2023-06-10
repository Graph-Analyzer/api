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

func TestDegreeCorrelation_Invoke(t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	service := domain.NewDegreeCorrelationService(createMockRepository(t))
	action := NewDegreeCorrelationAction(service)

	action.Invoke(ctx)

	expectedResponse := responder.DegreeCorrelationProjection{
		Values: []responder.DegreeCorrelationValueProjection{
			{
				Degree:  2,
				Average: 2.5,
			},
			{
				Degree:  3,
				Average: 2.33333,
			},
		},
	}

	var response responder.DegreeCorrelationProjection
	b, _ := io.ReadAll(w.Body)
	_ = json.Unmarshal(b, &response)

	response.Values[1].Average = math.Round(response.Values[1].Average*100000) / 100000

	assert.Equal(t, 200, w.Code)

	difference, err := diff.Diff(expectedResponse, response)
	assert.NoError(t, err)
	assert.Len(t, difference, 0)
}
