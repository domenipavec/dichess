package bluetooth

import (
	"context"

	"github.com/muka/go-bluetooth/api"
	"github.com/pkg/errors"
)

func StartController(ctx context.Context) error {
	// cmd := exec.CommandContext(ctx, "btattach", "-N", "-S", "115200", "-P", "bcm", "-B", "/dev/ttyAMA0")
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout
	// err := cmd.Start()
	// if err != nil {
	//     return errors.Wrap(err, "could not start btattach")
	// }
	//
	// go func() {
	//     err := cmd.Wait()
	//     log.Printf("btattach exited with: %v", err)
	// }()
	//
	// for i := 0; i < 60; i++ {
	//     time.Sleep(time.Second)
	//     exists, err := api.AdapterExists("hci0")
	//     if err != nil {
	//         return errors.Wrap(err, "could not check for adapter hci0")
	//     }
	//     if exists {
	//         break
	//     }
	// }

	adapter, err := api.GetAdapter("hci0")
	if err != nil {
		return errors.Wrapf(err, "couldn't find adapter %s", "hci0")
	}

	if err := adapter.SetProperty("Powered", true); err != nil {
		return errors.Wrap(err, "couldn't set powered to true")
	}
	if err := adapter.SetProperty("Alias", "dichess"); err != nil {
		return errors.Wrap(err, "couldn't set powered to true")
	}
	if err := adapter.SetProperty("DiscoverableTimeout", uint32(0)); err != nil {
		return errors.Wrap(err, "couldn't set discoverable timeout to 0")
	}
	if err := adapter.SetProperty("Discoverable", true); err != nil {
		return errors.Wrap(err, "couldn't set discoverable to true")
	}

	return nil
}
