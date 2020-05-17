package models

type MediaDS struct {
	Id 	 interface{} `bson:"_id" json:"id"`
	Filename string `bson:"file_name" json:"file_name"`
	Filepath string `bson:"file_path" json:"file_path"`
	Mime string `bson:"mime" json:"mime"`
}

type MediaLink struct {
	Url 	 string `json:"url"`
}

type MediaFilter struct {
	OnlyLink bool
	Query interface{}
}