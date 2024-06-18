package main

import (
	"TGU-MAP/utils"
	"TGU-MAP/web"
)

func main() {
	utils.InitLogger()
	if err := web.StartServer(); err != nil {
		//zap.S().Fatal(err.Error())
		utils.Fatal(err)
	}

}
