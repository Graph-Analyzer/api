package action

import (
	"bytes"
	"encoding/json"
	mocks "graph-analyzer/api/mocks/upload/domain"
	"graph-analyzer/api/upload/domain"
	"graph-analyzer/api/upload/responder"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/r3labs/diff/v3"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpload_Invoke(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("should return status 200 when successful upload", func(t *testing.T) {
		tmpfile, err := os.CreateTemp("", "*graph.gexf")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())

		graphData := `
			<gexf version="1.2" xmlns="http://www.gexf.net/1.2draft">
				<graph>
					<nodes>
						<node id="1" label="Node 1"/>
						<node id="2" label="Node 2"/>
					</nodes>
					<edges>
						<edge id="1" source="1" target="2"/>
					</edges>
				</graph>
			</gexf>
		`
		_, err = tmpfile.WriteString(graphData)
		if err != nil {
			t.Fatal(err)
		}
		tmpfile.Close()

		file, err := os.Open(tmpfile.Name())
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", filepath.Base(tmpfile.Name()))
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Fatal(err)
		}
		err = writer.Close()
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/upload", body)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = req

		mockService := mocks.NewUploadService(t)
		mockService.On("Invoke", mock.Anything).Return(domain.UploadDTO{
			Status: true,
		}, nil).Maybe()

		action := NewUploadAction(mockService)

		action.Invoke(ctx)

		expectedResponse := responder.UploadProjection{
			Status: true,
		}

		var response responder.UploadProjection
		b, _ := io.ReadAll(w.Body)
		_ = json.Unmarshal(b, &response)

		assert.Equal(t, 200, w.Code)

		difference, err := diff.Diff(expectedResponse, response)
		assert.NoError(t, err)
		assert.Len(t, difference, 0)
	})

	t.Run("should return status 400 when file not ends in .gexf", func(t *testing.T) {
		tmpfile, err := os.CreateTemp("", "*graph.invalid")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())

		graphData := `test`
		_, err = tmpfile.WriteString(graphData)
		if err != nil {
			t.Fatal(err)
		}
		tmpfile.Close()

		file, err := os.Open(tmpfile.Name())
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", filepath.Base(tmpfile.Name()))
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Fatal(err)
		}
		err = writer.Close()
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/upload", body)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = req

		mockService := mocks.NewUploadService(t)
		mockService.On("Invoke", mock.Anything).Return(domain.UploadDTO{
			Status: true,
		}, nil).Maybe()
		action := NewUploadAction(mockService)

		action.Invoke(ctx)

		expectedResponse := responder.UploadErrorProjection{
			Error: "File extension not supported",
		}

		var response responder.UploadErrorProjection
		b, _ := io.ReadAll(w.Body)
		_ = json.Unmarshal(b, &response)

		assert.Equal(t, 400, w.Code)

		difference, err := diff.Diff(expectedResponse, response)
		assert.NoError(t, err)
		assert.Len(t, difference, 0)
	})

	t.Run("should return status 400 when file is too large", func(t *testing.T) {
		tmpfile, err := os.CreateTemp("", "*graph.gexf")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())

		data := make([]byte, 10*1024*1024) // 10MB
		_, err = tmpfile.Write(data)
		if err != nil {
			panic(err)
		}
		tmpfile.Close()

		file, err := os.Open(tmpfile.Name())
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", filepath.Base(tmpfile.Name()))
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Fatal(err)
		}
		err = writer.Close()
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/upload", body)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = req

		mockService := mocks.NewUploadService(t)
		mockService.On("Invoke", mock.Anything).Return(domain.UploadDTO{
			Status: true,
		}, nil).Maybe()
		action := NewUploadAction(mockService)

		action.Invoke(ctx)

		expectedResponse := responder.UploadErrorProjection{
			Error: "http: request body too large",
		}

		var response responder.UploadErrorProjection
		b, _ := io.ReadAll(w.Body)
		_ = json.Unmarshal(b, &response)

		assert.Equal(t, 400, w.Code)

		difference, err := diff.Diff(expectedResponse, response)
		assert.NoError(t, err)
		assert.Len(t, difference, 0)
	})

	t.Run("should return status 400 when fieldname is not file", func(t *testing.T) {
		tmpfile, err := os.CreateTemp("", "*graph.gexf")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())

		graphData := `test`
		_, err = tmpfile.WriteString(graphData)
		if err != nil {
			t.Fatal(err)
		}
		tmpfile.Close()

		file, err := os.Open(tmpfile.Name())
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("test", filepath.Base(tmpfile.Name()))
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Fatal(err)
		}
		err = writer.Close()
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/upload", body)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = req

		mockService := mocks.NewUploadService(t)
		mockService.On("Invoke").Return(domain.UploadDTO{
			Status: true,
		}, nil).Maybe()
		action := NewUploadAction(mockService)

		action.Invoke(ctx)

		expectedResponse := responder.UploadErrorProjection{
			Error: "Key: 'UploadRequestModel.File' Error:Field validation for 'File' failed on the 'required' tag",
		}

		var response responder.UploadErrorProjection
		b, _ := io.ReadAll(w.Body)
		_ = json.Unmarshal(b, &response)

		assert.Equal(t, 400, w.Code)

		difference, err := diff.Diff(expectedResponse, response)
		assert.NoError(t, err)
		assert.Len(t, difference, 0)
	})
}
