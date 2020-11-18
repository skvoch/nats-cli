package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	natsc "github.com/skvoch/nats-cli/internal/nats"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var natsServer string
var natsSubject string
var natsClusterID string

var startDeltaFrom time.Duration

const (
	logNatsAddr = "nats"
	logNatsSubject = "subject"
)

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Aliases: []string{"sub"},
	Short: "Subscribe to subject",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.WithField(logNatsAddr, natsServer).Info("trying connect to nats...")
		nats, err := natsc.Connect(natsServer,natsClusterID, clientID())
		if err != nil {
			existWithError(err)
		}
		logrus.WithField(logNatsSubject, natsSubject).Info("trying subscribe to subject...")
		messages, err := nats.Subscribe(natsSubject, startDeltaFrom);
		if err != nil {
			existWithError(err)
		}
		logrus.WithField(logNatsSubject, natsSubject).Info("successful subscription to subject!")


		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		for {
			select {
				case <-c:
					logrus.Info("exist")
					return

				case msg := <-messages:
					f := &bytes.Buffer{}
					if err := json.Indent(f, msg, "", "  "); err != nil {
						logrus.Error(err)
					}
					logrus.Info(f)
			}
		}
	},
}

func clientID() string {
	return "nats-cli"
}

func existWithError(err error) {
	logrus.Fatal(err)
}

func init() {
	subscribeCmd.Flags().StringVarP(&natsServer,"addr", "a",  "","NATS server addr")
	subscribeCmd.Flags().StringVarP(&natsSubject,"subject", "s",  "","subject name")
	subscribeCmd.Flags().StringVarP(&natsClusterID,"cluster-id", "c",  "","cluster id")
	subscribeCmd.Flags().DurationVarP(&startDeltaFrom,"delta-time", "d",  0,"cluster id")


	if err := subscribeCmd.MarkFlagRequired("addr"); err != nil {
		logrus.Fatal(err)
	}

	if err := subscribeCmd.MarkFlagRequired("subject"); err != nil {
		logrus.Fatal(err)
	}

	if err := subscribeCmd.MarkFlagRequired("cluster-id"); err != nil {
		logrus.Fatal(err)
	}

	rootCmd.AddCommand(subscribeCmd)
}
