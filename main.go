package main

import (
	"github.com/bary822/gomemon-server/internal/application"
	web_application "github.com/bary822/gomemon-server/internal/application/web"
)

type Startup struct {
	launcher application.Launcher
}

func (s Startup) NewStartUp() Startup {
	return Startup{
		launcher: web_application.WebLauncher{},
	}
}

func main() {
	s := Startup{}.NewStartUp()
	s.launcher.Launch()
}
