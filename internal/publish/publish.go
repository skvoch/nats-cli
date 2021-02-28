package publish

import (
	"github.com/sirupsen/logrus"
	"github.com/skvoch/nats-cli/pkg/constants"

	natsc "github.com/skvoch/nats-cli/internal/nats"
)

const (
	logNatsAddr = "nats"
)

func Run(address, clusterId, subject string, data []byte, validateJSON bool) error {
	logrus.WithField(logNatsAddr, address).Info("trying connect to nats...")
	nats, err := natsc.Connect(address, clusterId, constants.ClientID)
	if err != nil {
		return err
	}

	logrus.WithField(logNatsAddr, address).Info("trying to publish message...")
	if err := nats.Publish(subject, data, validateJSON); err != nil {
		return err
	}

	logrus.WithField(logNatsAddr, address).Info("message has been published successfully")
	return nil
}
