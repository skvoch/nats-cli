package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/skvoch/nats-cli/internal/template"
	"github.com/spf13/cobra"
)

type TemplateVars struct {
	natsServer    string
	natsSubject   string
	natsClusterID string
	name          string
}

var templateVars TemplateVars

// publishCmd represents the publish command
var templateCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"tpl"},
	Short:   "Manage templates",
	Args:    cobra.MinimumNArgs(1),
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var templateRemove = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove template",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := template.Remove(templateVars.name); err != nil {
			logrus.Error(err)
			return
		}
		if err := template.Save(); err != nil {
			logrus.Error(err)
			return
		} else {
			logrus.Info("template has been removed successfully")
		}
	},
}

var templateCreate = &cobra.Command{
	Use:     "create",
	Aliases: []string{"cr"},
	Short:   "Create template",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := template.Add(templateVars.name, template.Item{
			NatsClusterID: templateVars.natsClusterID,
			NatsServer:    templateVars.natsServer,
			NatsSubject:   templateVars.natsSubject,
		}); err != nil {
			logrus.Error(err)
			return
		}

		if err := template.Save(); err != nil {
			logrus.Error(err)
		} else {
			logrus.Info("template has been added successfully")
		}
	},
}

var templateList = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List templates",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		templates := template.List()

		for name, item := range templates.Values {
			str := fmt.Sprintf(`name: %s | server: %s | cluster-id: %s | subject: %s`,
				name, item.NatsServer, item.NatsClusterID, item.NatsSubject)
			logrus.Info(str)
		}

		if len(templates.Values) == 0 {
			logrus.Info("there aren't templates yet")
		}
	},
}

func init() {
	templateCmd.AddCommand(templateList)
	templateCmd.AddCommand(templateCreate)

	templateCreate.Flags().StringVarP(&templateVars.name, "name", "n", "", "template name")
	templateCreate.Flags().StringVarP(&templateVars.natsServer, "addr", "a", "", "NATS server addr")
	templateCreate.Flags().StringVarP(&templateVars.natsSubject, "subject", "s", "", "subject name")
	templateCreate.Flags().StringVarP(&templateVars.natsClusterID, "cluster-id", "c", "", "cluster id")

	if err := templateCreate.MarkFlagRequired("addr"); err != nil {
		logrus.Fatal(err)
	}
	if err := templateCreate.MarkFlagRequired("subject"); err != nil {
		logrus.Fatal(err)
	}
	if err := templateCreate.MarkFlagRequired("cluster-id"); err != nil {
		logrus.Fatal(err)
	}
	if err := templateCreate.MarkFlagRequired("name"); err != nil {
		logrus.Fatal(err)
	}

	templateCmd.AddCommand(templateRemove)
	templateRemove.Flags().StringVarP(&templateVars.name, "name", "n", "", "template name")
	if err := templateRemove.MarkFlagRequired("name"); err != nil {
		logrus.Fatal(err)
	}

	rootCmd.AddCommand(templateCmd)
}
