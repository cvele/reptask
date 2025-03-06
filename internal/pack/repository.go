package pack

import (
	"database/sql"
	"errors"

	"github.com/cvele/reptask/internal/db"
)

func GetAllPacks() ([]Pack, error) {
	rows, err := db.DB.Query("SELECT id, size FROM packs ORDER BY size DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packs []Pack
	for rows.Next() {
		var pack Pack
		if err := rows.Scan(&pack.ID, &pack.Size); err != nil {
			return nil, err
		}
		packs = append(packs, pack)
	}

	return packs, nil
}

func GetPackByID(id int) (*Pack, error) {
	var pack Pack
	err := db.DB.QueryRow("SELECT id, size FROM packs WHERE id = ?", id).Scan(&pack.ID, &pack.Size)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &pack, nil
}

func AddPack(size int) error {
	_, err := db.DB.Exec("INSERT INTO packs (size) VALUES (?)", size)
	return err
}

func UpdatePack(id int, size int) error {
	res, err := db.DB.Exec("UPDATE packs SET size = ? WHERE id = ?", size, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func DeletePack(id int) error {
	res, err := db.DB.Exec("DELETE FROM packs WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
