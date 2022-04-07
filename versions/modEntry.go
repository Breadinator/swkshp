package versions

import (
	"fmt"
	"time"

	"github.com/breadinator/swkshp/utils"
)

const DBName string = ".swkshp.db"

type Entry struct {
	ID      int64     `db:"id"`      // Workshop ID.
	Path    string    `db:"path"`    // Relative path from game mod directory as given by config
	Sum     []byte    `db:"sum"`     // 16 byte MD5 checksum.
	Updated time.Time `db:"updated"` // Timestamp of the last time it was changed.
}

// Returns the newest entry of the two and a boolean of whether or not the checksum is the same (true means it is)
func Compare(a, b *Entry) (*Entry, bool) {
	if a.Updated.After(b.Updated) {
		return a, utils.SlicesEqual(a.Sum, b.Sum)
	} else {
		return b, utils.SlicesEqual(a.Sum, b.Sum)
	}
}

// workshopID: The Steam Workshop ID of the mod.
//
// zippedModPath: The local path to the mod as zipped by Steam.
func NewModEntry(workshopID int, zippedModPath string) (Entry, error) {
	// Fetch workshop timestamp and parse it
	url := fmt.Sprintf(`https://steamcommunity.com/workshop/filedetails/?id=%d`, workshopID)
	t, ok := utils.ParseWorkshopTimestamp(url)
	if !ok {
		return Entry{}, fmt.Errorf("couldn't parse %d's timestamp", workshopID)
	}

	// Get the MD5 checksum of the given zipped file
	md5, err := utils.GetFileMD5(zippedModPath, 512)
	if err != nil {
		return Entry{}, err
	}

	return Entry{
		ID:      int64(workshopID),
		Path:    zippedModPath[:len(zippedModPath)-4],
		Sum:     md5,
		Updated: t,
	}, nil
}
