package main

import (
	"encoding/json"
	"time"
)


type FormAggregate struct {
	FormId string   `json:"formId"`
	FieldId string  `json:"fieldId"`
	FormName string  `json:"formName"`
	FieldName *string `json:"fieldName"`
	DependsOn []string `json:"dependsOn"`
	DeleteAt time.Time `json:"deleteAt"`
}

func main() {
  // conn, _ :=	db.GetDB()
 
	// DependsOn, err := json.Marshal([]string{"1", "2"})

	// if err != nil {
	// 	panic(err)
	// }

	// str := string(DependsOn)
	// conn.Create(&model.FormsAggregate{
	// 	FormID: "12",
	// 	FieldID: "13",
	// 	FormName: "formName1111",
	// 	FieldName: 	"FieldName",
	// 	DependsOn: &str,
	// })

	
	data := make(map[string]string)

	data["formId"] = "12"
	data["fieldId"] = "13"

	res, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	println(string(res))
}