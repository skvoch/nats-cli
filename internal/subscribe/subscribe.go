package subscribe

import (
	"bytes"
	"encoding/json"
	"github.com/skvoch/nats-cli/pkg/constants"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	natsc "github.com/skvoch/nats-cli/internal/nats"
)

const (
	logNatsAddr    = "nats"
	logNatsSubject = "subject"
)

func Run(address, clusterId, subject string, delta time.Duration) error {
	logrus.WithField(logNatsAddr, address).Info("trying connect to nats...")
	nats, err := natsc.Connect(address, clusterId, constants.ClientID)
	if err != nil {
		return err
	}
	logrus.WithField(logNatsSubject, subject).Info("trying subscribe to subject...")
	messages, err := nats.Subscribe(subject, delta)
	if err != nil {
		return err
	}
	logrus.WithField(logNatsSubject, subject).Info("successful subscription to subject!")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-c:
			close(c)

			logrus.Info("exist")
			return nil

		case msg := <-messages:
			f := &bytes.Buffer{}
			if err := json.Indent(f, msg, "", "  "); err != nil {
				logrus.Error(err)
			}
			logrus.Info(f)
		}
	}
}
