package action

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLiveness_Invoke(t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	action := NewLivenessAction()

	action.Invoke(ctx)

	assert.Equal(t, 200, w.Code)
}
