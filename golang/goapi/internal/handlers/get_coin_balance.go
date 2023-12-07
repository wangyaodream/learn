package handlers

import (
    "encoding/json"
    "net/http"

    "gihub.com/wangyaodream/goapi/api"
    "gihub.com/wangyaodream/goapi/internal/tools"
    log "github.com/sirupsen/logrus"
    "github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
     var params = api.CoinBalanceParams{}
     var decoder *schema.Decoder = schema.NewDecoder()
     var err error

     err = decoder.Decode(&params, r.URL.Query())

     if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
     }

     var database *tools.DatabaseInterface
     database, err = tools.NewDatabase()
     if err != nil {
        api.InternalErrorHandler(w)
        return
     }

     var tokenDetail *tools.CoinDetails
     tokenDetail = (*database).GetUserCoins(params.Username)
     if tokenDetail == nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
     }

     var response = api.CoinBalanceResponse{
        Balance: (*tokenDetail).Coins,
        Code: http.StatusOK,
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }
}
