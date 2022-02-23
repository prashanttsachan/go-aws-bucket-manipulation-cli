/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/spf13/cobra"
	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    config "bucket/config"
)

// uploadFolderCmd represents the uploadFolder command
var uploadFolderCmd = &cobra.Command{
	Use:   "uploadFolder",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Uploading your folder...")
		env := config.Getenv()
		sess := config.ConnectAws();
		// Create S3 service client
		folder := args[0]

		files, err := ioutil.ReadDir(folder)
		if err != nil {
			fmt.Println("Error in opening folder: %q", err)
		}
		for _, object := range files {
			file, err := os.Open(folder + "/" + object.Name())
			if err != nil {
				config.ExitErrorf("Unable to open file %q, %v", err)
			}
			defer file.Close()

			uploader := s3manager.NewUploader(sess)
			success, err := uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String(env.Bucket),
				Key: aws.String(env.Folder + "/" + folder + "/" + object.Name()),
				Body: file,
			})
			if err != nil {
				// Print the error and exit.
				config.ExitErrorf("Unable to upload %q to %q, %v", object.Name(), env.Bucket, err)
			}
			fmt.Printf("%q\n", success)
			fmt.Println("File uploaded successfully %q to %q\n", object.Name(), env.Bucket)
		}
		fmt.Println("All files uploaded successfully.\n")
	},
}

func init() {
	rootCmd.AddCommand(uploadFolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadFolderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadFolderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
