package main

import (
	"fmt"
	"goweb/pkg/house"
)

func main() {
	driver, err := house.New(house.HzGovResource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 金雅苑，六塘公寓，景城花园，吉兴公寓，野风海天城，梦琴湾，湖景居
	list, err1 := driver.GetList("野风海天城", "1")
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	if list != nil {
		for _, v := range list.List {
			fmt.Println(
				v.UploadDate,
				fmt.Sprintf("%.2f", v.SellPrice/v.HouseArea)+"万/平",
				fmt.Sprintf("%.2f",v.HouseArea),
				v.SellPrice,
			)
		}
	}
	//app.Start()
}
