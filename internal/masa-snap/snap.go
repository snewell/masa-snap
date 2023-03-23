package masasnap

type SnapBind struct {
	Bind     string `yaml:",omitempty"`
	BindFile string `yaml:"bind-file,omitempty"`
}

type SnapInfo struct {
	Name        string
	Version     string
	Summary     string
	Description string
	Layout      map[string]SnapBind `yaml:",omitempty"`
}
