/*
 * HomeWork-9: Integration tests
 * Created on 13.12.2019 13:08
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("Wait 5s for service availability...")
	time.Sleep(5 * time.Second)
	os.Exit(2)
}
