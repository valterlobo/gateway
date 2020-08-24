package command

type Property map[string]interface{}

func (m Property) Get(key string) interface{} {
	if v, ok := m[key]; ok {
		return v
	}

	return nil
}

// Metadata is sent with every message to provide extra context without unmarshaling the message payload.
type Metadata map[string]string

func (m Metadata) Get(key string) string {
	if v, ok := m[key]; ok {
		return v
	}

	return ""
}

func (m Metadata) Set(key, value string) {
	m[key] = value
}

type Command struct {
	UUID        string
	Namespace   string
	CommandType string
	Property    Property
	Metadata    Metadata
}
