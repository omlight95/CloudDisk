package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
