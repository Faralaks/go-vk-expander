package main

import (
	"github.com/umputun/go-flags"
	"log"
	"os"
)

var config struct {
	App struct {
		Port    string `long:"port" env:"PORT" default:"80" description:"string | Application port"`
		Address string `long:"address" env:"ADDRESS" default:"0.0.0.0" description:"string | Application web address"`
	} `group:"app" namespace:"app" env-namespace:"VKEXP_APP"`
	OAuth struct {
		ClientId     string `long:"client_id" env:"CLIENT_ID" required:"true"  description:"string | VK OAuth Client ID"`
		Key          string `long:"key" env:"KEY" required:"true"  description:"string | VK OAuth Key"`
		RedirectPath string `long:"redirect_path" env:"REDIRECT_PATH" required:"true"  description:"string | VK OAuth Redirect path like \"/save_token\""`
		RedirectURL  string `long:"redirect_url" env:"REDIRECT_URL" required:"true"  description:"string | VK OAuth full Redirect URL"`
	} `group:"oauth" namespace:"oauth" env-namespace:"VKEXP_OAUTH"`
}

func main() {
	println("\t ---> Faralaks Vk-Expander starting!")

	// Loading configuration
	p := flags.NewParser(&config, flags.PrintErrors|flags.PassDoubleDash|flags.HelpFlag)
	if _, err := p.Parse(); err != nil {
		if err.(*flags.Error).Type != flags.ErrHelp {
			log.Printf("[ERROR] cli error: %v", err)
		}
		os.Exit(2)
	}
	os.Clearenv() // Clear Environment. Now only this process has this data

	println("\t <--- Faralaks Vk-Expander finish!!")

}
