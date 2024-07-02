package cmd

import (
	"encoding/json"
	"os"

	"github.com/paulmach/orb/geojson"
)

func Merge(filenames []string) (*geojson.FeatureCollection, error) {
	fc := geojson.NewFeatureCollection()

	for _, filename := range filenames {
		b, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		tempFc := geojson.NewFeatureCollection()
		err = json.Unmarshal(b, &tempFc)
		if err != nil {
			return nil, err
		}

		fc.Features = append(fc.Features, tempFc.Features...)
	}

	return fc, nil
}
