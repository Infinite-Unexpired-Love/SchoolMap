package main

import (
	"TGU-MAP/models"
	"TGU-MAP/service"
	"TGU-MAP/utils"
	"os"
)

func main() {
	if err := insertData(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	data, err := service.ListItemClient.FetchData()
	if err != nil {
		panic(err)
	}
	println(string(utils.Marshal(*data)))

}

func insertData() *models.CustomError {
	// 示例数据
	data := []models.ListItem{
		{
			Title:     "软件园",
			Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
			Latitude:  ptrFloat64(39.134213),
			Longitude: ptrFloat64(117.212123),
			Children: []models.ListItem{
				{
					Title:     "软件学院",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134213),
					Longitude: ptrFloat64(117.212123),
					Children: []models.ListItem{
						{
							Title:     "教学楼B栋",
							Desc:      "这里是冬暖夏凉",
							Latitude:  ptrFloat64(39.134213),
							Longitude: ptrFloat64(117.212123),
						},
					},
				},
				{
					Title:     "大软公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.221321),
				},
				{
					Title:     "大软食堂",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.212321),
				},
				{
					Title:     "北苑公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134231),
					Longitude: ptrFloat64(117.221321),
				},
			},
		},
		{
			Title:     "西苑",
			Desc:      "这里是西苑，有众多学生宿舍，学习氛围浓厚，是学生学习、生活的好地方",
			Latitude:  ptrFloat64(39.134213),
			Longitude: ptrFloat64(117.212123),
			Children: []models.ListItem{
				{
					Title:     "软件学院",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134213),
					Longitude: ptrFloat64(117.212123),
				},
				{
					Title:     "大软公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.221321),
				},
				{
					Title:     "大软食堂",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.212321),
				},
				{
					Title:     "北苑公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134231),
					Longitude: ptrFloat64(117.221321),
				},
			},
		},
		{
			Title:     "其他",
			Desc:      "这里是其他地方",
			Latitude:  ptrFloat64(39.134213),
			Longitude: ptrFloat64(117.212123),
			Children: []models.ListItem{
				{
					Title:     "软件学院",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134213),
					Longitude: ptrFloat64(117.212123),
				},
				{
					Title:     "大软公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.221321),
				},
				{
					Title:     "大软食堂",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.212321),
				},
				{
					Title:     "北苑公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134231),
					Longitude: ptrFloat64(117.221321),
				},
			},
		},
	}

	return service.ListItemClient.InsertData(&data)
}

func ptrFloat64(f float64) *float64 {
	return &f
}
