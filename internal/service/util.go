package service

import (
	"cufoon.litkeep.service/internal/box"
	"cufoon.litkeep.service/internal/db/entity"
)

func generateMap(kinds []entity.BillKind) map[string]*box.BillKindTree {
	resultMap := map[string]*box.BillKindTree{}
	for _, item := range kinds {
		if item.UpKind == "KindLitWorldRoot" {
			value, ok := resultMap[item.KindID]
			if ok {
				value.KindID = item.KindID
				value.Name = item.Name
				value.Description = item.Description
			} else {
				tmp := new(box.BillKindTree)
				tmp.KindID = item.KindID
				tmp.Name = item.Name
				tmp.Description = item.Description
				resultMap[item.KindID] = tmp
			}
			continue
		}
		if item.UpKind == "" {
			tmp := new(box.BillKindTree)
			tmp.KindID = item.KindID
			tmp.Name = item.Name
			tmp.Description = item.Description
			resultMap[item.KindID] = tmp
			continue
		}
		value, ok := resultMap[item.UpKind]
		if ok {
			value.Children = append(value.Children, box.BillKind{
				KindID:      item.KindID,
				Name:        item.Name,
				Description: item.Description,
			})
		} else {
			tmp := new(box.BillKindTree)
			tmp.Children = append(tmp.Children, box.BillKind{
				KindID:      item.KindID,
				Name:        item.Name,
				Description: item.Description,
			})
			resultMap[item.UpKind] = tmp
		}
	}
	return resultMap
}

func getKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
