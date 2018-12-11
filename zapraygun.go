package zapraygun

import (
	"github.com/MindscapeHQ/raygun4go"
	"go.uber.org/zap/zapcore"
	"log"
)

type RaygunHook struct {
	// Messages with a log level not contained in this array
	// will not be dispatched. If nil, all messages will be dispatched.
	AcceptedLevels []zapcore.Level
	RaygunClient   *raygun4go.Client
}

func NewRaygunHook(appName string, apiKey string, level ...zapcore.Level) *RaygunHook {
	raygun, err := raygun4go.New("appName", "apiKey")
	if err != nil {
		log.Println("Unable to create Raygun client:", err.Error())
	}

	return &RaygunHook{
		RaygunClient:   raygun,
		AcceptedLevels: level,
	}
}

func (rh *RaygunHook) GetHook() func(zapcore.Entry) error {

	return func(entry zapcore.Entry) error {

		if !rh.isAcceptedLevel(entry.Level) {
			return nil
		}

		return rh.RaygunClient.CreateError(entry.Message)
	}
}

func (rh *RaygunHook) isAcceptedLevel(level zapcore.Level) bool {

	for _, lv := range rh.AcceptedLevels {
		if lv == level {
			return true
		}
	}
	return false
}
