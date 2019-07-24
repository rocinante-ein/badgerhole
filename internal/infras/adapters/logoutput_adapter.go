package adapters

import "github.com/rocinante-ein/badgerhole/internal/entities"

// LogOutputAdapter log output interface
type LogOutputAdapter interface {
	Output(entities.WWWLog)
}
