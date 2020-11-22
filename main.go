package main

import (
	"github.com/sirupsen/logrus"

	"github.com/skvoch/nats-cli/cmd"
	"github.com/skvoch/nats-cli/internal/template"
)

func main() {
	if err := template.Read(); err != nil {
		logrus.Warn(err)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	cmd.Execute()
}
