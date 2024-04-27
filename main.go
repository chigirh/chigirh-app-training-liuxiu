package main

import (
	"chigirh-app-trainning-liuxiu/conf/drivers"
	"context"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	d, err := drivers.InitializeDriver(ctx)
	if err != nil {
		log.Printf("failed to create UserDriver: %s\n", err)
		os.Exit(2)
	}

	d.Start(ctx)
}
