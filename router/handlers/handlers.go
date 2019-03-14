package handlers

import (
	"BackPi/global"
	"BackPi/model"
	"BackPi/service"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

func DoThings(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flag.Parse()
		fmt.Println("Point hit")
		incomingBody := json.NewDecoder(r.Body)
		creds := model.Credentials{}
		incomingBody.Decode(&creds)

		fmt.Println("*****user******", creds.SSID)
		resp, err := ctx.Value(global.CredentialService).(service.Credential).WriteToFile(creds.SSID, creds.Password, creds.Email)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, *resp)
	}
}
