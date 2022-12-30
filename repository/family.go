package repository

import (
	"TestJavan/config"
	"database/sql"
)

type Family struct {
	Id     int           `json:"id"`
	Name   string        `json:"name"`
	Parent sql.NullInt64 `json:"parent"`
	Childs []Family      `json:"childs"`
	Assets []Assets      `json:"assets"`
}

type FamilyRepository struct{}

type IFamilyRepository interface {
	FamilyList() (data []Family, err error)
	Childs(id int) (data []Family, err error)
	AssetsList(id int) (data []Assets, err error)
}

func NewFamilyRepository() IFamilyRepository {
	return FamilyRepository{}
}

func (u FamilyRepository) FamilyList() (data []Family, err error) {
	rows, err := config.DB.Query("SELECT id, name, parent FROM family where parent is null")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var family Family
		if err := rows.Scan(&family.Id, &family.Name, &family.Parent); err != nil {
			return nil, err
		}
		data = append(data, family)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for i := 0; i < len(data); i++ {
		// assets
		data[i].Assets, err = u.AssetsList(data[i].Id)
		if err != nil {
			return nil, err
		}

		// two level
		data[i].Childs, err = u.Childs(data[i].Id)
		if err != nil {
			return nil, err
		}
		// three level
		for j := 0; j < len(data[i].Childs); j++ {
			data[i].Childs[j].Childs, err = u.Childs(data[i].Childs[j].Id)
			if err != nil {
				return nil, err
			}
		}
	}
	return data, nil
}

func (u FamilyRepository) Childs(id int) (data []Family, err error) {
	rows, err := config.DB.Query("SELECT id, name, parent FROM family where parent = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var family Family
		if err := rows.Scan(&family.Id, &family.Name, &family.Parent); err != nil {
			return nil, err
		}
		data = append(data, family)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for i := 0; i < len(data); i++ {
		// assets
		data[i].Assets, err = u.AssetsList(data[i].Id)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func (u FamilyRepository) AssetsList(id int) (data []Assets, err error) {
	// with eager loading assets
	rows, err := config.DB.Query("SELECT a.id, a.name, a.price FROM assets a inner join family_assets fa on a.id = fa.assetid where fa.familyid = ?", id)
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
