package adapters

import (
	"fmt"
	"os"
	"time"

	"github.com/rocinante-ein/badgerhole/internal/entities"
)

// FileLogOutputAdapter File output interface
type FileLogOutputAdapter struct {
	LogFormat string
}

// NewFileLogOutputAdapter Create new UniqueFileLogOutput instance
func NewFileLogOutputAdapter(logFormat string) *FileLogOutputAdapter {
	return &FileLogOutputAdapter{
		LogFormat: logFormat,
	}
}

// Output output file log
func (floa FileLogOutputAdapter) Output(wwwlog entities.WWWLog) {

	// open wirte only file
	logFile, err := os.Create(
		fmt.Sprintf(
			floa.LogFormat,
			time.Now().Format("20060102150405"),
			wwwlog.GetRequestID(),
		),
	)

	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	_, err = logFile.Write(wwwlog.ToJSON())
	if err != nil {
		panic(err)
	}
}
