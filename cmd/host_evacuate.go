package cmd

import (
	"github.com/urfave/cli"
)








func hostEvacuate(ctx *cli.Context) error {
	c, err := GetClient(ctx)
	if err != nil {
		return err
	}

	w, err := NewWaiter(ctx)
	if err != nil {
		return err
	}

	var lastErr error
	for _, id := range ctx.Args() {
		resource, err := Lookup(c, id, "host")
		if err != nil {
			lastErr = printErr(id, lastErr, err)
			continue
		}

		//resourceID, err := fn(c, resource)
		resourceID := resource.Id
		err = c.Action(resource.Type, "evacuate", resource, nil, resource)
		if resourceID == "" {
			resourceID = resource.Id
		}
		lastErr = printErr(resource.Id, lastErr, err)
		if resourceID != "" && resourceID != "-" {
			w.Add(resourceID)
		}
	}

	if lastErr != nil {
		return cli.NewExitError("", 1)
	}

	return w.Wait()
}
