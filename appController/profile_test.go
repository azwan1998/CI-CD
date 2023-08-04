package appController

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"gofrendi/structureExample/appModel"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestProfileController_GetAll(t *testing.T) {
	// Buat instance dari echo
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/profiles", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Buat instance dari ProfileController
	pc := NewProfileController(appModel.NewProfileMemModel(), "jwt_secret")

	// Panggil fungsi GetAll
	if assert.NoError(t, pc.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// TODO: Add more assertions to check the response body or other things
	}
}

func TestProfileController_GetById(t *testing.T) {
	// Buat instance dari echo
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/profiles/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Buat instance dari ProfileController
	pc := NewProfileController(appModel.NewProfileMemModel(), "jwt_secret")

	// Panggil fungsi GetById
	if assert.NoError(t, pc.GetById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// TODO: Add more assertions to check the response body or other things
	}
}

func TestProfileController_Add(t *testing.T) {
	// Buat instance dari echo
	e := echo.New()
	jsonData := `{"id_user": 1, "alamat": "Jl. ABC No. 123", "institusi": "Universitas XYZ", "isApprove": false}`
	req := httptest.NewRequest(http.MethodPost, "/profiles", strings.NewReader(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Buat instance dari ProfileController
	pc := NewProfileController(appModel.NewProfileMemModel(), "jwt_secret")

	// Panggil fungsi Add
	if assert.NoError(t, pc.Add(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// TODO: Add more assertions to check the response body or other things
	}
}

func TestProfileController_Edit(t *testing.T) {
	// Buat instance dari echo
	e := echo.New()

	// Data JSON untuk request
	jsonData := `{"alamat": "Jl. XYZ No. 456", "institusi": "Universitas ABC"}`
	req := httptest.NewRequest(http.MethodPut, "/profiles/1", bytes.NewBufferString(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Buat instance dari ProfileController
	pc := NewProfileController(appModel.NewProfileMemModel(), "jwt_secret")

	// Panggil fungsi Edit
	if assert.NoError(t, pc.Edit(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// TODO: Add more assertions to check the response body or other things
	}
}

func createMockToken() *jwt.Token {
	// Buat token palsu sesuai dengan struktur yang diharapkan
	// Misalnya, sesuaikan data user ID dan role berdasarkan kebutuhan unit test
	return &jwt.Token{
		Claims: jwt.MapClaims{
			"id_user": 1,       // Sesuaikan dengan user ID yang diharapkan
			"role":    "admin", // Sesuaikan dengan role yang diharapkan
		},
	}
}

func createMockContext(token *jwt.Token) echo.Context {
	// Buat MockContext dengan token palsu
	req := httptest.NewRequest(http.MethodPost, "/profile", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	c.Set("user", token) // Set token ke dalam konteks
	return c
}

func TestProfileController_uploadFile(t *testing.T) {
	// Buat instance dari echo
	e := echo.New()

	// Buat file dummy untuk diunggah
	file, err := os.Create("test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	file.WriteString("Test data")
	file.Close()
	defer os.Remove("test.jpg")

	// Buat request dengan file yang diunggah
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("foto", "test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	file, err = os.Open("test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	_, err = io.Copy(part, file)
	writer.Close()

	// Lakukan request
	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Buat instance dari ProfileController
	pc := NewProfileController(appModel.NewProfileMemModel(), "jwt_secret")

	// Panggil fungsi uploadFile
	fileName, err := pc.uploadFile(c, "foto")
	if err != nil {
		t.Fatal(err)
	}

	// Pastikan file terunggah dengan benar
	assert.NotEmpty(t, fileName)
	assert.FileExists(t, "C:/laragon/www/GoApp/IslamicNews/storage/file/"+fileName)
}

func TestProfileController_ApproveProfile(t *testing.T) {
	// Buat instance dari echo
	e := echo.New()

	// Data JSON untuk request
	jsonData := `{"isApprove": true}`
	req := httptest.NewRequest(http.MethodPut, "/profiles/1/approve", bytes.NewBufferString(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Buat instance dari ProfileController
	pc := NewProfileController(appModel.NewProfileMemModel(), "jwt_secret")

	// Panggil fungsi ApproveProfile
	if assert.NoError(t, pc.ApproveProfile(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// TODO: Add more assertions to check the response body or other things
	}
}
