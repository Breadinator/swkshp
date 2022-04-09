package versions

import (
	"github.com/breadinator/swkshp/config"
	"github.com/breadinator/swkshp/resource"
)

// The map maps the game strings to a map mapping IDs to bools, where the boolean represents whether an update is available.
func CheckForUpdate(games ...string) (map[string]map[*resource.Resource]bool, []error) {
	// checks all games
	if len(games) == 0 {
		for game := range config.Conf.Games {
			games = append(games, game)
		}
	}

	resourceUpdates := make(map[string]map[*resource.Resource]bool)
	errs := make([]error, 0)

	for _, game := range games {
		resources, newErrs := checkForUpdateSingle(game)
		resourceUpdates[game] = resources
		errs = append(errs, newErrs[:]...)
	}

	return resourceUpdates, errs
}

func checkForUpdateSingle(game string) (map[*resource.Resource]bool, []error) {
	resources := make(map[*resource.Resource]bool)
	errs := make([]error, 0)

	entries, err := GetAllEntries(game)
	if err != nil {
		errs = append(errs, err)
		return resources, errs
	}

	for _, entry := range entries {
		r := resource.ResourceFromID(entry.ID)
		latest, err := r.Updated()
		if err != nil {
			errs = append(errs, err)
			continue
		}
		resources[&r] = latest.After(entry.Updated)
	}

	return resources, errs
}
