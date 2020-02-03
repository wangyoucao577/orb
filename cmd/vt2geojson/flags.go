package main

import "flag"

var flags struct {
	mvtSource string
	layer     string
	summary   bool
	x         uint
	y         uint
	z         uint
}

func init() {
	flag.StringVar(&flags.mvtSource, "mvt", "", "Mapbox Vector Tile file path or URL, e.g. 'xxx.mvt' or 'xxx.vector.pbf' or 'https://api.mapbox.com/v4/mapbox.mapbox-streets-v6/9/150/194.mvt?access_token=YOUR_MAPBOX_ACCESS_TOKEN'.")
	flag.StringVar(&flags.layer, "layer", "", "Include only the specified layer of the Mapbox Vector Tile.")
	flag.BoolVar(&flags.summary, "summary", false, "Print layers summary only.")
	flag.UintVar(&flags.x, "x", 0, "Tile x coordinate")
	flag.UintVar(&flags.y, "y", 0, "Tile x coordinate")
	flag.UintVar(&flags.z, "z", 0, "Tile zoom level.")
}
