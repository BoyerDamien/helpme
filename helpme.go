package main

import (
	"flag"
	"fmt"
)

func main() {
	sorted := flag.Bool("s", false, "Sort elements by score")
	help := flag.Bool("h", false, "Help")
	number := flag.Int("n", 0, "Number of elements to display")
	flag.Parse()
	if (len(flag.CommandLine.Args()) == 0) || (*help) {
		fmt.Println(helpmeUsage)
	} else if len(flag.Args()) == 1 {
		tail := flag.Args()
		req := Request{content: tail[0], url: first, sort: *sorted, n: *number}
		req.buildURL()
		req.fetch()
	} else {
		fmt.Println(formatError)
	}
}
