package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs (e.g. 1m)>", cmd.Name)
	}

	timeArg := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(timeArg)
	if err != nil {
		return fmt.Errorf("couldn't parse time: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
