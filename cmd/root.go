package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/lulu73211/go_crm_tp3/config"
	"github.com/lulu73211/go_crm_tp3/internal/service"
	"github.com/lulu73211/go_crm_tp3/internal/store"
	gormstore "github.com/lulu73211/go_crm_tp3/internal/store/gorm"
	jsonstore "github.com/lulu73211/go_crm_tp3/internal/store/json"
	mem "github.com/lulu73211/go_crm_tp3/internal/store/memory"
)

var (
	cfgPath string
	app     *service.CRM
)

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "Mini-CRM CLI (contacts CRUD)",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Charger la configuration
		cfg, err := config.Load(cfgPath)
		if err != nil {
			return err
		}

		// Choisir le backend de persistance
		var st store.Storer
		switch cfg.Type {
		case "gorm":
			st = gormstore.New(cfg.DBPath)
		case "json":
			st = jsonstore.New(cfg.JSONPath)
		case "memory":
			st = mem.New()
		default:
			return fmt.Errorf("unknown storage type: %s", cfg.Type)
		}

		// Initialiser le store et construire le service
		if err := st.Init(); err != nil {
			return err
		}
		app = service.New(st)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Sans sous-commande, afficher l'aide
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cwd, _ := os.Getwd()
	defaultCfg := filepath.Join(cwd, "config.yaml")
	rootCmd.PersistentFlags().StringVar(&cfgPath, "config", defaultCfg, "path to config file")
}
