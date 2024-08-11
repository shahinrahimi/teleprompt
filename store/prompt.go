package store

import "github.com/shahinrahimi/teleprompt/models"

func (s *SqliteStore) GetPrompt(id int) (*models.Prompt, error) {
	var prompt models.Prompt
	if err := s.db.QueryRow(models.SELECT_PROMPT, id).Scan(prompt.ToFields()...); err != nil {
		s.l.Printf("error scanning row for a prompt: %v", err)
		return nil, err
	}
	return &prompt, nil
}

func (s *SqliteStore) GetPrompts() ([]models.Prompt, error) {
	rows, err := s.db.Query(models.SELECT_PROMPTS)
	if err != nil {
		return nil, err
	}
	var prompts []models.Prompt
	for rows.Next() {
		var prompt models.Prompt
		if err := rows.Scan(prompt.ToFields()...); err != nil {
			s.l.Printf("error scanning row for a prompt: %v", err)
			continue
		}
		prompts = append(prompts, prompt)
	}

	return prompts, nil
}

func (s *SqliteStore) CreatePrompt(p *models.Prompt) error {
	if _, err := s.db.Exec(models.INSERT_PROMPT, p.ToArgs()...); err != nil {
		s.l.Printf("error inserting a new prompt to DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) DeletePrompt(id int) error {
	if _, err := s.db.Exec(models.DELETE_PROMPT, id); err != nil {
		s.l.Printf("error deleting a user from DB: %v", err)
		return err
	}
	return nil
}
