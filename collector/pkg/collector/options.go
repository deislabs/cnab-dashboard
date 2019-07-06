package collector

import (
	"errors"
	"time"
)

type Options struct {
	ClaimSources string
	Recent       string

	recent time.Duration
}

func (o *Options) Validate() error {
	if o.ClaimSources == "" {
		return errors.New("CLAIM_SOURCES is a required positional argument")
	}

	return o.ValidateRecent()
}

func (o *Options) ValidateRecent() error {
	if o.Recent != "" {
		d, err := time.ParseDuration(o.Recent)
		if err != nil {
			return errors.New("invalid --recent, should be a time duration such as 7d or 3w")
		}
		o.recent = d
	} else {
		o.recent = 7 * 24 * time.Hour
	}

	return nil
}

func (o *Options) GetRecent() time.Duration {
	return o.recent
}
