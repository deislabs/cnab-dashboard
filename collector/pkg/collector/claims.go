package collector

import (
	"io/ioutil"
	"path/filepath"

	"github.com/deislabs/cnab-go/claim"
	"github.com/deislabs/cnab-go/utils/crud"
	"github.com/pkg/errors"
)

func ListClaims(claimSources string) ([]claim.Claim, error) {
	claims := make([]claim.Claim, 0, 10)

	sources, err := ioutil.ReadDir(claimSources)
	if err != nil {
		return nil, errors.Wrapf(err, "error listing claim sources directory %s", claimSources)
	}

	for _, source := range sources {
		claimDir := filepath.Join(claimSources, source.Name())
		claimfs := crud.NewFileSystemStore(claimDir, "json")
		store := claim.NewClaimStore(claimfs)

		c, err := store.ReadAll()
		if err != nil {
			return nil, errors.Wrapf(err, "error listing claims in %s", claimDir)
		}
		claims = append(claims, c...)
	}

	return claims, nil
}
