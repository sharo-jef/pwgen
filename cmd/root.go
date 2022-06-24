/*
Copyright © 2022 sharo-jef

*/
package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pwgen",
	Short: "Generate password.",
	Long:  `Generate password.`,
	Run: func(cmd *cobra.Command, args []string) {
		l, err := cmd.Flags().GetInt("length")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		base := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ()[]!?$+=/"
		result := ""
		for i := 0; i < l; i++ {
			n, err := rand.Int(rand.Reader, big.NewInt(int64(len(base))))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
			result += string(base[int(n.Int64())])
		}
		fmt.Printf(result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("length", "l", 16, "Password length")
}
