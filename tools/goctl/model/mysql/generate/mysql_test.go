package generate

import (
	"log"
	"testing"
)

func TestMysqlt(t *testing.T) {
	err := Do(&Context{
		DataSource: "ugozero@tcp(127.0.0.1:3306)/gozero",
		Pattern:    "user",
		File:       "$GOCTLHOME",
		Output:     ".",
	})

	if err != nil {
		log.Fatal(err)
	}
}
