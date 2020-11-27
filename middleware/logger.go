package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if ip != "" {
			ip = r.Header.Get("X-FORWARDED-FOR")
		}
		log.WithFields(log.Fields{
			"ip":  ip,
			"uri": r.RequestURI,
		}).Info()

		next.ServeHTTP(w, r)
	})
}

func Log(message string, level string) {

	loglevel := LogLevels(level)

	switch loglevel {
	case log.InfoLevel:
		log.WithFields(log.Fields{
			"message": message,
		}).Infoln()
	case log.ErrorLevel:
		log.WithFields(log.Fields{
			"message": message,
		}).Errorln()
	default:
		log.WithFields(log.Fields{
			"message": message,
		}).Debugln()
	}

}

func LogError(message error) {

	log.WithFields(log.Fields{
		"error": message,
	}).Errorln()

}

func LogLevels(level string) log.Level {
	switch level {
	case "info":
		return log.InfoLevel

	case "error":
		return log.ErrorLevel
	}
	return log.DebugLevel
}
