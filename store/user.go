package store

import "github.com/shahinrahimi/teleprompt/models"

func (s *SqliteStore) GetUser(user_id int64) (*models.User, error) {
	var user models.User
	if err := s.db.QueryRow(models.SELECT_USER_BY_USER_ID, user_id).Scan(user.ToFields()...); err != nil {
		s.l.Printf("error scanning row for a user: %v", err)
		return nil, err
	}
	return &user, nil
}

func (s *SqliteStore) GetUsers() ([]models.User, error) {
	rows, err := s.db.Query(models.SELECT_USERS)
	if err != nil {
		return nil, err
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(user.ToFields()...); err != nil {
			s.l.Printf("error scanning row for a user: %v", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *SqliteStore) CreateUser(u *models.User) error {
	if _, err := s.db.Exec(models.INSERT_USER, u.ToArgs()...); err != nil {
		s.l.Printf("error inserting a new user to DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) DeleteUser(user_id int64) error {
	// TODO delete prompts before deleting the user
	if _, err := s.db.Exec(models.DELETE_USER_BY_USER_ID, user_id); err != nil {
		s.l.Printf("error deleting a user from DB: %v", err)
		return err
	}
	return nil
}
