package users

import (
	"fmt"
	"time"

	"github.com/diazadriel0/go-course/models"
)

func RegisterUser() {
	user := new(models.User)
	user.AddUser(1, "Adri", time.Now(), true)
	fmt.Println(user)
}