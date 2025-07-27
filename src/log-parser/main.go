package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type parser struct {
	sum     map[string]result // total visits per domain
	domains []string       // unique domain names
	total   int            // total visits to all domains
	lines   int            // number of parsed lines (for the error messages)
}

type result struct{
	domain string
	visits int
}

func main() {

	p := parser{
		sum: make(map[string]result),
	}



	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		// Parse the fields
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, p.lines)
			return
		}

		domain := fields[0]

		// Sum the total visits per domain
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], p.lines)
			return
		}

		// Collect the unique domains
		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}

		// Keep track of total and per domain visits
		p.total += visits

		p.sum[domain] = result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}
	}

	// Print the visits per domain
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

	// Let's handle the error
	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}