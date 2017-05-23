package deployer

import boshdir "github.com/cloudfoundry/bosh-cli/director"

type deployment struct {
	dep boshdir.Deployment
}

func (d *deployment) Update(dryRun bool, log Infoer) error {
	log.Info("deployement", "Updating %q (dryRun: %t) ...", d.dep.Name(), dryRun)
	defer log.Info("deployement", "Done updating %q", d.dep.Name())

	mf, err := d.dep.Manifest()
	if err != nil {
		return err
	}

	if mf == "" {
		// nothing to update if manifest is empty
		return nil
	}

	return d.dep.Update([]byte(mf), boshdir.UpdateOpts{DryRun: dryRun})
}
