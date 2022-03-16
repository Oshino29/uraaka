package main

import (
	// "oshino29/uraaka/storage"
	"testing"
)

// just a hack to run function loadCensorList() to load censorWords into database
func TestCensor(t *testing.T) {
	loadCensorList()
}