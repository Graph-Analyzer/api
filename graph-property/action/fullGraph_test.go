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

func TestFullGraph_Invoke(t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	service := domain.NewFullGraphService(createMockRepository(t))
	action := NewFullGraphAction(service)

	action.Invoke(ctx)

	expectedResponse := responder.FullGraphProjection{
		Nodes: []responder.FullGraphNodeProjection{
			{
				Id:    "1",
				Label: "XR-1",
				Size:  3,
			},
			{
				Id:    "2",
				Label: "XR-2",
				Size:  2,
			},
			{
				Id:    "3",
				Label: "XR-3",
				Size:  2,
			},
			{
				Id:    "4",
				Label: "XR-4",
				Size:  3,
			},
			{
				Id:    "5",
				Label: "XR-5",
				Size:  2,
			},
			{
				Id:    "6",
				Label: "XR-6",
				Size:  2,
			},
		},
		Edges: []responder.FullGraphEdgeProjection{
			{
				Id:     "1",
				From:   "1",
				To:     "2",
				Weight: 1,
			},
			{
				Id:     "2",
				From:   "2",
				To:     "1",
				Weight: 1,
			},
			{
				Id:     "3",
				From:   "2",
				To:     "3",
				Weight: 1,
			},
			{
				Id:     "4",
				From:   "3",
				To:     "2",
				Weight: 1,
			},
			{
				Id:     "5",
				From:   "3",
				To:     "1",
				Weight: 1,
			},
			{
				Id:     "6",
				From:   "1",
				To:     "3",
				Weight: 1,
			},
			{
				Id:     "7",
				From:   "4",
				To:     "5",
				Weight: 1,
			},
			{
				Id:     "8",
				From:   "5",
				To:     "4",
				Weight: 1,
			},
			{
				Id:     "9",
				From:   "5",
				To:     "6",
				Weight: 1,
			},
			{
				Id:     "10",
				From:   "6",
				To:     "5",
				Weight: 1,
			},
			{
				Id:     "11",
				From:   "6",
				To:     "4",
				Weight: 1,
			},
			{
				Id:     "12",
				From:   "4",
				To:     "6",
				Weight: 1,
			},
			{
				Id:     "13",
				From:   "1",
				To:     "4",
				Weight: 1,
			},
			{
				Id:     "14",
				From:   "4",
				To:     "1",
				Weight: 1,
			},
		},
	}

	var response responder.FullGraphProjection
	b, _ := io.ReadAll(w.Body)
	_ = json.Unmarshal(b, &response)

	assert.Equal(t, 200, w.Code)

	difference, err := diff.Diff(expectedResponse, response)
	assert.NoError(t, err)
	assert.Len(t, difference, 0)
}
