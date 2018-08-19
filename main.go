package main

import (
	"log"
	"os"

	formats "github.com/wuyuMk7/kancolle-extract/lib/ship"
)

func imagesExtraction(dirName string, format formats.ShipList) {
	log.Println("image extraction process...")
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

	var dataFileName string
	var kc formats.KC

	dataFileName = "getdata.json"
	log.Println("loading data...")
	if err := kc.LoadInfo(dataFileName); err != nil {
		log.Fatal("data load failed - ", err)
	}
	log.Println("data loaded.")

	imagesExtraction("./imgs", &kc)

	log.Println("Extraction accomplished.")
}
