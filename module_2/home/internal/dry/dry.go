package dry

import "fmt"

type Logger struct{}

func (l Logger) Log(level, message string) {
	fmt.Printf("%s: %s\n", level, message)
}

type Config struct {
	ConnectionString string
}

type DatabaseService struct {
	Config Config
}

func (db DatabaseService) Connect() {
	fmt.Println("Connecting with:", db.Config.ConnectionString)
}

type LoggingService struct {
	Config Config
}

func (ls LoggingService) Log(message string) {
	fmt.Println("Log to DB with:", ls.Config.ConnectionString, message)
}
