package usecase

// Contains the business logic or application-specific logic.
import "chatroom/internal/domain"

type UserUsecase interface {
	CreateUser(user *domain.User) error
	GetUser(id int) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id int) error
}

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (uc *userUsecase) CreateUser(user *domain.User) error {
	return uc.userRepo.Create(user)
}

func (uc *userUsecase) GetUser(id int) (*domain.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *userUsecase) GetAll() ([]*domain.User, error) {
	return uc.userRepo.GetAll()
}

func (uc *userUsecase) UpdateUser(user *domain.User) error {
	return uc.userRepo.Update(user)
}

func (uc *userUsecase) DeleteUser(id int) error {
	return uc.userRepo.Delete(id)
}
