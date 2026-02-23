package storage

import (
	"demo/bins/bins"
	"demo/bins/file"
	"encoding/json"
	"errors"
	"fmt"
)

func SaveBin(b bins.Bin, path string) error {
	binList, err := GetBinList(path)
	if err != nil {
		return err
	}

	binList = append(binList, &b)

	data, err := json.MarshalIndent(binList, "", "  ")
	if err != nil {
		return err
	}

	file.WriteFile(data, path)
	if err != nil {
		return err
	}
	return nil
}

func GetBinList(path string) (bins.BinList, error) {
	data, err := file.ReadJSONFile(path)
	if err != nil {
		if errors.Is(err, file.ErrFileNotJSON) {
			fmt.Println("файл не является JSON")
			return nil, err
		}

		return nil, err
	}

	var bl bins.BinList
	err = json.Unmarshal(data, &bl)
	if err != nil {
		return nil, fmt.Errorf("произошла ошибка при unmarshal: %w", err)
	}

	return bl, nil
}
