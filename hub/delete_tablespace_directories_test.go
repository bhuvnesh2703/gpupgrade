//  Copyright (c) 2017-2020 VMware, Inc. or its affiliates
//  SPDX-License-Identifier: Apache-2.0

package hub_test

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/greenplum-db/gp-common-go-libs/dbconn"

	"github.com/greenplum-db/gpupgrade/greenplum"
	"github.com/greenplum-db/gpupgrade/hub"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/idl/mock_idl"
	"github.com/greenplum-db/gpupgrade/step"
	"github.com/greenplum-db/gpupgrade/testutils"
	"github.com/greenplum-db/gpupgrade/testutils/testlog"
	"github.com/greenplum-db/gpupgrade/upgrade"
)

func TestDeleteTablespaceDirectories(t *testing.T) {
	testlog.SetupLogger()

	target := hub.MustCreateCluster(t, []greenplum.SegConfig{
		{DbID: 1, ContentID: -1, Hostname: "master", DataDir: "/data/qddir", Role: greenplum.PrimaryRole},
		{DbID: 6, ContentID: -1, Hostname: "standby", DataDir: "/data/standby", Role: greenplum.MirrorRole},
		{DbID: 2, ContentID: 0, Hostname: "sdw1", DataDir: "/data/dbfast1/seg1", Role: greenplum.PrimaryRole},
		{DbID: 3, ContentID: 0, Hostname: "msdw1", DataDir: "/data/dbfast_mirror1/seg1", Role: greenplum.MirrorRole},
		{DbID: 4, ContentID: 1, Hostname: "sdw2", DataDir: "/data/dbfast2/seg2", Role: greenplum.PrimaryRole},
		{DbID: 5, ContentID: 1, Hostname: "msdw2", DataDir: "/data/dbfast_mirror2/seg2", Role: greenplum.MirrorRole},
	})
	target.Version = dbconn.NewVersion("6.1.0")

	t.Run("deletes tablespace directories only on the master", func(t *testing.T) {
		tsDir1, _, tsLocation1 := testutils.MustMakeTablespaceDir(t, 16386)
		defer testutils.MustRemoveAll(t, tsLocation1)

		tsDir2, _, tsLocation2 := testutils.MustMakeTablespaceDir(t, 16387)
		defer testutils.MustRemoveAll(t, tsLocation2)

		systemTsDir, systemDbIdDir, systemTsLocation := testutils.MustMakeTablespaceDir(t, 1700)
		defer testutils.MustRemoveAll(t, systemTsLocation)

		masterTablespaces := greenplum.SegmentTablespaces{
			16386: {
				Location:    tsLocation1,
				UserDefined: 1,
			},
			16387: {
				Location:    tsLocation2,
				UserDefined: 1,
			},
			1700: {
				Location:    systemTsLocation,
				UserDefined: 0,
			},
		}

		err := hub.DeleteTargetTablespacesOnMaster(step.DevNullStream, target, masterTablespaces, "301908232")
		if err != nil {
			t.Errorf("DeleteTargetTablespacesOnMaster returned error %+v", err)
		}

		// verify user tablespace directories are deleted
		for _, dir := range []string{tsDir1, tsDir2} {
			if upgrade.PathExists(dir) {
				t.Errorf("expected tablespace directory %q to be deleted", dir)
			}

			dbIdDir := filepath.Dir(filepath.Clean(dir))
			if upgrade.PathExists(dbIdDir) {
				t.Errorf("expected parent dbid directory %q to be deleted", dbIdDir)
			}
		}

		// verify system tablespace directories are not deleted
		if !upgrade.PathExists(systemTsDir) {
			t.Errorf("expected system tablespace directory %q to not be deleted", systemTsDir)
		}

		if !upgrade.PathExists(systemDbIdDir) {
			t.Errorf("expected system tablespace parent dbid directory %q to not be deleted", systemDbIdDir)
		}
	})

	t.Run("deletes tablespace directories only on the primaries", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tablespaces := map[int]greenplum.SegmentTablespaces{
			1: {
				16386: {
					Location:    "/tmp/testfs/master/demoDataDir-1/16386",
					UserDefined: 1,
				},
				16387: {
					Location:    "/tmp/testfs/master/demoDataDir-1/16387",
					UserDefined: 1,
				},
				1663: {
					// system tablespace locations do not include the tablespace oid
					Location:    "/data/qddir/demoDataDir-1",
					UserDefined: 0,
				},
			},
			2: {
				16386: {
					Location:    "/tmp/testfs/primary1/dbfast1/16386",
					UserDefined: 1,
				},
				16387: {
					Location:    "/tmp/testfs/primary1/dbfast1/16387",
					UserDefined: 1,
				},
				1663: {
					// system tablespace locations do not include the tablespace oid
					Location:    "/data/dbfast1/seg1",
					UserDefined: 0,
				},
			},
			4: {
				16386: {
					Location:    "/tmp/testfs/primary2/dbfast2/16386",
					UserDefined: 1,
				},
				16387: {
					Location:    "/tmp/testfs/primary2/dbfast2/16387",
					UserDefined: 1,
				},
				1663: {
					// system tablespace locations do not include the tablespace oid
					Location:    "/data/dbfast2/seg2",
					UserDefined: 0,
				},
			},
		}

		sdw1 := mock_idl.NewMockAgentClient(ctrl)
		sdw1.EXPECT().DeleteTablespaceDirectories(
			gomock.Any(),
			equivalentRequest(&idl.DeleteTablespaceRequest{
				Dirs: []string{
					"/tmp/testfs/primary1/dbfast1/16386/2/GPDB_6_301908232",
					"/tmp/testfs/primary1/dbfast1/16387/2/GPDB_6_301908232",
				}}),
		).Return(&idl.DeleteTablespaceReply{}, nil)

		sdw2 := mock_idl.NewMockAgentClient(ctrl)
		sdw2.EXPECT().DeleteTablespaceDirectories(
			gomock.Any(),
			equivalentRequest(&idl.DeleteTablespaceRequest{
				Dirs: []string{
					"/tmp/testfs/primary2/dbfast2/16386/4/GPDB_6_301908232",
					"/tmp/testfs/primary2/dbfast2/16387/4/GPDB_6_301908232",
				}}),
		).Return(&idl.DeleteTablespaceReply{}, nil)

		master := mock_idl.NewMockAgentClient(ctrl)
		standby := mock_idl.NewMockAgentClient(ctrl)

		agentConns := []*hub.Connection{
			{nil, sdw1, "sdw1", nil},
			{nil, sdw2, "sdw2", nil},
			{nil, master, "master", nil},
			{nil, standby, "standby", nil},
		}

		err := hub.DeleteTargetTablespacesOnPrimaries(agentConns, target, tablespaces, "301908232")
		if err != nil {
			t.Errorf("DeleteTargetTablespacesOnPrimaries returned error %+v", err)
		}
	})

	t.Run("errors when failing to delete tablespace directories on the primaries", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sdw1 := mock_idl.NewMockAgentClient(ctrl)
		sdw1.EXPECT().DeleteTablespaceDirectories(
			gomock.Any(),
			gomock.Any(),
		).Return(&idl.DeleteTablespaceReply{}, nil)

		expected := errors.New("permission denied")
		failedClient := mock_idl.NewMockAgentClient(ctrl)
		failedClient.EXPECT().DeleteTablespaceDirectories(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, expected)

		agentConns := []*hub.Connection{
			{nil, sdw1, "sdw1", nil},
			{nil, failedClient, "sdw2", nil},
		}

		err := hub.DeleteTargetTablespacesOnPrimaries(agentConns, target, nil, "")

		if !errors.Is(err, expected) {
			t.Errorf("got error %#v, want %#v", err, expected)
		}
	})

	t.Run("must not error out when target is not yet created", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sdw1 := mock_idl.NewMockAgentClient(ctrl)
		sdw2 := mock_idl.NewMockAgentClient(ctrl)

		agentConns := []*hub.Connection{
			{nil, sdw1, "sdw1", nil},
			{nil, sdw2, "sdw2", nil},
		}

		err := hub.DeleteTargetTablespacesOnPrimaries(agentConns, nil, nil, "")
		if err != nil {
			t.Errorf("unexpected error %#v", err)
		}
	})
}

// equivalentRequest is a Matcher that can handle differences in order between
// two instances of DeleteTablespaceRequest.Dirs
func equivalentRequest(req *idl.DeleteTablespaceRequest) gomock.Matcher {
	return reqMatcher{req}
}

type reqMatcher struct {
	expected *idl.DeleteTablespaceRequest
}

func (r reqMatcher) Matches(x interface{}) bool {
	actual, ok := x.(*idl.DeleteTablespaceRequest)
	if !ok {
		return false
	}

	// The key here is that Datadirs can be in any order. Sort them before
	// comparison.
	sort.Strings(r.expected.Dirs)
	sort.Strings(actual.Dirs)

	return reflect.DeepEqual(r.expected, actual)
}

func (r reqMatcher) String() string {
	return fmt.Sprintf("is equivalent to %v", r.expected)
}
