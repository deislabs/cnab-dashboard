package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deislabs/cnab-dashboard/collector/pkg/collector"
)

func Run(opts collector.Options) error {
	wireup := func(pattern string, getData func(collector.Options) (interface{}, error)) {
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			data, err := getData(opts)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintln(w, err.Error())
			}

			j, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintln(w, err.Error())
			}

			fmt.Fprintln(w, string(j))
		})
	}

	wireup("/installs", func(opts collector.Options) (interface{}, error) {
		return collector.CollectInstalls(opts)
	})
	wireup("/recent", func(opts collector.Options) (interface{}, error) {
		return collector.ListRecent(opts)
	})
	wireup("/status", func(opts collector.Options) (interface{}, error) {
		return collector.GetStatus(opts)
	})

	return http.ListenAndServe(":3031", nil)
}
