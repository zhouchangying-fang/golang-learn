package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Printf("%s", sql.Drivers())
}
