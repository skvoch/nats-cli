package cmd

import (
	"github.com/skvoch/nats-cli/pkg/constants"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/skvoch/nats-cli/internal/subscribe"
	"github.com/skvoch/nats-cli/internal/template"
)

var subscribeVars struct {
	natsServer     string
	natsSubject    string
	natsClusterID  string
	natsClientID   string
	templateName   string
	startDeltaFrom time.Duration
}

var subscribeTemplateCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"tpl"},
	Short:   "Use template for subscribe",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		tpl, err := template.Get(subscribeVars.templateName)
		if err != nil {
			logrus.Error(err)
			return
		}

		if err := subscribe.Run(
			tpl.NatsServer,
			tpl.NatsClusterID,
			tpl.NatsClientID,
			tpl.NatsSubject,
			subscribeVars.startDeltaFrom); err != nil {
			logrus.Error(err)
			return
		}
	},
}

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "Subscribe to subject",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := subscribe.Run(
			subscribeVars.natsServer,
			subscribeVars.natsClusterID,
			subscribeVars.natsClientID,
			subscribeVars.natsSubject,
			subscribeVars.startDeltaFrom); err != nil {
			logrus.Error(err)
			return
		}
	},
}

func init() {
	subscribeTemplateCmd.Flags().StringVarP(&subscribeVars.templateName, "name", "n", "", "template name")
	subscribeTemplateCmd.Flags().DurationVarP(&subscribeVars.startDeltaFrom, "delta-time", "d", 0, "cluster id")
	if err := subscribeTemplateCmd.MarkFlagRequired("name"); err != nil {
		logrus.Fatal(err)
	}

	subscribeCmd.AddCommand(subscribeTemplateCmd)

	subscribeCmd.Flags().StringVarP(&subscribeVars.natsServer, "addr", "a", "", "NATS server addr")
	subscribeCmd.Flags().StringVarP(&subscribeVars.natsSubject, "subject", "s", "", "subject name")
	subscribeCmd.Flags().StringVarP(&subscribeVars.natsClusterID, "cluster-id", "c", "", "cluster id")
	subscribeCmd.Flags().StringVarP(&subscribeVars.natsClientID, "client-id", "i", constants.ClientID, "client id")
	subscribeCmd.Flags().DurationVarP(&subscribeVars.startDeltaFrom, "delta-time", "d", 0, "delta-time")

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
