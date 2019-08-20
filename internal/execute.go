package internal

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kapustkin/go_ms_template/controllers/calendar"
	"github.com/kapustkin/go_ms_template/internal/config"
	"github.com/kapustkin/go_ms_template/internal/logger"
	"github.com/sirupsen/logrus"
)

// Run основной обработчик
func Run(args []string) error {
	c := config.InitConfig()

	r := chi.NewRouter()
	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Logging
	switch c.Logging {
	case 1:
		r.Use(middleware.Logger)
	case 2:
		log := logrus.New()
		log.Formatter = &logrus.JSONFormatter{
			DisableTimestamp: true,
		}
		r.Use(logger.NewChiLogger(log))
	default:
		log.Printf("Warning! Starting without logging... \n")
	}

	// Routes
	r.HandleFunc("/calendar/{user}", calendar.GetUserCalandar)

	return http.ListenAndServe(fmt.Sprintf("%s:%v", c.Host, c.Port), r)
}
