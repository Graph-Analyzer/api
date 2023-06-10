package action

import (
	"encoding/json"
	"graph-analyzer/api/graph-property/domain"
	"graph-analyzer/api/graph-property/responder"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff/v3"
	"github.com/stretchr/testify/assert"
)

func TestQuery_Invoke_Cut_Edges(t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	service := domain.NewQueryService(createMockRepository(t))
	action := NewQueryAction(service)

	jsonParam, _ := json.Marshal(QueryRequestModel{
		Type: "cut_edges",
	})
	reader := strings.NewReader(string(jsonParam))

	request, _ := http.NewRequest(http.MethodPost, "/query", reader)
	ctx.Request = request

	action.Invoke(ctx)

	expectedResponse := responder.QueryProjection{
		Nodes: []responder.QueryNodeProjection{},
		Edges: []responder.QueryEdgeProjection{
			{
				From:      "1",
				FromLabel: "XR-1",
				To:        "4",
				ToLabel:   "XR-4",
			},
		},
	}

	var response responder.QueryProjection
	b, _ := io.ReadAll(w.Body)
	_ = json.Unmarshal(b, &response)

	assert.Equal(t, 200, w.Code)

	difference, err := diff.Diff(expectedResponse, response)
	assert.NoError(t, err)
	assert.Len(t, difference, 0)
}

func TestQuery_Invoke_Cut_Vertices(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	service := domain.NewQueryService(createMockRepository(t))
	action := NewQueryAction(service)

	jsonParam, _ := json.Marshal(QueryRequestModel{
		Type: "cut_vertices",
	})
	reader := strings.NewReader(string(jsonParam))

	request, _ := http.NewRequest(http.MethodPost, "/query", reader)
	ctx.Request = request

	action.Invoke(ctx)

	expectedResponse := responder.QueryProjection{
		Nodes: []responder.QueryNodeProjection{
			{
				Id:    "1",
				Label: "XR-1",
			},
			{
				Id:    "4",
				Label: "XR-4",
			},
		},
		Edges: []responder.QueryEdgeProjection{},
	}

	var response responder.QueryProjection
	b, _ := io.ReadAll(w.Body)
	_ = json.Unmarshal(b, &response)

	assert.Equal(t, 200, w.Code)

	difference, err := diff.Diff(expectedResponse, response)
	assert.NoError(t, err)
	assert.Len(t, difference, 0)
}
