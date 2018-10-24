redmine
=======

A go package that can be used to access the Redmine API (v2.2.3)

$ go get github.com/woli/redmine

#### Usage example:

	package main

	import (
		"fmt"
		"github.com/woli/redmine"
		"net/http"
	)

	func main() {
		auth := &redmine.ApiKeyAuth{"<apikey>"}
		s, _ := redmine.New("<baseurl>", auth, http.DefaultClient)

		feed, err := s.Issues.List().Do()
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("TotalCount:%v Limit:%v Offset:%v\n", feed.TotalCount, feed.Limit, feed.Offset)
			for _, issue := range feed.Issues {
				fmt.Printf("%+v\n", issue)
			}
		}
		
		// Example adding an issue
		assignedTo := redmine.Name{
			Id: 1, // Map these to real users
		}
		status := redmine.Name{
			Id: 3, // Resolved
		}
		currentVersion := redmine.Name{
			Id: 1, // some release id
		}
		projectName := redmine.Name{
			Id: 1, // some project id
		}
		author := redmine.Name{
			Id: 1, // some user
		}
		tracker := redmine.Name{
			Id: 2, // feature
		}
		newIssue := redmine.Issue{
			DoneRatio:    100,
			Status:       &status,
			AssignedTo:   &assignedTo,
			Author:       &author,
			Tracker:      &tracker,
			Project:      &projectName,
			FixedVersion: &currentVersion,
			Subject:      "blah test",
			Description:  "blah description",
		}

		customField := redmine.CustomField{
			Id:    3,
			Name:  "FW Version",
			Value: "1234",
		}

		newIssue.CustomFields = append(newIssue.CustomFields, &customField)
		_, _ = s.Issues.Insert(&newIssue).Do()
		
		// Example adding a comment to an issue
		
		
		issue, err := s.Issues.Get(6897).Journals(true).Do()
		if err == nil {
			issue.Notes = "Some new note and comment"
			err = s.Issues.Update(issue).Do()
			if err == nil {
				fmt.Println("this worked")
			}
		}
	}
