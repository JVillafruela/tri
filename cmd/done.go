//Package cmd command line
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/JVillafruela/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark item as done",
	Run:     doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doneRun(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Fatalf("Read items : %v\n", err)
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label", err)
	}
	if i > 0 && i < len(items) {
		items[i-1].Done = true
		fmt.Printf("%s %v\n", items[i-1].Text, "marked done")
		sort.Sort(todo.ByPri(items))
		err = todo.SaveItems(dataFile, items)
		if err != nil {
			log.Fatalf("Save items : %v\ns", err)
		}
	} else {
		log.Println(i, "doesnt match any item")
	}
}
