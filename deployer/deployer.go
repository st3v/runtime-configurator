package deployer

import boshdir "github.com/cloudfoundry/bosh-cli/director"

type Deployment interface {
	Update(bool, Infoer) error
	Name() string
}

type Director interface {
	Deployments() ([]boshdir.Deployment, error)
	FindDeployment(string) (boshdir.Deployment, error)
}

type Infoer interface {
	Info(string, string, ...interface{})
}

type Directory interface {
	Deployments() ([]Deployment, error)
	Deployment(string) (Deployment, error)
}

type deployer struct {
	dir Directory
	log Infoer
	dry bool
}

func New(director Director, dryRun bool, info Infoer) *deployer {
	return &deployer{
		dir: &directory{director},
		dry: dryRun,
		log: info,
	}
}

func (d *deployer) DeployAllBut(skip []string) error {
	d.log.Info("deployer", "Deploying existing deployments ...")
	defer d.log.Info("deployer", "Done deploying existing deployments")

	deps, err := d.dir.Deployments()
	if err != nil {
		return err
	}

	for _, dep := range deps {
		if contains(skip, dep.Name()) {
			// skip deployment
			d.log.Info("deployer", "Skipping deployment %q ...", dep.Name())
			continue
		}

		if err := dep.Update(d.dry, d.log); err != nil {
			// break on first error
			return err
		}
	}

	return nil
}

func (d *deployer) DeployAll() error {
	return d.DeployAllBut([]string{})
}

func (d *deployer) Deploy(name string) error {
	d.log.Info("deployer", "Deploying existing deployment %q ...", name)
	defer d.log.Info("deployer", "Done deploying existing deployment %q", name)

	dep, err := d.dir.Deployment(name)
	if err != nil {
		return err
	}

	return dep.Update(d.dry, d.log)
}

func contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}
