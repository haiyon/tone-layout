package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sample/internal/conf"
	"sample/internal/data/ent/migrate"
	"sample/pkg/utils"

	atlas "ariga.io/atlas/sql/migrate"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	// postgres driver
	_ "github.com/lib/pq"
)

var (

	// flagConf is the config flag.
	flagConf string
	//flagDir is the migrate dir
	flagDir string
	//flagName is the migrate script name
	flagName string
)

func init() {
	flag.StringVar(&flagConf, "conf", "config.yaml", "config path, eg: -conf config.yaml")
	flag.StringVar(&flagDir, "dir", "internal/data/migrations", "migrate dir, eg: -dir migrations")
	flag.StringVar(&flagName, "name", "", "migrate script name, eg: -name user")
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagConf),
		),
	)
	defer c.Close()

	if err := c.Load(); utils.IsNotNil(err) {
		fmt.Println(err)
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); utils.IsNotNil(err) {
		panic(err)
	}

	dir, err := atlas.NewLocalDir(flagDir)
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	if len(flagName) == 0 && len(os.Args) == 0 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod cmd/migrate/main.go -name <schema>'")
	} else if len(flagName) == 0 && len(os.Args) > 0 {
		flagName = os.Args[1]
	}

	// Write migration diff.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
		//schema.DisableChecksum(),
	}
	err = migrate.NamedDiff(context.Background(), bc.Data.Database.Source, flagName, opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
