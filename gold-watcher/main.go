package main

import (
	"database/sql"
	"gold-watcher/repository"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	_ "github.com/glebarez/go-sqlite"
	"github.com/joho/godotenv"
)

type Config struct {
	App                            fyne.App
	InfoLog                        *log.Logger
	ErrorLog                       *log.Logger
	DB                             repository.Repository
	MainWindow                     fyne.Window
	PriceContainer                 *fyne.Container
	Toolbar                        *widget.Toolbar
	PriceChartContainer            *fyne.Container
	Holdings                       [][]any
	HoldingsTable                  *widget.Table
	HTTPClient                     *http.Client
	AddHoldingsPurchaseAmountEntry *widget.Entry
	AddHoldingsPurchaseDateEntry   *widget.Entry
	AddHoldingsPurchasePriceEntry  *widget.Entry
}

func main() {
	var myApp Config

	// create a fyne application
	fyneApp := app.NewWithID("piatoss.goldwatcher")
	myApp.App = fyneApp

	// set http client
	myApp.HTTPClient = &http.Client{}

	// create loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		myApp.ErrorLog.Println("error loading environment variables from .env file")
	}

	// open a connection to the database
	sqlDB, err := myApp.ConnectSQL()
	if err != nil {
		log.Panic(err)
	}

	// create a database repository
	myApp.setupDB(sqlDB)

	// create and size the fyne window
	myApp.MainWindow = fyneApp.NewWindow("Gold Watcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	// get application UIs
	myApp.makeUI()

	// show and run the application
	myApp.MainWindow.ShowAndRun()
}

// connect to SQLite DB
func (app *Config) ConnectSQL() (*sql.DB, error) {
	path := ""

	// check if "DB_PATH" environment variable exists
	// else set fyne app storage path as DB source path
	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		app.InfoLog.Println("db in:", path)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	app.InfoLog.Println("connected to SQLite")

	return db, nil
}

// setup SQLite DB
func (app *Config) setupDB(sqlDB *sql.DB) {
	app.DB = repository.NewSQLiteRepo(sqlDB)

	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Panic()
	}
}
