package geojson

import (
	"encoding/json"
	"os"
	"path/filepath"

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

func Output(path string, outputFilename string, fc *geojson.FeatureCollection) error {
	b, err := json.Marshal(fc)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(filepath.Dir(path), outputFilename), b, 0644)
	if err != nil {
		return err
	}

	return nil
}
