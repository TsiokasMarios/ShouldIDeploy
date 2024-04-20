package server

import (
	"ShouldIDeploy/cmd/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)
	mux.HandleFunc("/shouldIDeploy", s.ShouldIDeploy)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

// A route that accepts a GET request with the paremeters of TimeZone and Date
func (s *Server) ShouldIDeploy(w http.ResponseWriter, r *http.Request) {
	//Get the timezone and date from the request
	timezone := r.URL.Query().Get("timezone")
	//date := r.URL.Query().Get("date")
	language := r.URL.Query().Get("language")

	utils.SetUpTimezone(timezone)

	//Get the reasons based on the timezone and date
	shouldI := utils.ShouldIDeploy()
	reason := utils.GetRandom(utils.GetReason())

	//Send the response that has the reason and shouldI
	resp := make(map[string]string)
	resp["shouldI"] = strconv.FormatBool(shouldI)
	resp["reason"] = utils.Localize(reason, language)
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
