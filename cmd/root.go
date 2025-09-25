package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	jsonstore "github.com/lulu73211/go_crm_tp3/internal/store/json"

	"github.com/lulu73211/go_crm_tp3/config"
	"github.com/lulu73211/go_crm_tp3/internal/service"
	"github.com/lulu73211/go_crm_tp3/internal/store"
	mem "github.com/lulu73211/go_crm_tp3/internal/store/memory"
)

var (
	cfgPath string
	app     *service.CRM
)

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "Mini-CRM CLI (étape 2 : Viper + add + mémoire)",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(cfgPath)
		if err != nil {
			return err
		}

		var st store.Storer
		switch cfg.Type {
		case "json":
			st = jsonstore.New(cfg.JSONPath)
		case "memory":
			st = mem.New()
		default:
			return fmt.Errorf("unknown storage type: %s", cfg.Type)
		}

		if err := st.Init(); err != nil {
			return err
		}
		app = service.New(st)
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
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
