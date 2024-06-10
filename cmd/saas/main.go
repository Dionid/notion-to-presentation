package main

import (
	"context"
	"log"
	"os"
	"path"
	"strings"

	"github.com/Dionid/notion-to-presentation/cmd/saas/httph"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/mails"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	config, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	gctx, _ := context.WithCancel(context.Background())

	// # Pocketbase
	app := pocketbase.New()

	// # Migrations
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	curPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun,
		Dir:         path.Join(curPath, "pb_migrations"),
	})

	// # HTTP API
	httph.InitApi(httph.Config{
		PreviewId: config.PreviewId,
	}, app, gctx)

	// # Send verification email on sign-up
	app.OnRecordAfterCreateRequest("users").Add(func(e *core.RecordCreateEvent) error {
		return mails.SendRecordVerification(app, e.Record)
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
