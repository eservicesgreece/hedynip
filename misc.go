package main

import "fmt"

func fullversion() string {
	return fmt.Sprintf("hedynip v%s\n Build: %s\n GIT: %s\n Copyright (c) 2018 eServices Greece - https://esgr.in\n", version, buildstamp, githash)
}
