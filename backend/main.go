package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rg-km/final-project-engineering-12/backend/config"
	"github.com/rg-km/final-project-engineering-12/backend/route"
)

func main() {
	configuration := config.New()
	initialized := route.NewInitializedServer(configuration)
	// Run
	PORT := fmt.Sprintf(":%v", configuration.Get("APP_PORT"))
	teenager(PORT)

	err := initialized.Run(PORT)
	if err != nil {
		panic(err)
	}
}

func teenager(port string) {
	fmt.Print(`
┏━━━━┓
┃┏┓┏┓┃
┗┛┃┃┣┻━┳━━┳━┓┏━━┳━━┳━━┳━┓
╋╋┃┃┃┃━┫┃━┫┏┓┫┏┓┃┏┓┃┃━┫┏┛
╋╋┃┃┃┃━┫┃━┫┃┃┃┏┓┃┗┛┃┃━┫┃
╋╋┗┛┗━━┻━━┻┛┗┻┛┗┻━┓┣━━┻┛
╋╋╋╋╋╋╋╋╋╋╋╋╋╋╋╋┏━┛┃
╋╋╋╋╋╋╋╋╋╋╋╋╋╋╋╋┗━━┛
`)
}
