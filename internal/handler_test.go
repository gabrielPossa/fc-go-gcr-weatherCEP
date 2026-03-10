package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func newTestRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/cep/{CEP}/weather", GetWeatherByCEP)
	return r
}

func TestGetWeatherByCEP_InvalidZipcode(t *testing.T) {
	tests := []struct {
		name string
		cep  string
	}{
		{"too short", "1234567"},
		{"too long", "123456789"},
		{"letters", "abcdefgh"},
		{"special chars", "12345_678"},
		{"empty", ""},
	}

	router := newTestRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/cep/"+tt.cep+"/weather", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != http.StatusUnprocessableEntity {
				t.Errorf("expected status 422, got %d", w.Code)
			}
		})
	}
}

func TestGetWeatherByCEP_NotFound(t *testing.T) {
	router := newTestRouter()

	req := httptest.NewRequest("GET", "/cep/99999999/weather", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}
}

func TestGetWeatherByCEP_Success(t *testing.T) {
	router := newTestRouter()

	req := httptest.NewRequest("GET", "/cep/01001000/weather", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d, body: %s", w.Code, w.Body.String())
	}

	var resp response
	err := json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.TempC == 0 && resp.TempF == 0 && resp.TempK == 0 {
		t.Error("all temperatures are zero, expected real values")
	}

	expectedF := resp.TempC*9/5 + 32
	if resp.TempF != expectedF {
		t.Errorf("temp_F = %v, expected %v (from temp_C %v)", resp.TempF, expectedF, resp.TempC)
	}

	expectedK := resp.TempC + 273
	if resp.TempK != expectedK {
		t.Errorf("temp_K = %v, expected %v (from temp_C %v)", resp.TempK, expectedK, resp.TempC)
	}
}
