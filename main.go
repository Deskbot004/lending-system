package main

import (
	sql "lending-system/sql"
	"lending-system/web"
	"log"
)

func main() {
	// Setting up the DB
	client, err := sql.SetupDB()
	if err != nil {
		log.Fatalf("Failed to setup the db: %v", err)
	}
	defer client.Close()

	/* Setting up the timers
	update := os.Getenv("UPDATETIMER")
	reset := os.Getenv("RESETTIMER")

	// starting schedules
	if update == "" {
		log.Println("No UPDATETIMER set")
	} else {
		s := gocron.NewScheduler(time.UTC)
		_, err = s.Every(update).Do(func() {
			err = sql.UpdateAllProjects(client)
			if err != nil {
				log.Printf("Failed updating the projects: %v", err)
			}
		})
		if err != nil {
			log.Printf("failed setting up sheduler for updates with UPDATETIMER: %s", update)
		}
		s.StartAsync()
	}

	if reset == "" {
		log.Println("No RESETTIMER set")
	} else {
		t := gocron.NewScheduler(time.UTC)
		_, err = t.Every(reset).Do(func() {
			client.Close()
			os.Remove("db/game_db.db")
			client, err = sql.SetupDB()
			if err != nil {
				log.Fatalf("Can't setup the db: %v", err)
			}
		})
		if err != nil {
			log.Printf("failed setting up sheduler for reset with RESETTIMER: %s", reset)
		}
		t.StartAsync()
	}
	*/
	web.StartServer(client)
	// TODO Sort by ausgeliehen
	// TODO Jonaskopf
	// TODO search bar for game search
}
