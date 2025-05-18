package main

import (
	"fmt"

	"github.com/yesetoda/bet365-evaluator-go/excuter/cricket_excuter"
	"github.com/yesetoda/bet365-evaluator-go/excuter/volleyball_excuter"
)

func main() {
	cricket_excuter.CricketExecutor()
	fmt.Println("Cricket evaluation completed.")
	fmt.Println("_________________________________________________________________________________________________________________________________")
	fmt.Println("Starting Volleyball evaluation...")
	volleyball_excuter.VolleyballExecutor()
	fmt.Println("_________________________________________________________________________________________________________________________________")
	fmt.Println("Volleyball evaluation completed.")	
}
