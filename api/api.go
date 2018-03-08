package api

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Excerpt  string `json:"excerpt"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Status   string `json:"status"`
}

type Medias struct {
	gorm.Model
	UserId    string `json:"userId"`
	MediaId   string `json:"mediaId"`
	ArticleId string `json:"articleId"`
	Type      string `json:"type"`
	Status    string `json:"status"`
}

type Media struct {
	gorm.Model
	Source      string `json:"source"`
	Description string `json:"description"`
	UserId      string `json:"userId"`
	CategoryId  string `json:"categoryId"`
	ArticleId   string `json:"articleId"`
	Type        string `json:"type"`
	Status      string `json:"status"`
}

type Catalog struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	CatalogId   string `json:catalogId`
	Type        string `json:"type"`
	Status      string `json:"status"`
}

type Option struct {
	gorm.Model
	Key       string `json:"key"`
	Value     string `json:"value"`
	CatalogId string `json:"catalogId"`
	Type      string `json:"type"`
	Status    string `json:"status"`
}

type Profile struct {
	gorm.Model
	Birthdate         string `json:"birthdate"`
	Location          string `json:"location"`
	Greeting          string `json:"greeting"`
	Realm             string `json:"realm"`
	EmailVerified     string `json:"emailVerified"`
	VerificationToken string `json:"verificationToken"`
	Type              string `json:"type"`
	Status            string `json:"status"`
}

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   string `json:"status"`
	Type     string `json:"type"`
}

type Comment struct {
	gorm.Model
	Title     string `json:"title"`
	Content   string `json:"content"`
	Parent    string `json:"parent"`
	Karma     string `json:"karma"`
	Author    string `json:"author"`
	UserId    string `json:"userId"`
	ArticleId string `json:"articleId"`
	Type      string `json:"type"`
	Status    string `json:"status"`
}

type ACL struct {
	Model         string `json:"model"`
	Property      string `json:"property"`
	AccessType    string `json:"accessType"`
	Permission    string `json:"permission"`
	PrincipalType string `json:"principalType"`
	PrincipalId   string `json:"princiaplId"`
}
