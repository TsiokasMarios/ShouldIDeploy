package server

import (
	"ShouldIDeploy/cmd/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/shouldIDeploy", s.ShouldIDeploy)

	return mux
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// A route that accepts a GET request with the paremeters of TimeZone and Date
func (s *Server) ShouldIDeploy(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	//Get the timezone, date and language from the request
	timezone := r.URL.Query().Get("timezone")
	//date := r.URL.Query().Get("date")
	language := r.URL.Query().Get("language")

	utils.SetUpTimezone(timezone)

	//Get the reasons based on the timezone and date
	shouldI := utils.ShouldIDeploy()
	reason := utils.GetRandom(utils.GetReason())

	date, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST"))
	if err != nil {
		log.Fatalf("error handling time parse. Err: %v", err)
	}

	//Send the response that has the reason and shouldI
	resp := make(map[string]string)
	resp["timezone"] = timezone
	resp["date"] = date.String()
	resp["should I Deploy"] = strconv.FormatBool(shouldI)
	resp["message"] = utils.Localize(reason, language)
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
