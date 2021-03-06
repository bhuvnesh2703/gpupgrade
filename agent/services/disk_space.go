package services

import (
	"context"

	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/utils/disk"
)

func (s *AgentServer) CheckDiskSpace(ctx context.Context, in *idl.CheckSegmentDiskSpaceRequest) (*idl.CheckDiskSpaceReply, error) {
	failed, err := disk.CheckUsage(disk.Local, in.Request.Ratio, in.Datadirs...)
	if err != nil {
		return nil, err
	}

	return &idl.CheckDiskSpaceReply{Failed: failed}, nil
}
