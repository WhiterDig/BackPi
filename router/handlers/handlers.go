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

func SaveCredential(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flag.Parse()
		incomingBody := json.NewDecoder(r.Body)
		creds := model.Credentials{}
		incomingBody.Decode(&creds)

		resp, err := ctx.Value(global.CredentialService).(service.Credential).SaveCredential(creds)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, *resp)
	}
}
