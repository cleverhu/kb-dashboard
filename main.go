package main


import (
	"fmt"
	"knowledgeBase/src/models/DocGrpModel"
	"knowledgeBase/src/models/DocModel"
	"knowledgeBase/src/models/KbModel"
	"knowledgeBase/src/models/KbUserModel"
	"time"
)

func main() {
	dpm := DocGrpModel.New().
		Mutate(DocGrpModel.WithCreateTime(time.Now())).
		Mutate(DocGrpModel.WithKbID(123))
	fmt.Println(dpm)

	dm:=DocModel.New().
		Mutate(DocModel.WithDocID(1),DocModel.WithGroupID(123))
	fmt.Println(dm)


	kbm:=KbModel.New().
		Mutate(KbModel.WithKbID(1))
	fmt.Println(kbm)

	kbum:=KbUserModel.New(KbUserModel.WithCanEdit("Y")).
		Mutate(KbUserModel.WithKbID(1))
	fmt.Println(kbum)

}