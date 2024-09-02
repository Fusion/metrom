package net

import "net"

type Resolver struct {
	cache map[string][]string
}

func NewResolver() Resolver {
	return Resolver{
		cache: make(map[string][]string),
	}
}

func (r *Resolver) ResolveAddress(address string) []string {
	resolved, ok := r.cache[address]
	if !ok {
		names, _ := net.LookupAddr(address)
		if len(names) == 0 {
			resolved = []string{address}
		} else {
			resolved = names
		}
		r.cache[address] = resolved
	}
	return resolved
}

func (r *Resolver) Cleanup() {
	r.cache = make(map[string][]string)
}
