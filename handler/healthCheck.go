package handler

type checkResponse struct {
	status	string
}

func HealthCheck() (checkResponse){
	response := checkResponse{}
	return response
}