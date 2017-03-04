package model

//Tenant represents a tenant
type Tenant struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

//Idea represents an idea on a tenant board
type Idea struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//User represents an user inside our application
type User struct {
	ID        int64
	Name      string
	Email     string
	Providers []*UserProvider
}

//UserProvider represents the relashionship between an User and an Authentication provide
type UserProvider struct {
	Name string
	UID  string
}
