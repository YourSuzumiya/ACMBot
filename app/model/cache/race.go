package cache

import (
	"encoding/json"
	"github.com/YourSuzumiya/ACMBot/app/model"
	"time"
)

func keyRace(source string) string {
	return "race:" + source
}

func SetRace(source string, data []model.Race, exp time.Duration) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return rdb.Set(ctx, keyRace(source), res, exp).Err()
}

func GetRace(source string) ([]model.Race, error) {
	races, err := rdb.Get(ctx, keyRace(source)).Result()
	if err != nil {
		return nil, err
	}
	var result []model.Race
	err = json.Unmarshal([]byte(races), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}