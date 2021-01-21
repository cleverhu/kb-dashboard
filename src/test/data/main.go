package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"knowledgeBase/src/common"
	"knowledgeBase/src/models/DocGrpModel"
	"knowledgeBase/src/models/DocModel"
	"knowledgeBase/src/models/KbModel"
	"knowledgeBase/src/models/KbUserModel"
	"knowledgeBase/src/test/utils/myHttp"
	"time"
)

func main() {
	for k := 0; k < 5; k++ {
		_, str, _, _ := myHttp.Request("get", "https://www.yuque.com/api/zsearch?p=1&q=linux&scope=&type=book", nil, nil, "", "", 30)
		bookId := ""
		index := 0
		var j interface{}
		_ = json.Unmarshal([]byte(str), &j)
		d := j.(map[string]interface{})["data"].(map[string]interface{})["hits"].([]interface{})
		for _, v := range d {
			m := v.(map[string]interface{})
			bookId = m["id"].(string)
			fmt.Println(bookId)
			if m["_record"].(map[string]interface{})["toc_yml"] != nil {
				kb := KbModel.KbImpl{
					ID:         0,
					Name:       m["book_name"].(string),
					Desc:       m["description"].(string),
					Kind:       1,
					CreatorID:  10000 + index,
					IsPrivate:  "N",
					CreateTime: time.Now(),
				}

				common.Orm.Table("kbs").Save(&kb)
				//fmt.Println(kb)
				//return
				for m := 0; m <= 5; m++ {
					user := KbUserModel.KbUserImpl{
						KbID:     kb.ID,
						UserID:   10000 + m,
						JoinTime: time.Now(),
						CanEdit:  "Y",
					}
					if m == 0 && index > 5 {
						user = KbUserModel.KbUserImpl{
							KbID:     kb.ID,
							UserID:   kb.CreatorID,
							JoinTime: time.Now(),
							CanEdit:  "Y",
						}
					}

					common.Orm.Table("kb_users").Create(&user)

					//return
				}

				record := m["_record"].(map[string]interface{})["toc_yml"].(string)
				var y []map[string]interface{}
				_ = yaml.Unmarshal([]byte(record), &y)
				for _, v := range y {
					if v["type"].(string) == "DOC" {
						dgm := DocGrpModel.DocGrpImpl{
							GroupName: v["title"].(string),
							KbID:      kb.ID,
							CreatorID: kb.CreatorID,
						}

						common.Orm.Table("doc_grps").Save(&dgm)

						_, resStr, _, _ := myHttp.Request("get", "https://www.yuque.com/api/docs/"+v["url"].(string)+"?book_id="+bookId+"&include_contributors=true&include_hits=true&include_like=true&include_pager=true&include_suggests=true", nil, nil, "", "", 30)

						var content map[string]interface{}
						_ = json.Unmarshal([]byte(resStr), &content)

						dm := DocModel.DocImpl{
							KbID:      kb.ID,
							Title:     v["title"].(string),
							TitleUrl:  v["url"].(string),
							Content:   content["data"].(map[string]interface{})["content"].(string),
							CreatorID: kb.CreatorID,
							GroupID:   dgm.ID,
						}

						common.Orm.Table("docs").Save(&dm)
					}

				}

			}
			index++
			//return
		}
	}

}
