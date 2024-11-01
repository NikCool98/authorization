package storage

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func NewDB(storagePath string) (*sql.DB, error) {
	//константа для ошибок, для удобства их нахождения
	const fn = "storage.sqlite.NewDB"

	db, err := sql.Open("sqlite", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	//Подготовка sql запроса для создания таблицы
	stmt, err := db.Prepare(`
CREATE TABLE IF NOT EXISTS motivations(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content VARCHAR(128) NOT NULL DEFAULT '',
    author VARCHAR(128) NOT NULL DEFAULT '');
`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	//выполнение sql запроса и создание таблицы
	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	// Проверяем наличие данных в таблице
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM motivations")
	if err = row.Scan(&count); err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	// Если таблица пуста, вставляем данные
	if count == 0 {
		stmt, err := db.Prepare(`
INSERT INTO motivations (content,author) VALUES
    ('А если ты не уверен в себе ничего хорошего никогда не получится. Ведь если ты в себя не веришь, кто же поверит?', 'Кто-то умный'),
	('Без идеи не может быть ничего великого! Без великого не может быть ничего прекрасного.', 'Гюстав Флобер'),
	('Быстрее всего учишься в трех случаях — до 7 лет, на тренингах, и когда жизнь загнала тебя в угол.', 'Стивен Кови'),
	('В вашем подсознании скрыта сила, способная перевернуть мир.', 'Уильям Джеймс'),
	('В моем словаре нет слова «невозможно».', 'Наполеон Бонапарт'),
	('Важно верить, что талант нам даётся не просто так – и что любой ценой его нужно для чего-то использовать.', 'Мари Кюри'),
	('Ваше время ограничено, не тратьте его, живя чужой жизнью', 'Стив Джобс'),
	('Велики те, кто видит, что миром правят мысли.', 'Ральф Эмерсон');                            
`)
		_, err = stmt.Exec()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		log.Println("Data inserted successfully")
	} else {
		log.Println("Data already exists.")
	}
	return db, nil
}

func NewStore(db *sql.DB) Storage {
	return Storage{db: db}
}

func (s *Storage) GetRandomMotivation() (string, error) {
	var value string
	query := `SELECT content FROM motivations ORDER BY RANDOM() LIMIT 1;`
	err := s.db.QueryRow(query).Scan(&value)
	if err != nil {
		return "", fmt.Errorf(`{"error":"не указан id задачи"}`)
	}
	return value, nil
}
