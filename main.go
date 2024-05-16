package main

import (
	"TGU-MAP/models"
	"TGU-MAP/service"
	"reflect"
)

func main() {
	ptr := service.ListItemClient.FetchData()
	items := reflect.ValueOf(ptr).Elem().Interface()
	listItems := items.([]models.ListItem)
	for _, item := range listItems {
		println(item.Title)
	}
}
