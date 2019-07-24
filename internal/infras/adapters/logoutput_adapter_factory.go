package adapters

// LogOutputAdapterFactory LogOutputAdapter factory pattern
type LogOutputAdapterFactory struct{}

// NewLogOutputAdapterFactory Create new LogOutputAdapterFactory instance
func NewLogOutputAdapterFactory() *LogOutputAdapterFactory {
	return &LogOutputAdapterFactory{}
}

// Create factory pattern at LogOutputAdapter
func (loaf LogOutputAdapterFactory) Create(pattern string, connection string) LogOutputAdapter {
	switch pattern {
	case "file":
		return NewFileLogOutputAdapter(connection)
	default:
		return NewFileLogOutputAdapter(connection)
	}
}

func GetConnectionString() {
}
