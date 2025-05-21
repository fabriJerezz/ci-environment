package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name string
		sentArgument string
		expectedBody string
	} {
		{
			"real province",
			"cordoba",
			"La capital es Córdoba",
		},
		{
			"another word",
			"scrum",
			"No se encontró la provincia scrum",
		},
		{
			"empty string",
			"",
			"Por favor ingrese una provincia",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/obtener-capital?provincia=%s", tt.sentArgument)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()
			
			Handler(w, req)
			
			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status 200, got %d", res.StatusCode)
			}
			
			body := w.Body.String()
			
			
			if body != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, body)
			}
 		})
	}
}
	
	




