package guacamole

type ConnectResponse struct {
	AuthToken            string   `json:"authToken"`
	Username             string   `json:"username"`
	DataSource           string   `json:"dataSource"`
	AvailableDataSources []string `json:"availableDataSource"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	// TranslatableMessage []map[string]string `json:"translatableMessage"`
	StatusCode string `json:"statusCode"`
	Expected   string `json:"expected"`
	Type       string `json:"type"`
}
