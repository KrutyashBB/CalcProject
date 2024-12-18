package application_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/KrutyashBB/CalcProject/internal/application"
)

func TestCalcHandler_Success(t *testing.T) {
	reqBody := application.Request{
        Expression: "2 + 2",
    }
    body, _ := json.Marshal(reqBody)

    req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(body))
    w := httptest.NewRecorder()

    application.CalcHandler(w, req)

    res := w.Result()
    if res.StatusCode != http.StatusOK {
        t.Errorf("expected status %d, got %d", http.StatusOK, res.StatusCode)
    }

    var response application.Response
    err := json.NewDecoder(res.Body).Decode(&response)
    if err != nil {
        t.Fatalf("failed to decode response: %v", err)
    }

    if response.Result != "4.000000" {
        t.Errorf("expected result %s, got %s", "4.000000", response.Result)
    }
}

func TestCalcHandler_InvalidJSON(t *testing.T) {
    req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader([]byte("invalid json")))
    w := httptest.NewRecorder()

    application.CalcHandler(w, req)

    res := w.Result()
    if res.StatusCode != http.StatusInternalServerError {
        t.Errorf("expected status %d, got %d", http.StatusInternalServerError, res.StatusCode)
    }

    var response application.Response
    err := json.NewDecoder(res.Body).Decode(&response)
    if err != nil {
        t.Fatalf("failed to decode response: %v", err)
    }

    if response.Error != "Internal server error" {
        t.Errorf("expected error message %s, got %s", "Internal server error", response.Error)
    }
}

func TestCalcHandler_InvalidExpression(t *testing.T) {
    reqBody := application.Request{
        Expression: "2 / 0", 
    }
    body, _ := json.Marshal(reqBody)

    req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(body))
    w := httptest.NewRecorder()

    application.CalcHandler(w, req)

    res := w.Result()
    if res.StatusCode != http.StatusUnprocessableEntity {
        t.Errorf("expected status %d, got %d", http.StatusUnprocessableEntity, res.StatusCode)
    }

    var response application.Response
    err := json.NewDecoder(res.Body).Decode(&response)
    if err != nil {
        t.Fatalf("failed to decode response: %v", err)
    }

    if response.Error != "Expression is not valid" {
        t.Errorf("expected error message %s, got %s", "Expression is not valid", response.Error)
    }
}

