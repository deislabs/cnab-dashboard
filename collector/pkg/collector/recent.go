package collector

import (
	"time"
)

type RecentBundle struct {
	Bundle string `json:"bundle"`
	Action string `json:"action"`
}

func ListRecent(opts Options) ([]RecentBundle, error) {
	now := time.Now()
	cutoff := 7 * 24 // 1 week in hours

	recents := make([]RecentBundle, 0)

	claims, err := ListClaims(opts.ClaimSources)
	if err != nil {
		return nil, err
	}

	for _, claim := range claims {
		if now.Sub(claim.Modified).Hours() < float64(cutoff) || true {
			recents = append(recents, RecentBundle{
				Bundle: claim.Name,
				Action: claim.Result.Action,
			})
		}
	}

	return recents, nil
}
