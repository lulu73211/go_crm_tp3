package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

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
		// Charger config
		cfg, err := config.Load(cfgPath)
		if err != nil {
			return err
		}

		// Pour l’instant, on ne supporte que mémoire (cfg.Type = "memory")
		var st store.Storer = mem.New()
		if err := st.Init(); err != nil {
			return err
		}
		app = service.New(st)
		_ = cfg // on le garde pour la suite (JSON/GORM)
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
