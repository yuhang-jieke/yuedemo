package inits

import (
	"fmt"

	"github.com/olivere/elastic/v7"
)

var (
	ElasticClient *elastic.Client
	err           error
)

func EsInit() {
	ElasticClient, err = elastic.NewClient(
		elastic.SetURL("http://115.190.57.118:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		// Handle error
		panic("es连接失败")
	}
	fmt.Println("es连接成功")

}
