/*
Copyright 2019 Sensible

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sdt"
	app.Usage = "Schema Translator"
	app.Action = func(c *cli.Context) error {

		connectString := c.Args().First()
		if connectString == "" {
			panic("Pass a connection string url. Like this: postgres://username:password@localhost/db?sslmode=disable")
		}

		db, err := sql.Open("postgres", connectString)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		sqlStatement := `SELECT table_schema, table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema NOT IN ('pg_catalog', 'information_schema');`
		var (
			table_schema string
			table_name   string
		)

		rows, err := db.Query(sqlStatement)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&table_schema, &table_name)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(table_schema, table_name)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
