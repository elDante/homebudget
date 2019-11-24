package cmd

import (
	"log"

	"github.com/elDante/homebudget/config"
	"github.com/elDante/homebudget/contrib"
	"github.com/elDante/homebudget/database"
	"github.com/elDante/homebudget/models"
	"github.com/spf13/cobra"
)

var fullName string
var email string
var password string

var createUserCmd = &cobra.Command{
	Use:   "createuser",
	Short: "Create user",
	Long:  `This subcommand creates user in database`,
	Run: func(cmd *cobra.Command, args []string) {
		cnf := config.Parse(&ConfigPath)
		db := database.Connector(&cnf.Database)
		user := models.User{
			FullName: fullName,
			Email:    email,
			Password: contrib.SecretString(password, cnf.Site.Secret),
		}
		db.Where(&user).FirstOrCreate(&user)
		log.Println("User created")
	},
}

func init() {
	RootCmd.AddCommand(createUserCmd)
	createUserCmd.PersistentFlags().StringVar(&fullName, "fullname", "", "User fullname")
	createUserCmd.PersistentFlags().StringVar(&email, "email", "", "User email")
	createUserCmd.PersistentFlags().StringVar(&password, "password", "", "User password")
}
