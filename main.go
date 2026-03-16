package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func mask(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	reveal := int(math.Ceil(float64(n) * 0.10))

	if reveal*2 >= n {
		return strings.Repeat("*", n)
	}

	prefix := s[:reveal]
	suffix := s[n-reveal:]
	masked := strings.Repeat("*", n-reveal*2)

	return prefix + masked + suffix
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: echos <env_var>")
		os.Exit(1)
	}

	name := os.Args[1]
	value, ok := os.LookupEnv(name)
	if !ok || value == "" {
		fmt.Fprintf(os.Stderr, "env var %q is not set or empty\n", name)
		os.Exit(1)
	}

	fmt.Println(mask(value))
}
