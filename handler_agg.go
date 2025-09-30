package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.arguments) < 1 {
		return fmt.Errorf("agg expects 1 argument: time_betwen_reqs")
	}

	time_betwen_reqs, err := time.ParseDuration(cmd.arguments[0])

	if err != nil {
		return fmt.Errorf("time invalid:\n%v", err)
	}

	fmt.Printf("Collecting feeds every %v\n", time_betwen_reqs)

	ticker:=time.NewTicker(time_betwen_reqs)

	for ;;<-ticker.C{
		scrapeFeeds(s)
	}

}
