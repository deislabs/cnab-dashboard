package collector

type InstallPoint struct {
	Date             int `json:"x"`
	InstalledBundles int `json:"y"`
}

type Installs struct {
	InstalledBundles int            `json:"installedBundles"`
	InstallHistory   []InstallPoint `json:"installHistory"`
}

func CollectInstalls(opts Options) (Installs, error) {
	return Installs{InstalledBundles: 25,
		InstallHistory: []InstallPoint{
			{1, 1},
			{7, 7},
			{10, 15},
			{15, 23},
			{17, 25},
		},
	}, nil
}
