package models

import "time"

type Image struct {
	AltText string `json:"altText" bson:"altText"`
	ID      string `json:"id" bson:"_id"`
	Class	string `json:"_class" bson:"_class"`
	Created time.Time `json:"_created" bson:"_created"`
	Caption  string `json:"caption" bson:"caption"`
	EncodingFormat string `json:"encodingFormat" bson:"encodingFormat"`
	Height int `json:"height" bson:"height"`
	Width int `json:"width" bson:"width"`
	DatePublished time.Time `json:"datePublished" bson:"datePublished"`
	Source Source `json:"source" bson:"source"`
	Authors []NameType `json:"author" bson:"author"`
	Headline string `json:"headline" bson:"headline"`
	Publisher NameType `json:"publisher" bson:"publisher"`
	ExternalSource string `json:"externalSource" bson:"externalSource"`
	ExternalId string `json:"externalId" bson:"externalId"`
	Name string `json:"name" bson:"name"`
	Url string `json:"url" bson:"url"`
	Etag string `json:"etag" bson:"etag"`
	DateModified time.Time `json:"dateModified" bson:"dateModified"`
	Type string `json:"type" bson:"type"`
	MetaData []KeyValue `json:"_metadata" bson:"_metadata"`
	MessageId string `json:"messageId" bson:"messageId"`
	User string `json:"user" bson:"user"`
	GraphicContent bool `json:"graphicContent" bson:"graphicContent"`
}

type Source struct {
	Copyright string `json:"copyright" bson:"copyright"`
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
}

type NameType struct {
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
}

type KeyValue struct {
	Key string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}