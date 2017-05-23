package deployer

type directory struct {
	director Director
}

func (d *directory) Deployments() ([]Deployment, error) {
	deps, err := d.director.Deployments()
	if err != nil {
		return nil, err
	}

	all := make([]Deployment, len(deps))
	for i, dep := range deps {
		all[i] = &deployment{dep}
	}

	return all, nil
}

func (d *directory) Deployment(name string) (Deployment, error) {
	dep, err := d.director.FindDeployment(name)
	if err != nil {
		return nil, err
	}

	return &deployment{dep}, nil
}
