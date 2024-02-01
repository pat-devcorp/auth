package domain

type User struct{
   UserId   string   `json:"userId"`
   Name     string   `json:"name"`
   Email    string   `json:"email"`
   Password string   `json:"password"`
}

var Users = []User{
   {UserId: "1", Name: "Patrick Carpio", Email: "patricik18483@gmail.com", Password: "BatmanIsBruceWayne"},
   {UserId: "1", Name: "Bruno Carpio", Email: "bruno@gmail.com", Password: "BatmanIsBruceWayne"},
}
