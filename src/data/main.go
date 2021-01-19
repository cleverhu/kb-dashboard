package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"knowledgeBase/src/dbs"
	"knowledgeBase/src/models/DocGrpModel"
	"knowledgeBase/src/models/DocModel"
	"knowledgeBase/src/models/KbModel"
	"knowledgeBase/src/models/KbUserModel"
	"knowledgeBase/src/utils/myHttp"
)

func main() {
	_, str, _, _ := myHttp.Request("get", "https://www.yuque.com/api/zsearch?p=1&q=linux&scope=&type=book", nil, nil, "", "", 30)
	//ioutil.WriteFile("666.txt",[]byte(str),os.ModePerm)
	//
	//return
	bookId := ""
	index := 0
	var j interface{}
	_ = json.Unmarshal([]byte(str), &j)
	d := j.(map[string]interface{})["data"].(map[string]interface{})["hits"].([]interface{})
	for _, v := range d {
		//fmt.Println(v)
		m := v.(map[string]interface{})
		bookId = m["id"].(string)
		fmt.Println(bookId)
		//return
		if m["_record"].(map[string]interface{})["toc_yml"] != nil {
			kb := KbModel.KbImpl{
				Name:      m["book_name"].(string),
				Desc:      m["description"].(string),
				Kind:      1,
				CreatorId: 10000 + index,
			}
			fmt.Println(kb)

			dbs.Orm.Table("kbs").Save(&kb)
			fmt.Println(kb)

			user := KbUserModel.KbUserImpl{
				KbID:   kb.ID,
				UserID: kb.CreatorId,
			}
			dbs.Orm.Table("kb_users").Save(&user)

			record := m["_record"].(map[string]interface{})["toc_yml"].(string)
			var y []map[string]interface{}
			_ = yaml.Unmarshal([]byte(record), &y)
			//fmt.Println(y)
			for _, v := range y {
				fmt.Println(v)
				if v["type"].(string) == "DOC" {
					dgm := DocGrpModel.DocGrpImpl{
						GroupName: v["title"].(string),
						KbID:      kb.ID,
						CreatorId: kb.CreatorId,
					}

					dbs.Orm.Table("doc_grps").Save(&dgm)
					fmt.Println(dgm)
					fmt.Println(v["url"])
					//https://www.yuque.com/api/docs/ls5hwx?book_id=2058470&include_contributors=true&include_hits=true&include_like=true&include_pager=true&include_suggests=true
					//https://www.yuque.com/api/docs/qeti83?book_id=2058470&include_contributors=true&include_hits=true&include_like=true&include_pager=true&include_suggests=true
					_, resStr, _, _ := myHttp.Request("get", "https://www.yuque.com/api/docs/"+v["url"].(string)+"?book_id="+bookId+"&include_contributors=true&include_hits=true&include_like=true&include_pager=true&include_suggests=true", nil, nil, "", "", 30)
					fmt.Println(resStr)
					//ioutil.WriteFile("777.txt",[]byte(resStr),os.ModePerm)
					var content map[string]interface{}
					_ = json.Unmarshal([]byte(resStr), &content)

					dm := DocModel.DocImpl{
						KbID:      kb.ID,
						Title:     v["title"].(string),
						TitleUrl:  v["url"].(string),
						Content:   content["data"].(map[string]interface{})["content"].(string),
						CreatorId: kb.CreatorId,
						GroupID:   dgm.ID,
					}

					dbs.Orm.Table("docs").Save(&dm)

					//return
				}

			}
			index++

			//return
		}

		//return

	}
}
