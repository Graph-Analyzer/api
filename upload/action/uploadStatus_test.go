package action

import (
	"encoding/json"
	mocks "graph-analyzer/api/mocks/upload/domain"
	"graph-analyzer/api/upload/domain"
	"graph-analyzer/api/upload/responder"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/r3labs/diff/v3"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUploadStatus_Invoke(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("should return status 200 when successful check", func(t *testing.T) {
		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		ctx, _ := gin.CreateTestContext(w)

		mockService := mocks.NewUploadStatusService(t)
		mockService.On("Invoke").Return(domain.UploadStatusDTO{
			Healthy: true,
		}, nil).Maybe()
		action := NewUploadStatusAction(mockService)

		action.Invoke(ctx)

		expectedResponse := responder.UploadStatusProjection{
			Healthy: true,
		}

		var response responder.UploadStatusProjection
		b, _ := io.ReadAll(w.Body)
		_ = json.Unmarshal(b, &response)

		assert.Equal(t, 200, w.Code)

		difference, err := diff.Diff(expectedResponse, response)
		assert.NoError(t, err)
		assert.Len(t, difference, 0)
	})
}
