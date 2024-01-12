package app

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/SCR-NETWORK/SCR_Network/infrastructure/config"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/db/database"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/db/database/ldb"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/logger"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/os/execenv"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/os/limits"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/os/signal"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/os/winservice"
	"github.com/SCR-NETWORK/SCR_Network/util/panics"
	"github.com/SCR-NETWORK/SCR_Network/util/profiling"
	"github.com/SCR-NETWORK/SCR_Network/version"
)

const (
	leveldbCacheSizeMiB = 256
	defaultDataDirname  = "datadir2"
)

var desiredLimits = &limits.DesiredLimits{
	FileLimitWant: 2048,
	FileLimitMin:  1024,
}

var serviceDescription = &winservice.ServiceDescription{
<<<<<<< HEAD
	Name:        "SCR-NETWORKsvc",
	DisplayName: "SCR-NETWORK Service",
=======
	Name:        "SCR_Networksvc",
	DisplayName: "SCR_Network Service",
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
	Description: "Downloads and stays synchronized with the SCR blockDAG and " +
		"provides DAG services to applications.",
}

<<<<<<< HEAD
type SCR-NETWORKApp struct {
	cfg *config.Config
}

// StartApp starts the SCR-NETWORK app, and blocks until it finishes running
=======
type SCR_NetworkApp struct {
	cfg *config.Config
}

// StartApp starts the SCR_Network app, and blocks until it finishes running
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
func StartApp() error {
	execenv.Initialize(desiredLimits)

	// Load configuration and parse command line. This function also
	// initializes logging and configures it accordingly.
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	defer logger.BackendLog.Close()
	defer panics.HandlePanic(log, "MAIN", nil)

<<<<<<< HEAD
	app := &SCR-NETWORKApp{cfg: cfg}
=======
	app := &SCR_NetworkApp{cfg: cfg}
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4

	// Call serviceMain on Windows to handle running as a service. When
	// the return isService flag is true, exit now since we ran as a
	// service. Otherwise, just fall through to normal operation.
	if runtime.GOOS == "windows" {
		isService, err := winservice.WinServiceMain(app.main, serviceDescription, cfg)
		if err != nil {
			return err
		}
		if isService {
			return nil
		}
	}

	return app.main(nil)
}

<<<<<<< HEAD
func (app *SCR-NETWORKApp) main(startedChan chan<- struct{}) error {
=======
func (app *SCR_NetworkApp) main(startedChan chan<- struct{}) error {
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
	// Get a channel that will be closed when a shutdown signal has been
	// triggered either from an OS signal such as SIGINT (Ctrl+C) or from
	// another subsystem such as the RPC server.
	interrupt := signal.InterruptListener()
	defer log.Info("Shutdown complete")

	// Show version at startup.
	log.Infof("Version %s", version.Version())

	// Enable http profiling server if requested.
	if app.cfg.Profile != "" {
		profiling.Start(app.cfg.Profile, log)
	}
	profiling.TrackHeap(app.cfg.AppDir, log)

	// Return now if an interrupt signal was triggered.
	if signal.InterruptRequested(interrupt) {
		return nil
	}

	if app.cfg.ResetDatabase {
		err := removeDatabase(app.cfg)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	// Open the database
	databaseContext, err := openDB(app.cfg)
	if err != nil {
		log.Errorf("Loading database failed: %+v", err)
		return err
	}

	defer func() {
		log.Infof("Gracefully shutting down the database...")
		err := databaseContext.Close()
		if err != nil {
			log.Errorf("Failed to close the database: %s", err)
		}
	}()

	// Return now if an interrupt signal was triggered.
	if signal.InterruptRequested(interrupt) {
		return nil
	}

	// Create componentManager and start it.
	componentManager, err := NewComponentManager(app.cfg, databaseContext, interrupt)
	if err != nil {
<<<<<<< HEAD
		log.Errorf("Unable to start SCR-NETWORK: %+v", err)
=======
		log.Errorf("Unable to start SCR_Network: %+v", err)
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
		return err
	}

	defer func() {
<<<<<<< HEAD
		log.Infof("Gracefully shutting down SCR-NETWORK...")
=======
		log.Infof("Gracefully shutting down SCR_Network...")
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4

		shutdownDone := make(chan struct{})
		go func() {
			componentManager.Stop()
			shutdownDone <- struct{}{}
		}()

		const shutdownTimeout = 2 * time.Minute

		select {
		case <-shutdownDone:
		case <-time.After(shutdownTimeout):
			log.Criticalf("Graceful shutdown timed out %s. Terminating...", shutdownTimeout)
		}
<<<<<<< HEAD
		log.Infof("SCR-NETWORK shutdown complete")
=======
		log.Infof("SCR_Network shutdown complete")
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
	}()

	componentManager.Start()

	if startedChan != nil {
		startedChan <- struct{}{}
	}

	// Wait until the interrupt signal is received from an OS signal or
	// shutdown is requested through one of the subsystems such as the RPC
	// server.
	<-interrupt
	return nil
}

// dbPath returns the path to the block database given a database type.
func databasePath(cfg *config.Config) string {
	return filepath.Join(cfg.AppDir, defaultDataDirname)
}

func removeDatabase(cfg *config.Config) error {
	dbPath := databasePath(cfg)
	return os.RemoveAll(dbPath)
}

func openDB(cfg *config.Config) (database.Database, error) {
	dbPath := databasePath(cfg)

	err := checkDatabaseVersion(dbPath)
	if err != nil {
		return nil, err
	}

	log.Infof("Loading database from '%s'", dbPath)
	db, err := ldb.NewLevelDB(dbPath, leveldbCacheSizeMiB)
	if err != nil {
		return nil, err
	}

	return db, nil
}
