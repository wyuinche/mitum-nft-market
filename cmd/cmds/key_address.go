package cmds

import (
	mc "github.com/spikeekips/mitum-currency"
	"github.com/spikeekips/mitum/util/logging"
)

type KeyAddressCommand struct {
	printCommand
	Keys []KeyFlag `arg:"" name:"key" help:"key for address (ex: \"<public key>,<weight>\")" sep:"@" optional:""`
}

func (cmd *KeyAddressCommand) Run(log logging.Logger) error {
	keys := make([]mc.Key, len(cmd.Keys))
	for i := range cmd.Keys {
		keys[i] = cmd.Keys[i].Key
	}

	log.Debug().Int("number_of_keys", len(keys)).Interface("keys", keys).Msg("keys loaded")

	if a, err := mc.NewAddressFromKeys(keys); err != nil {
		return err
	} else {
		cmd.print(a.String())
	}

	return nil
}
