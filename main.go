package main

import (
	"fmt"

	"go-micro.dev/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"luna.gameworks/echo/model" // TODO: transform this into a proper module
)

func main() {
	// TODO: yeah i know, this is not the right way to do it but theres isnt a infra established yet
	dsn := "host=localhost user=postgres password=257800 dbname=echo port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Collective{}, &model.Project{})
	collective := &model.Collective{CollectiveName: "lunalabs"}
	db.Create(collective)
	db.Create(&model.Project{ProjectName: "lunalabs", ProjectType: model.OBJECT_ORIENTED, Collective: *collective})

	var collectiveNew model.Collective
	db.First(&collectiveNew)

	fmt.Println("Hello, World!")
	fmt.Printf("Collective: %s\n", collectiveNew.CollectiveName)

	service := micro.New("luna.gameworks.echo")
	service.Run()

}
