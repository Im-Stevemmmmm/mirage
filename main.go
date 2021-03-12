package main

import "github.com/Im-Stevemmmmm/mirage/vcs"

func main() {
	c := vcs.Commit{}
	c.RevertTo(vcs.Hard)
}
