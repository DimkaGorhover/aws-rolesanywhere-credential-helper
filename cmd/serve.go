package cmd

import (
	helper "github.com/aws/rolesanywhere-credential-helper/aws_signing_helper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	host     string
	port     int
	loglevel = "info"
)

func init() {
	initCredentialsSubCommand(serveCmd)
	serveCmd.PersistentFlags().StringVar(&host, "host", helper.LocalHostAddress, "The host used to run the server (default: 127.0.0.1)")
	serveCmd.PersistentFlags().IntVar(&port, "port", helper.DefaultPort, "The port used to run the server (default: 9911)")
	serveCmd.PersistentFlags().StringVar(&loglevel, "log-level", "info", "logging level")
}

var serveCmd = &cobra.Command{
	Use:   "serve [flags]",
	Short: "Serve AWS credentials through a endpoint",
	Long:  "Serve AWS credentials through a endpoint that is compatible with IMDSv2",
	Run: func(cmd *cobra.Command, args []string) {
		err := PopulateCredentialsOptions()
		if err != nil {
			log.WithError(err).Fatal("populate credentials options error")
		}

		log.SetFormatter(&log.JSONFormatter{
			PrettyPrint: false,
		})

		level, err := log.ParseLevel(loglevel)
		if err != nil {
			level = log.WarnLevel
		}
		log.SetLevel(level)
		if credentialsOptions.Debug {
			log.SetLevel(log.DebugLevel)
		}

		helper.Debug = credentialsOptions.Debug

		helper.Serve(host, port, credentialsOptions)
	},
}
