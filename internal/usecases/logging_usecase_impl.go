package usecases

import (
	"net/http"

	"github.com/rocinante-ein/badgerhole/internal/entities"
	"github.com/rocinante-ein/badgerhole/internal/infras/adapters"
	"github.com/spf13/viper"
)

// LoggingUsecaseImpl badger hole logging usecase implements
type LoggingUsecaseImpl struct{}

// NewLoggingUsecase create new logging usecase instance
func NewLoggingUsecase() LoggingUsecase {
	return &LoggingUsecaseImpl{}
}

// Logging output log to other interface
func (lui LoggingUsecaseImpl) Logging(r *http.Request) {

	// create WWWLog format instance
	var wwwlog = entities.NewWWWLog(r)

	// create LogOutputAdapterFactory instance
	var logOutputAdapterFactory = adapters.NewLogOutputAdapterFactory()

	// create FileLogOutput instance
	var fileoutput = logOutputAdapterFactory.Create(
		viper.GetString("LogOutputAdapterType"),
		viper.GetString("LogOutputAdapterConnection"),
	)

	// output wwwlog to file
	fileoutput.Output(wwwlog)
}
