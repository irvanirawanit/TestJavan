package repository

import (
	"TestJavan/config"
)

type Assets struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type AssetsRepository struct{}

type IAssetsRepository interface {
	AssetsList() (data []Assets, err error)
	AssetsInsert(name string, price int) (err error)
}

func NewAssetsRepository() IAssetsRepository {
	return AssetsRepository{}
}

func (u AssetsRepository) AssetsList() (data []Assets, err error) {
	rows, err := config.DB.Query("SELECT id, name, price FROM assets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var assets Assets
		if err := rows.Scan(&assets.Id, &assets.Name, &assets.Price); err != nil {
			return nil, err
		}
		data = append(data, assets)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func (u AssetsRepository) AssetsInsert(name string, price int) (err error) {
	_, err = config.DB.Exec("INSERT INTO assets (name, price) VALUES (?, ?)", name, price)
	if err != nil {
		return err
	}
	return nil
}
