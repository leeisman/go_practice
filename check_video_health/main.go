package main

import (
	"context"
	"fmt"
	"gopkg.in/vansante/go-ffprobe.v2"
	"log"
	"os"
	"time"
)

func main() {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	fileName := "/Users/frankieli/goProjects/src/gitlab.silkrode.com.tw/team_golang/mobile_lib/libs/test/universal_download/download/e2c808d0-0276-4999-bb93-7cab0d53677b.mp4"
	fileReader, err := os.Open(fileName)
	if err != nil {
		log.Panicf("Error opening test file: %v", err)
	}

	data, err := ffprobe.ProbeReader(ctx, fileReader)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}
	fmt.Print(data)
}
