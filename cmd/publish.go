package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/skvoch/nats-cli/internal/publish"
	"github.com/skvoch/nats-cli/internal/template"
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

var publishCmdTemplateCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"tpl"},
	Short:   "Use template for publish",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		tpl, err := template.Get(publishVars.templateName)
		if err != nil {
			logrus.Error(err)
			return
		}

		if err := publish.Run(
			tpl.NatsServer,
			tpl.NatsClusterID,
			tpl.NatsSubject,
			[]byte(publishVars.message),
			true); err != nil {
			logrus.Error(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	publishCmdTemplateCmd.Flags().StringVarP(&publishVars.message, "message", "m", "", "JSON message")
	publishCmdTemplateCmd.Flags().StringVarP(&publishVars.templateName, "name", "n", "", "template name")

	publishCmd.Flags().StringVarP(&publishVars.natsServer, "addr", "a", "", "NATS server addr")
	publishCmd.Flags().StringVarP(&publishVars.natsSubject, "subject", "s", "", "subject name")
	publishCmd.Flags().StringVarP(&publishVars.natsClusterID, "cluster-id", "c", "", "cluster id")
	publishCmd.Flags().StringVarP(&publishVars.message, "message", "m", "", "JSON message")

	publishCmd.AddCommand(publishCmdTemplateCmd)

	if err := publishCmdTemplateCmd.MarkFlagRequired("name"); err != nil {
		logrus.Fatal(err)
	}

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

	if err := publishCmdTemplateCmd.MarkFlagRequired("message"); err != nil {
		logrus.Fatal(err)
	}
}
