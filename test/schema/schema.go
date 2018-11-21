package schema

var (

	// Response schema
	Response = `{
		"type": "object",
	    "properties": {
			"message": {"type": "string"},
			"payload": {"type": ["object", "array"]},
			"status":  {"type": "string"}
	    },
	    "required": ["message", "status", "payload"]
	}`
)
