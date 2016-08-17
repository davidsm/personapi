package person

import (
	"strconv"

	"github.com/masenius/personapi/data"
)

type Address struct {
	Street       string        `json:"streetAddress"`
	StreetNumber string        `json:"streetNumber"`
	Code         string        `json:"code"`
	Locality     string        `json:"locality"`
	Municipality *municipality `json:"municipality"`
	Region       *region       `json:"region"`
}

type municipality struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type region struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Letter string `json:"letter"`
}

func RandomAddress() *Address {
	addr := data.Addresses[randgen.Intn(len(data.Addresses))]

	streetNumber := strconv.Itoa(randomBetween(addr.NumMin, addr.NumMax+1))

	muniName := data.Municipalities[addr.MunicipalityCode]
	municipality := municipality{muniName, addr.MunicipalityCode}

	regionData := data.Regions[addr.RegionCode]
	region := region{regionData.Name, addr.RegionCode, regionData.Letter}

	return &Address{
		Street:       addr.Street,
		StreetNumber: streetNumber,
		Code:         addr.Code,
		Locality:     addr.Locality,
		Municipality: &municipality,
		Region:       &region,
	}
}
