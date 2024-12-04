package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"strconv"
	"time"
)

type Review struct {
	//为什么64位
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Score       uint8     `json:"score"`
	Content     string    `json:"content"`
	Tags        []Tag     `json:"tags"`
	Status      int       `json:"status"`
	PublishTime time.Time `json:"publish_time"`
}

type Tag struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
}

var (
	cli *elasticsearch.Client
	err error
)

func InitES() {
	cfg := elasticsearch.Config{
		Addresses: []string{"127.0.0.1:9200"},
		Username:  "RAIKI_SAKURA",
	}

	cli, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("ES init success")
}

func CreateIndex(client *elasticsearch.TypedClient) {
	res, err := client.Indices.
		Create("persons").
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("crated index success")
	fmt.Println(res.Index) //index?
}

func CreateDoc(client *elasticsearch.TypedClient) {
	d1 := Review{
		ID:      1,
		UserID:  114514,
		Score:   1,
		Content: "哼哼啊啊啊啊啊啊啊",
		Tags: []Tag{
			{Code: 1919, Title: "食雪汉"}, {Code: 810, Title: "好臭哼哼"},
		},
		Status:      1,
		PublishTime: time.Now(),
	}

	res, err := client.Index("person").
		Id(strconv.FormatInt(d1.ID, 10)).
		Document(d1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("created doc success")
	fmt.Println(res.Result) //result?

}

func GetDocBYid(client *elasticsearch.TypedClient, id string) {
	res, err := client.Get("person", id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Source_) //??
}

func GetAddDoc(client *elasticsearch.TypedClient) {
	res, err := client.Search().
		Index("person").
		Query(&types.Query{
			MatchAll: &types.MatchAllQuery{},
		}).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("total: %d \n", res.Hits.Total.Value) //??

	for _, hit := range res.Hits.Hits { //??
		fmt.Println(hit.Source_) //??
	}
}

func GetSpecialDoc(client *elasticsearch.TypedClient, want string) {
	res, err := client.Search().
		Index("person").
		Query(&types.Query{
			MatchPhrase: map[string]types.MatchPhraseQuery{
				"content": {Query: want},
			},
		}).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("total: %d \n", res.Hits.Total.Value) //??

	for _, hit := range res.Hits.Hits { //??
		fmt.Println(hit.Source_) //??
	}
}

// aggregationDemo 聚合
//在 my-review-1 上运行一个平均值聚合，得到所有文档 score 的平均值。
//不理解
//	func aggregationDemo(client *elasticsearch.TypedClient) {
//		avgScoreAgg, err := client.Search().
//			Index("my-review-1").
//			Request(
//				&search.Request{
//					Size: some.Int(0),
//					Aggregations: map[string]types.Aggregations{
//						"avg_score": { // 将所有文档的 score 的平均值聚合为 avg_score
//							Avg: &types.AverageAggregation{
//								Field: some.String("score"),
//							},
//						},
//					},
//				},
//			).Do(context.Background())
//		if err != nil {
//			fmt.Printf("aggregation failed, err:%v\n", err)
//			return
//		}
//		fmt.Printf("avgScore:%#v\n", avgScoreAgg.Aggregations["avg_score"])
//	}

func UpdateDoc_string(client *elasticsearch.TypedClient, id string) {
	new_d1 := Review{
		ID:      1,
		UserID:  114514,
		Score:   5,
		Content: "我超，冰",
		Tags: []Tag{
			{Code: 1919, Title: "食雪汉"},
			{Code: 810, Title: "好臭哼哼"},
			{Code: 233, Title: "www"},
		},
		Status:      1,
		PublishTime: time.Now(),
	}

	res, err := client.Update("person", id).
		Doc(new_d1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("update doc success")
	fmt.Println(res.Result) //??
}

func UpdateDoc_json(client *elasticsearch.TypedClient) {
	new_d1 := `{
					"id":1,
					"userID":147982601,
					"score":5,
					"content":"这是一个二次修改后的好评！",
					"tags":[
						{
							"code":1000,
							"title":"好评"
						},
						{
							"code":9000,
							"title":"有图"
						}
					],
					"status":2,
					"publishDate":"2023-12-10T15:27:18.219385+08:00"
				}`

	res, err := client.Update("person", "1").
		Request(&update.Request{
			Doc: json.RawMessage(new_d1),
		}).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Println("update doc success")
	fmt.Println(res.Result) //??
}

func DeleteDoc(client elasticsearch.TypedClient) {
	res, _ := client.Delete("person", "1").
		Do(context.Background())
	fmt.Println(res.Result) //??
}

func Deleteindex(client elasticsearch.TypedClient) {
	res, _ := client.Indices.Delete("person").
		Do(context.Background())
	fmt.Println(res.Acknowledged) //??
}

func main() {

}
