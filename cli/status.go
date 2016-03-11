package cli

import (
	"fmt"
	"os"

	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"

	"github.com/giantswarm/inago/controller"
)

var (
	statusCmd = &cobra.Command{
		Use:   "status [group-slice] ...",
		Short: "status of a group",
		Long:  "status of a group",
		Run:   statusRun,
	}
)

func statusRun(cmd *cobra.Command, args []string) {
	req, err := createRequest(args)
	if err != nil {
		newLogger.Error(nil, "%#v", maskAny(err))
		os.Exit(1)
	}

	group := dirnameFromSlices(args)
	statusList, err := newController.GetStatus(req)
	if controller.IsUnitNotFound(err) || controller.IsUnitSliceNotFound(err) {
		if req.SliceIDs == nil {
			newLogger.Error(nil, "Failed to find group '%s'.", req.Group)
		} else if len(req.SliceIDs) == 0 {
			newLogger.Error(nil, "Failed to find all slices of group '%s'.", req.Group)
		} else {
			newLogger.Error(nil, "Failed to find %d slices for group '%s': %v.", len(req.SliceIDs), req.Group, req.SliceIDs)
		}
		os.Exit(1)
	} else if err != nil {
		newLogger.Error(nil, "%#v", maskAny(err))
		os.Exit(1)
	}

	data, err := createStatus(group, statusList)
	if err != nil {
		newLogger.Error(nil, "%#v", maskAny(err))
		os.Exit(1)
	}
	fmt.Println(columnize.SimpleFormat(data))
}
