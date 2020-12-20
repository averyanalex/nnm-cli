package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "send",
	Short: "Send message",
	Long:  `nnm-cli send [message]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		req, err := http.NewRequest("GET", "http://localhost:8080/send/msg", nil)
		q := req.URL.Query()
		q.Add("content", args[0])
		req.URL.RawQuery = q.Encode()
		fmt.Println(req.URL.String())
		client := &http.Client{}
		resp, err := client.Do(req)
		if err == nil && resp.StatusCode == 200 {
			fmt.Println("Succesfully sent!")
		} else {
			fmt.Println("Error occured!")
			fmt.Println(err)
			fmt.Println(resp.Status)
			fmt.Println(resp.Body)
		}
	},
}

func init() {
	RootCmd.AddCommand(helloCmd)
}
