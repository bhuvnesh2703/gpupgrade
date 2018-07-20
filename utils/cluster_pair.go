package utils

import (
	"encoding/json"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/greenplum-db/gp-common-go-libs/cluster"
)

type ClusterPair struct {
	OldCluster *cluster.Cluster
	NewCluster *cluster.Cluster
	OldBinDir  string
	NewBinDir  string
}

func (cp *ClusterPair) Load(stateDir string) error {
	var err error

	err = cp.ReadOldConfig(stateDir)
	if err != nil {
		return errors.Wrap(err, "couldn't read old config file")
	}
	err = cp.ReadNewConfig(stateDir)
	if err != nil {
		return errors.Wrap(err, "couldn't read new config file")
	}

	return nil
}

func (cp *ClusterPair) Commit(stateDir string) error {
	err := cp.WriteOldConfig(stateDir)
	if err != nil {
		return errors.Wrap(err, "couldn't write old config file")
	}

	err = cp.WriteNewConfig(stateDir)
	if err != nil {
		return errors.Wrap(err, "couldn't write new config file")
	}

	return nil
}

func GetConfigFilePath(baseDir string) string {
	return filepath.Join(baseDir, "cluster_config.json")
}

func GetNewConfigFilePath(baseDir string) string {
	return filepath.Join(baseDir, "new_cluster_config.json")
}

/*
 * We need to use an intermediary struct for reading and writing fields not
 * present in cluster.Cluster
 */
type ClusterConfig struct {
	SegConfigs []cluster.SegConfig
	BinDir     string
}

func ReadClusterConfig(configFilePath string) (*cluster.Cluster, string, error) {
	contents, err := System.ReadFile(configFilePath)
	if err != nil {
		return nil, "", err
	}
	clusterConfig := &ClusterConfig{}
	err = json.Unmarshal([]byte(contents), clusterConfig)
	if err != nil {
		return nil, "", err
	}
	return cluster.NewCluster(clusterConfig.SegConfigs), clusterConfig.BinDir, nil
}

func WriteClusterConfig(configFilePath string, c *cluster.Cluster, binDir string) error {
	segConfigs := make([]cluster.SegConfig, 0)
	clusterConfig := &ClusterConfig{BinDir: binDir}

	for _, contentID := range c.ContentIDs {
		segConfigs = append(segConfigs, c.Segments[contentID])
	}

	clusterConfig.SegConfigs = segConfigs

	return WriteJSONFile(configFilePath, clusterConfig)
}

func (cp *ClusterPair) ReadOldConfig(baseDir string) error {
	var err error
	cp.OldCluster, cp.OldBinDir, err = ReadClusterConfig(GetConfigFilePath(baseDir))
	return err
}

func (cp *ClusterPair) ReadNewConfig(baseDir string) error {
	var err error
	cp.NewCluster, cp.NewBinDir, err = ReadClusterConfig(GetNewConfigFilePath(baseDir))
	return err
}

func (cp *ClusterPair) WriteOldConfig(baseDir string) error {
	return WriteClusterConfig(GetConfigFilePath(baseDir), cp.OldCluster, cp.OldBinDir)
}

func (cp *ClusterPair) WriteNewConfig(baseDir string) error {
	return WriteClusterConfig(GetNewConfigFilePath(baseDir), cp.NewCluster, cp.NewBinDir)
}

func (cp *ClusterPair) GetPortsAndDataDirForReconfiguration() (int, int, string) {
	return cp.OldCluster.GetPortForContent(-1), cp.NewCluster.GetPortForContent(-1), cp.NewCluster.GetDirForContent(-1)
}

func (cp *ClusterPair) GetMasterPorts() (int, int) {
	return cp.OldCluster.GetPortForContent(-1), cp.NewCluster.GetPortForContent(-1)
}

func (cp *ClusterPair) GetMasterDataDirs() (string, string) {
	return cp.OldCluster.GetDirForContent(-1), cp.NewCluster.GetDirForContent(-1)
}

func (cp *ClusterPair) GetHostnames() []string {
	hostnameMap := make(map[string]bool, 0)
	for _, seg := range cp.OldCluster.Segments {
		hostnameMap[seg.Hostname] = true
	}
	hostnames := make([]string, 0)
	for host := range hostnameMap {
		hostnames = append(hostnames, host)
	}
	return hostnames
}
