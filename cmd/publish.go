package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/skvoch/nats-cli/internal/publish"
	"github.com/spf13/cobra"
)

var publishVars struct {
	natsServer    string
	natsSubject   string
	natsClusterID string
	templateName  string
	message       string
}

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:     "publish",
	Aliases: []string{"pub"},
	Short:   "Publish to subject",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := publish.Run(
			publishVars.natsServer,
			publishVars.natsClusterID,
			publishVars.natsSubject,
			[]byte(publishVars.message), true); err != nil {
			logrus.Error(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	publishCmd.Flags().StringVarP(&publishVars.natsServer, "addr", "a", "", "NATS server addr")
	publishCmd.Flags().StringVarP(&publishVars.natsSubject, "subject", "s", "", "subject name")
	publishCmd.Flags().StringVarP(&publishVars.natsClusterID, "cluster-id", "c", "", "cluster id")
	publishCmd.Flags().StringVarP(&publishVars.message, "message", "m", "", "JSON message")

	if err := publishCmd.MarkFlagRequired("addr"); err != nil {
		logrus.Fatal(err)
	}

	if err := publishCmd.MarkFlagRequired("subject"); err != nil {
		logrus.Fatal(err)
	}

	if err := publishCmd.MarkFlagRequired("cluster-id"); err != nil {
		logrus.Fatal(err)
	}

	if err := publishCmd.MarkFlagRequired("message"); err != nil {
		logrus.Fatal(err)
	}
}
