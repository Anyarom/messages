package postgres

// метод для отправки значений в таблицу БД
func (pgConnPool *PgxClient) DbInsertMessage(phone string, text string) error {
	_, err := pgConnPool.PgxConnPool.Exec(`INSERT INTO messages(phone, text) VALUES ($1, $2) `, phone, text)
	if err != nil {
		return err
	}
	return nil
}
