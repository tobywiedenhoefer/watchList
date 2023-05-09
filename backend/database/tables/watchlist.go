package watchlist

import (
	"database/sql"
	"fmt"

	"github.com/tobywiedenhoefer/watchList/database/models"
)

func rowCount(rows *sql.Rows) int {
	incr := 0
	for rows.Next() {
		incr += 1
	}
	return incr
}

func GetAll(db *sql.DB) ([]models.WatchListRow, error) {
	var watchListRows []models.WatchListRow
	rows, err := db.Query("SELECT * FROM watchlist.watchlist")
	if err != nil {
		return nil, fmt.Errorf("watchlist.GetAll: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var wlr models.WatchListRow
		if err := rows.Scan(&wlr.ID, &wlr.Title, &wlr.MediaType, &wlr.Genre, &wlr.StreamingPlatform, &wlr.ShortNote); err != nil {
			return nil, fmt.Errorf("watchlist.GetAll.rows.Next: %v", err)
		}
		watchListRows = append(watchListRows, wlr)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("watchlist.GetAll.rows.Err: %v", err)
	}
	return watchListRows, nil
}

func GetOne(db *sql.DB, reqId int) ([]models.WatchListRow, error) {
	var watchListRows []models.WatchListRow
	rows, err := db.Query("SELECT * FROM watchlist.watchlist WHERE id=? LIMIT 1", reqId)
	if err != nil {
		return nil, fmt.Errorf("watchlist.GetOne: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var wlr models.WatchListRow
		if err := rows.Scan(&wlr.ID, &wlr.Title, &wlr.MediaType, &wlr.Genre, &wlr.StreamingPlatform, &wlr.ShortNote); err != nil {
			return nil, fmt.Errorf("watchlist.GetOne.rows.Next: %v", err)
		}
		watchListRows = append(watchListRows, wlr)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("watchlist.GetOne.rows.Err: %v", err)
	}
	return watchListRows, nil
}

func PostOne(db *sql.DB, wlr models.WatchListRow) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO watchlist (title, mediatype, genre, streamingplatform, shortnote) VALUES (?, ?, ?, ?, ?)",
		wlr.Title, wlr.MediaType, wlr.Genre, wlr.StreamingPlatform, wlr.ShortNote,
	)
	if err != nil {
		return 0, fmt.Errorf("watchlist.PostOne.INSERT: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("watchlist.PostOne: %v", err)
	}
	return id, nil
}

func PutOne(db *sql.DB, wlr models.WatchListRow) (int64, error) {
	sqlStatement := "UPDATE watchlist.watchlist SET mediaType=?, genre=?, streamingPlatform=?, shortNote=? WHERE id=?"
	rows, err := db.Exec(sqlStatement, wlr.MediaType, wlr.Genre, wlr.StreamingPlatform, wlr.ShortNote, wlr.ID)
	if err != nil {
		return 0, fmt.Errorf("watchlist.PutOne.UPDATE.Exec: %v", err)
	}
	count, err := rows.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("watchlist.PutOne.UPDATE.RowsAffected: %v", err)
	}
	return count, nil
}

func Delete(db *sql.DB, reqId *int64) (int64, error) {
	res, err := db.Exec("DELETE FROM watchlist.watchlist WHERE id=?", reqId)
	if err != nil {
		return 0, fmt.Errorf("watchlist.Delete.rows.Exec: %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("watchlist.Delete.rows.RowsAffected: %v", err)
	}
	return count, nil
}
