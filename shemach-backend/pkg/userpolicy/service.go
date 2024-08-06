package userpolicy

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/platforms/helper"
)

// Service ...
type Service interface {
	IsOwnerOfPost(userID uint, postID uint) bool
	DoesAdminWithEmailExist(email string) bool
}

type Repository interface {
	IsOwnerOfPost(userID uint, postID uint) bool
	DoesThisEmailExist(ctx context.Context) bool
}

type service struct {
	authR Repository
}

// NewService ...
func NewService(r Repository) Service {
	return &service{r}
}

// CanAddPost ...
func (s *service) IsOwnerOfPost(userID uint, postID uint) bool {
	if userID == 0 || postID == 0 {
		return false
	}
	return s.authR.IsOwnerOfPost(userID, postID)
}
func (s *service) DoesAdminWithEmailExist(email string) bool {
	if email == "" || !(helper.MatchesPattern(email, helper.EmailRX)) {
		return false
	}
	return false
}
