package log

import (
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

// Logger는 httplog에서 생성된 zerolog.Logger입니다.
var Logger zerolog.Logger

func init() {
	// httplog.NewLogger를 사용하여 로거를 생성합니다.
	Logger = httplog.NewLogger("stylekey-api", httplog.Options{
		JSON:     true,
		Concise:  true,
		LogLevel: "info",
	})
}
