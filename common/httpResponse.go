package common

import (
	"net/http"
	"encoding/json"
)

//Sends a 200 OK response with the requested content
func Return200(w http.ResponseWriter, resp []byte){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(resp)
}

//Sends a 201 new record created response
func Return201(w http.ResponseWriter, resp []byte){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(resp)
}

//Sends a 204 no content found if the result set is empty
func EmptyResult (w http.ResponseWriter){
	var resp ResposneStruct
	resp.Message = "Result set is empty"
	resp.Status = 204
	response, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	w.Write(response)
}

//This is a specific response sent if the marshalling or unmarshal fails
func Return210ErrorResponse(w http.ResponseWriter){
	var resp ResposneStruct
	resp.Status = 210
	resp.Message = "Error in marshalling the response from DB, contact admin"
	w.WriteHeader(210)
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(resp)
	w.Write(response)
}
//Standard 400 bad request
func Return400(w http.ResponseWriter){
	var resp ResposneStruct
	resp.Status = 400
	resp.Message = "Bad Request"
	response, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	w.Write(response)
}

//A 409 represents conflict used for duplicate entry in DB, could be used for others as well
func Return409Error(w http.ResponseWriter, error ResposneStruct){
	w.WriteHeader(409)
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(error)
	w.Write(response)
}

func Return417ErrorResponse(w http.ResponseWriter){
	var resp ResposneStruct
	resp.Status = 417
	resp.Message = "Error in input validation, please check the input fields"
	w.WriteHeader(417)
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(resp)
	w.Write(response)
}