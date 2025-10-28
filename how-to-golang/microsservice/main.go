package main

import (
	"log"
)

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)

	// fact, err := svc.GetCatFact(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", fact.Fact)

	apiServer := NewApiServer(svc)
	log.Fatal(apiServer.Start(":3000"))

}
