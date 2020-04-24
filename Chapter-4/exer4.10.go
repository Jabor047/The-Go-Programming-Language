// Issues prints a table of Github issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"../Chapter-4/github"
)

func main(){
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	pastDayIssues := make([]*github.Issue, 0)
	pastMonthIssues := make([]*github.Issue, 0)
	pastYearIssues := make([]*github.Issue, 0)

	// add a -1 to the respective time.Now.AddDate year, month, day fields
	pastDay := now.AddDate(0, 0, -1)
	pastMonth := now.AddDate(0, -1, 0)
	pastYear := now.AddDate(-1, 0, 0)

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(pastDay):
			pastDayIssues = append(pastDayIssues, item)
		case item.CreatedAt.After(pastMonth):
			pastMonthIssues = append(pastMonthIssues, item)
		case item.CreatedAt.After(pastYear):
			pastYearIssues = append(pastYearIssues, item)
		}
	}
	if len(pastDayIssues) > 0 {
		fmt.Println("\n Past Day Issues \n")
		for _, item := range pastDayIssues{
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number,
				  item.User.Login, item.Title)
		}
		
	}
	if len(pastMonthIssues) > 0 {
		fmt.Println("\n Past Month Issues \n")
		for _, item := range pastMonthIssues{
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number,
				  item.User.Login, item.Title)
		}
		
	}
	if len(pastYearIssues) > 0 {
		fmt.Println("\n Past Year Issues \n")
		for _, item := range pastYearIssues{
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number,
				  item.User.Login, item.Title)
		}
		
	}   
}