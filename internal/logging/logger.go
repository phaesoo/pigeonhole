package logging

import (
	"io/ioutil"
	"os"

	"github.com/phaesoo/pigeonhole/configs"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmlogrus"
)

// Logger is a logger that can
type Logger logrus.FieldLogger

type LogErrField struct {
	Err string `json:"error"`
}

func NewErrorField(e error) LogErrField {
	return LogErrField{
		Err: e.Error(),
	}
}

func AddHookToLogger(l Logger, h logrus.Hook) {
	l.(*logrus.Entry).Logger.AddHook(h)
}

// New creates a new logger with the give name.
func New(name string, config configs.LogConfig) Logger {
	logger := logrus.New()

	if config.FormatJSON {
		logger.Formatter = newJSONFormatter()
	} else {
		logger.Formatter = &logrus.TextFormatter{
			ForceColors: true,
		}
	}

	logger.SetLevel(config.Level)
	// Logrus default output channel is stderr
	logger.SetOutput(os.Stdout)

	return logger.WithFields(logrus.Fields{
		"name": name,
		"env":  config.Environment,
	})
}

// Discard is a logger that has all of the same functionality as an ordinary logger, but discards
// all logging output. Useful for testing.
var Discard Logger

func init() {
	logrus.AddHook(&apmlogrus.Hook{})
	logger := logrus.New()
	logger.Out = ioutil.Discard
	Discard = logger
}
