package test

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")

	m.Run() //eksekusi semua unit test

	fmt.Println("Setelah Unit Test")
}
