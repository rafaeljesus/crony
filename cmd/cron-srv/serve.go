package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/EmpregoLigado/cron-srv/pkg/checker"
	"github.com/EmpregoLigado/cron-srv/pkg/config"
	"github.com/EmpregoLigado/cron-srv/pkg/datastore"
	"github.com/EmpregoLigado/cron-srv/pkg/handlers"
	"github.com/EmpregoLigado/cron-srv/pkg/repos"
	"github.com/EmpregoLigado/cron-srv/pkg/scheduler"
	log "github.com/Sirupsen/logrus"
	"github.com/nbari/violetear"
	"github.com/spf13/cobra"
)

func Serve(cmd *cobra.Command, args []string) {
	log.WithField("version", version).Info("Cron Service starting...")

	env, err := config.LoadEnv()
	failOnError(err, "Failed to load config!")

	level, err := log.ParseLevel(strings.ToLower(env.LogLevel))
	failOnError(err, "Failed to get log level!")
	log.SetLevel(level)

	ds, err := datastore.New(env.DatastoreURL)
	failOnError(err, "Failed to init dababase connection!")
	defer ds.Close()

	checkers := map[string]checker.Checker{
		"api":      checker.NewApi(),
		"postgres": checker.NewPostgres(env.DatastoreURL),
	}
	healthzHandler := handlers.NewHealthzHandler(checkers)

	eventRepo := repos.NewEvent(ds)

	sc := scheduler.New()
	go sc.ScheduleAll(eventRepo)

	eventsHandler := handlers.NewEventsHandler(eventRepo, sc)

	r := violetear.New()
	r.LogRequests = true
	r.RequestID = "X-Request-ID"
	r.AddRegex(":id", `^\d+$`)

	r.HandleFunc("/v1/healthz", healthzHandler.HealthzIndex, "GET")

	r.HandleFunc("/v1/events", eventsHandler.EventsIndex, "GET")
	r.HandleFunc("/v1/events", eventsHandler.EventsCreate, "POST")
	r.HandleFunc("/v1/events/:id", eventsHandler.EventsShow, "GET")
	r.HandleFunc("/v1/events/:id", eventsHandler.EventsUpdate, "PUT")
	r.HandleFunc("/v1/events/:id", eventsHandler.EventsDelete, "DELETE")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", env.Port), r))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.WithError(err).Panic(msg)
	}
}
