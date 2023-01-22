package config

import "flag"

type Config struct {
	InputOpts InputOptions
}

type InputOptions struct {
	InFile       string
	OfFile       string
	Scheme       string
	Domain       string
	Wordlist     string
	Silent       bool
	PermuteLevel int
}

func (c *Config) ParseFlags() {
	flag.StringVar(&c.InputOpts.InFile, "i", c.InputOpts.InFile, "Path to the input URLs file")
	flag.StringVar(&c.InputOpts.OfFile, "o", c.InputOpts.OfFile, "Path to the output file")
	flag.StringVar(&c.InputOpts.Domain, "d", c.InputOpts.Domain, "The target domain if wordlist is supplied or if supplied paths have different domains.")
	flag.StringVar(&c.InputOpts.Wordlist, "w", c.InputOpts.Wordlist, "Path to the wordlist file")
	flag.BoolVar(&c.InputOpts.Silent, "s", c.InputOpts.Silent, "Silent mode")
	flag.IntVar(&c.InputOpts.PermuteLevel, "level", 3, "level of permutation")
	flag.IntVar(&c.InputOpts.PermuteLevel, "l", 3, "level of permutation (shorthand)")
	flag.Parse()
}
