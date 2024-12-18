package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/KrutyashBB/CalcProject/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	response := Response{}

	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Error = fmt.Errorf("Internal server error").Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := calculation.Calc(request.Expression)

	if err != nil {
		response.Error = "Expression is not valid"
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		response.Result = fmt.Sprintf("%f", result)
	}
	json.NewEncoder(w).Encode(response)
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
