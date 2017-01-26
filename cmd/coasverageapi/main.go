package main

import (
	"github.com/gorilla/handlers"
	"github.com/udistrital/coasverage-api/internal/coasverageapi"
	"github.com/urfave/cli"
	"net/http"
	"os"
)

var app *cli.App
var listen string
var port string
var cors_allowed_origins cli.StringSlice = []string{"*"}

func init() {
	app = cli.NewApp()
	app.Commands = []cli.Command{
		cli.Command{
			Name: "run",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "listen",
					EnvVar:      "LISTEN",
					Destination: &listen,
					Value:       "127.0.0.1",
				},
				cli.StringFlag{
					Name:        "port",
					EnvVar:      "PORT",
					Destination: &port,
					Value:       "3000",
				},
				cli.StringSliceFlag{
					Name:   "cors_allowed_origins",
					EnvVar: "CORS_ALLOWED_ORIGINS",
					Value:  &cors_allowed_origins,
				},
			},
			Action: action,
		},
	}
}

func action(ctx *cli.Context) (err error) {
	router := coasverageapi.NewRouter()
	handler := handlers.CORS(handlers.AllowedOrigins(cors_allowed_origins))(router)
	return http.ListenAndServe(listen+":"+port, handler)
}

func main() {
	app.Run(os.Args)
}
