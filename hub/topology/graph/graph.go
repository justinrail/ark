package graph

import (
	"ark/hub/topology/bolt"
	"ark/hub/topology/spout"

	flow "github.com/trustmaster/goflow"
)

//TopologyGraph topology logic net
type TopologyGraph struct {
	flow.Graph
}

//NewGraph A constructor that creates network structure
func NewGraph() *TopologyGraph {
	// Create a new graph
	net := new(TopologyGraph)
	net.InitGraphState()
	// Add graph nodes
	net.Add(new(spout.COGSpout), "cog_spout")
	net.Add(new(bolt.COGAshbin), "cog_ashbin")
	net.Add(new(bolt.COGStateUpdater), "cog_state_updater")

	net.Add(new(spout.CORSpout), "cor_spout")
	net.Add(new(bolt.CORAshbin), "cor_ashbin")
	net.Add(new(bolt.CORStateUpdater), "cor_state_updater")

	net.Add(new(spout.COVSpout), "cov_spout")
	net.Add(new(bolt.COVAshbin), "cov_ashbin")
	net.Add(new(bolt.COVStateUpdater), "cov_state_updater")
	net.Add(new(bolt.COVDataUpdater), "cov_data_updater")
	net.Add(new(bolt.COVStateNormalizer), "cov_state_normalizer")
	net.Add(new(bolt.CoreActiveEventUpdater), "cov_coreactiveevent_updater")
	net.Add(new(bolt.HisCoreLiveEventAppender), "cov_hsicoreactiveevent_appender")
	net.Add(new(bolt.PhoenixCOVHooker), "phoenix_cov_hooker")

	net.Connect("cog_spout", "COGOut1", "cog_state_updater", "In")
	net.Connect("cog_spout", "COGOut2", "cog_ashbin", "In")

	net.Connect("cor_spout", "COROut1", "cor_state_updater", "In")
	net.Connect("cor_spout", "COROut2", "cor_ashbin", "In")

	net.Connect("cov_spout", "COVOut1", "cov_data_updater", "In")
	net.Connect("cov_data_updater", "StateOut", "cov_state_normalizer", "In")
	net.Connect("cov_state_normalizer", "StateOut", "cov_state_updater", "In")
	net.Connect("cov_state_updater", "CoreLiteEventOut1", "cov_coreactiveevent_updater", "In")
	net.Connect("cov_state_updater", "CoreLiteEventOut2", "phoenix_cov_hooker", "In")
	net.Connect("cov_coreactiveevent_updater", "HisCoreLiveEventOut", "cov_hsicoreactiveevent_appender", "In")
	net.Connect("cov_spout", "COVOut2", "cov_ashbin", "In")

	// Network ports
	net.MapInPort("cog_in", "cog_spout", "In")
	net.MapInPort("cor_in", "cor_spout", "In")
	net.MapInPort("cov_in", "cov_spout", "In")

	return net
}
