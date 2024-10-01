package main

func initDatabase() {
	db, err := sql.Open(driver, dataSourceName)

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
