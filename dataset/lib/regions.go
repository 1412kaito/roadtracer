package lib

import (
	"github.com/mitroadmaps/gomapinfer/common"
	"github.com/mitroadmaps/gomapinfer/googlemaps"

	"math"
)

const ZOOM = 18

var regionMap map[string][4]float64 = map[string][4]float64{
// 	"surabaya-tengah-kota": [4]float64{-7.259136, 112.747124, -7.259136, 112.747124},
// 	"alun-alun-malang": [4]float64{-7.97710, 112.63412, -7.97710, 112.63412},
// 	"chicago": [4]float64{41.903117, -87.679072, 41.90, -87.67},
	"sby-atum": [4]float64{-7.241825, 112.743753, -7.241825, 112.743753},
}

type Region struct {
	Name string
	RadiusX int
	RadiusY int
	CenterGPS common.Point
	CenterWorld common.Point
}

func GetRegions() []Region {
	var regions []Region
	for name, array := range regionMap {
		centerGPS := common.Point{
			(array[1] + array[3]) / 2,
			(array[0] + array[2]) / 2,
		}
		extreme := googlemaps.LonLatToPixel(common.Point{array[1], array[0]}, centerGPS, ZOOM)
		radiusX := int(math.Ceil(math.Abs(extreme.X) / 4096))
		radiusY := int(math.Ceil(math.Abs(extreme.Y) / 4096))
		if name == "denver" || name == "kansas city" || name == "san diego" || name == "pittsburgh" || name == "montreal" || name == "vancouver" || name == "tokyo" || name == "saltlakecity" || name == "paris" || name == "amsterdam" {
			radiusX = 1
			radiusY = 1
		}
		regions = append(regions, Region{
			Name: name,
			RadiusX: radiusX,
			RadiusY: radiusY,
			CenterGPS: centerGPS,
			CenterWorld: googlemaps.LonLatToMeters(centerGPS),
		})
	}
	/*
	regions = append(regions, Region{
		Name: "boston",
		RadiusX: 3,
		RadiusY: 3,
		CenterGPS: common.Point{-71.107810, 42.352373},
		CenterWorld: googlemaps.LonLatToMeters(common.Point{-71.107810, 42.352373}),
	})
	*/
	return regions
}
