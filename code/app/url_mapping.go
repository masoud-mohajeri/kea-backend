package app

import "github.com/masoud-mohajeri/kea-backend/routes"

func url_mapping() {
	routes.NewPing("ping", app)
	routes.NewAuth("api/v1/auth", app)
}
