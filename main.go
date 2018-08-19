package main

import (
	"fmt"
	"log"
	"os"

	formats "github.com/wuyuMk7/kancolle-extract/lib/ship"
)

func imagesExtraction(dirName string, dataFileName string, format formats.ShipList) {
	log.Println("image extraction process...")
	if err := format.LoadInfo(dataFileName); err != nil {
		log.Fatal("image extraction - load ships info failed - ", err)
	}
	log.Println("ships information loaded.")

	log.Println("images extracting...")
	if err := format.GetImage(dirName); err != nil {
		log.Fatal("image extraction - image download failed - ", err)
	}
	log.Println("images extration accomplished.")
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("KCExtract: ")
	log.SetFlags(log.LstdFlags)

	log.Println("KanColle resources extraction program...")

	var kc formats.KC
	imagesExtraction("./imgs", "getdata.json", &kc)

	log.Println("Extraction accomplished.")
}
