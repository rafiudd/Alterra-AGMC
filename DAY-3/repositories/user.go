package repositories

import (
	"DAY_3/models"
)

func (r *repositories) CreateUser(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositories) UpdateUser(user *models.User, id int) error {
	if err := r.db.Model(user).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositories) DeleteUser(id int) error {
	var user *models.User
	if err := r.db.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositories) GetUserById(id int) (models.User, error) {
	var user models.User
	result := r.db.Where("id = ?", id).First(&user)
	return user, result.Error
}

func (r *repositories) GetAllUsers(keywords string) ([]models.User, error) {
	var users []models.User
	result := r.db.Where("email LIKE ? OR username LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&users)

	return users, result.Error
}

func (r *repositories) UserLogin(username string) (models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	return user, result.Error
}
