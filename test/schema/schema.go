package schema

var (

	// Response schema
	Response = `{
		"type": "object",
	    "properties": {
			"message": {"type": "string"},
			"payload": {"type": "object"},
			"status":  {"type": "string"}
	    },
	    "required": ["message", "status", "payload"]
	}`
)
