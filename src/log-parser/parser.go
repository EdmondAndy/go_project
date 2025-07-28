package main

import (
	"fmt"
	"strconv"
	"strings"
)

type parser struct {
	sum     map[string]result // total visits per domain
	domains []string       // unique domain names
	total   int            // total visits to all domains
	lines   int            // number of parsed lines (for the error messages)
	lerr	error          // last error
}

type result struct{
	domain string
	visits int
}

func newParser() *parser {
	return &parser{sum: make(map[string]result)}
}

func parse(p *parser, line string) (r result) {
	if p.lerr != nil {
		return
	}
	p.lines++

	// Parse the fields
	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return r
	}

	var err error

	r.domain = fields[0]

	// Sum the total visits per domain
	r.visits, err = strconv.Atoi(fields[1])
	if r.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
		return r
	}

	return r
}

func update(p *parser, r result) {
	if p.lerr != nil {
		return
	}

	// Collect the unique domains
	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	// Keep track of total and per domain visits
	p.total += r.visits

	p.sum[r.domain] = result{
		domain: r.domain,
		visits: r.visits + p.sum[r.domain].visits,
	}
}

func err(p *parser) error {
	return p.lerr
}