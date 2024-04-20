package main

import (
	"ShouldIDeploy/internal/server"
	"fmt"
)

func main() {

	newServer := server.NewServer()
	err := newServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start newServer: %s", err))
	}

	//Get date based on timezone
	//loc, _ := time.LoadLocation("America/New_York")

	//Convert 2024-12-24 16:00:00 to a time object
	//t, _ := time.Parse("2006-01-02 15:04:05", "2024-12-24 16:00:00")

	//println(t.String())
	//Get date based on loc
	//time.Local = loc
	//println("local time: ", time.Now().String())
	//Setup timezone at japan
	//println(utils.ShouldIDeploy())
	//msg := utils.GetRandom(utils.GetReason())
	//println("util date: ", utils.GetDate().String())
	//println("msg: ", utils.Localize(msg, "en"))

	//println(utils.Localize("i_don't_see_why_not", "el"))

}
