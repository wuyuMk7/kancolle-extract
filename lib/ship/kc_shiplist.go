package ship

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	utils "github.com/wuyuMk7/kancolle-extract/lib/download"
)

type KC struct {
	Data              KCData `json:"api_data"`
	ResponseResult    int    `json:"api_result"`
	ResponseResultMsg string `json:"api_result_msg"`
}

type KCData struct {
	KCShips []KCShip `json:"api_mst_ship"`
}

type KCShip struct {
	ID   int    `json:"api_id"`
	Name string `json:"api_name"`
}

var resources = [...]int{6657, 5699, 3371, 8909, 7719, 6229, 5449, 8561, 2987, 5501, 3127, 9319, 4365, 9811, 9927, 2423, 3439, 1865, 5925, 4409, 5509, 1517, 9695, 9255, 5325, 3691, 5519, 6949, 5607, 9539, 4133, 7795, 5465, 2659, 6381, 6875, 4019, 9195, 5645, 2887, 1213, 1815, 8671, 3015, 3147, 2991, 7977, 7045, 1619, 7909, 4451, 6573, 4545, 8251, 5983, 2849, 7249, 7449, 9477, 5963, 2711, 9019, 7375, 2201, 5631, 4893, 7653, 3719, 8819, 5839, 1853, 9843, 9119, 7023, 5681, 2345, 9873, 6349, 9315, 3795, 9737, 4633, 4173, 7549, 7171, 6147, 4723, 5039, 2723, 7815, 6201, 5999, 5339, 4431, 2911, 4435, 3611, 4423, 9517, 3243}

func (ship KCShip) suffix(imageType string) string {
	r := ship.ID
	s := 0
	a := 1

	if imageType != "" {
		a = len(imageType)
		for _, ch := range imageType {
			s = s + int(ch)
		}
	}
	suffix := 17*(r+7)*resources[(s+r*a)%100]%8973 + 1000

	return fmt.Sprint(suffix)

}

func (ship KCShip) GetImage(dirname string) error {
	if dirname == "" {
		dirname = "."
	}

	shipPrefix := fmt.Sprintf("%04d", ship.ID)

	baseUrl := "http://125.6.189.71/kcs2/resources/ship/"
	imageUrls := map[string]string{
		"card":    baseUrl + "card/" + shipPrefix + "_" + ship.suffix("ship_card") + ".png",
		"cardDmg": baseUrl + "card_dmg/" + shipPrefix + "_" + ship.suffix("ship_card_dmg") + ".png",
		"full":    baseUrl + "full/" + shipPrefix + "_" + ship.suffix("ship_full") + ".png",
		"fullDmg": baseUrl + "full_dmg/" + shipPrefix + "_" + ship.suffix("ship_full_dmg") + ".png",
	}

	log.Println("Downloading ship id ", ship.ID, "...")

	baseDir := fmt.Sprint(dirname, "/", ship.ID, "_", ship.Name, "/")
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		if err = os.MkdirAll(baseDir, 0744); err != nil {
			return err
		}
	}

	nameFile, err := os.Create(fmt.Sprint(baseDir, "name"))
	if err != nil {
		return err
	}
	defer nameFile.Close()
	nameFile.WriteString(ship.Name)

	for key, url := range imageUrls {
		log.Print(key, " image ...")
		err := utils.Download(url, fmt.Sprint(baseDir, key, ".png"))
		if err != nil {
			return err
		}
		log.Println("Done")
	}

	return nil
}

func (kc *KC) LoadInfo(dataFileName string) error {
	if dataFileName == "" {
		return errors.New("No data file specified.")
	}

	dataFile, err := os.Open(dataFileName)
	if err != nil {
		return err
	}
	defer dataFile.Close()
	log.Println("Successfully loaded datafile " + dataFileName)

	data, err := ioutil.ReadAll(dataFile)

	err = json.Unmarshal(data, kc)
	if err != nil {
		return err
	}

	return nil
}

func (kc KC) GetImage(imageDirName string) error {
	if imageDirName == "" {
		imageDirName = "."
	}
	if _, err := os.Stat(imageDirName); os.IsNotExist(err) {
		if err = os.MkdirAll(imageDirName, 0744); err != nil {
			return err
		}
	}

	if len(kc.Data.KCShips) > 0 {
		log.Println("Extracting images...")
		for _, ship := range kc.Data.KCShips {
			if err := ship.GetImage(imageDirName); err != nil {
				log.Print(fmt.Sprint(ship.ID, " image extraction failed."))
				log.Println(err)
			}
		}
	} else {
		log.Println("No ship information input. Abort.")
	}

	return nil
}
