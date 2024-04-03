package metadata

var (
	tracetestSource     = "tracetest.source"
	xK6TracetestVersion = "xk6.tracetest.version"
)

type Metadata map[string]string

func (m Metadata) Merge(other Metadata) Metadata {
	for k, v := range other {
		m[k] = v
	}

	return m
}

func GetMetadata() Metadata {
	// TODO: add more metadata after getting the response from the k6 team
	// https://github.com/grafana/k6/issues/1320#issuecomment-2032734378
	return Metadata{
		tracetestSource:     "xk6-tracetest",
		xK6TracetestVersion: "0.1.8",
	}
}
