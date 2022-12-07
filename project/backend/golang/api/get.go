package api

import (
	"encoding/json"
	"golang/commons"
	"golang/database"
	"log"
	"net/http"
	"strconv"
)

func get(writer http.ResponseWriter, request *http.Request) {

	apiReturn := func(httpStatusCode int) {
		commons.Api.StatusCodeReturn(commons.Api{}, "get", &writer, httpStatusCode)
	}

	if request.Method != "GET" {
		apiReturn(http.StatusMethodNotAllowed)
		return
	}

	log.Println("endpoint \"/get\" contact!")

	data := commons.Api.NewTable(commons.Api{})
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		log.Println("error on json decode:", err)
		apiReturn(http.StatusUnprocessableEntity)
		return
	}

	var where string
	var and string
	var total float32

	table := "finance"

	if optional := data.Optional; optional == "revenue" {
		where = "WHERE type = 'REVENUE'"
		and = "AND status = 'PAID'"
	}

	// SELECT f.value FROM openfinance.finance f WHERE f.type = 'REVENUE' AND f.status = 'PAID'
	response, err := database.List(table, where, and)
	if err != nil {
		log.Println("error on database list!")
		apiReturn(http.StatusInternalServerError)
		return
	}

	// float -> string: strconv.FormatFloat(float_variable, 'f', 2, 32)
	for _, row := range response {
		temporaryValue, err := strconv.ParseFloat(row["value"], 32)
		if err != nil {
			log.Println("error on converting string to float!")
			apiReturn(http.StatusInternalServerError)
			return
		}
		total += float32(temporaryValue)
	}

	json.NewEncoder(writer).Encode(total)

}
