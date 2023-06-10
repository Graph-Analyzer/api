package action

import (
	"encoding/json"
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff/v3"
	"github.com/stretchr/testify/assert"
)

func TestGeneralInformation_Invoke(t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	service := domain.NewGeneralInformationService(createMockRepository(t))
	action := NewGeneralInformationAction(service)

	action.Invoke(ctx)

	expectedResponse := responder.GeneralInformationProjection{
		Name:      "test",
		Nodes:     6,
		Edges:     14,
		CreatedAt: "2022-11-30T13:59:59Z",
	}

	var response responder.GeneralInformationProjection
	b, _ := io.ReadAll(w.Body)
	_ = json.Unmarshal(b, &response)

	assert.Equal(t, 200, w.Code)

	difference, err := diff.Diff(expectedResponse, response)
	assert.NoError(t, err)
	assert.Len(t, difference, 0)
}
