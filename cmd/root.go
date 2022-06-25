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
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		dc, err := cmd.Flags().GetBool("disable-check")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		n := "1234567890"
		l := "abcdefghijklmnopqrstuvwxyz"
		u := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		s := "()[]!?$+=/"
		base := n + l + u + s
		var result string
		for {
			result = ""
			for i := 0; i < length; i++ {
				n, err := rand.Int(rand.Reader, big.NewInt(int64(len(base))))
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(2)
				}
				result += string(base[int(n.Int64())])
			}
			if dc {
				break
			}
			if len(result) < 4 || (check(result, n) && check(result, l) && check(result, u) && check(result, s)) {
				break
			}
		}
		fmt.Printf(result)
	},
}

func check(target string, chars string) bool {
	for _, v := range target {
		for _, c := range chars {
			if v == c {
				return true
			}
		}
	}
	return false
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("length", "l", 16, "password length")
	rootCmd.Flags().BoolP("disable-check", "d", false, "disable check")
}
