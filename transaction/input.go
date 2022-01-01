package transaction

import "bwa_go/user"

type InputGetTransaction struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
