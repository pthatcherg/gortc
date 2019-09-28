package gortc

import "net"

// An IceTransport has an IceGatherer
// An IceGatherer has a NetworkManager
// A NetworkManager can list Networks
// Networks have adapter types
// An IceGatherer gathers in sessions, one per local parameters
// (one per ICE restart/regather, like RFC 8445 defines)
// An IceGatheringSession has many IceGatheringSequences, one per network

// TODO:
// - Consider making an IceGatherer be per ICE restart like
//   (make more than one to regather or restart)

type networkAdapterType int

const (
	unknownNetworkAdapterType networkAdapterType = iota
	wifi                      networkAdapterType = iota
	cellular                  networkAdapterType = iota
	otherNetworkAdapterType   networkAdapterType = iota
)

// like rtc::Network
type network interface {
	// ***
	isIpv6() bool
	adapterType() networkAdapterType
	listenUdp() net.PacketConn
}

// like rtc::NetworkManager and VirtualSocketServer
type networkManager interface {
	networks() []network
	// Like NetworkManager::StartUpdating and SignalNetworksChanged
	getNetworkUpdates() <-chan struct{}
	// ***
	createUdpSocket() interface{}
}

type iceGatheringState int

const (
	gatheringNotStarted iceGatheringState = iota
	gathering           iceGatheringState = iota
	gatheringComplete   iceGatheringState = iota
	gatheringStopped    iceGatheringState = iota
)

// Like "allocation flags" and "candidates filters" at the same time
type iceGatheringConfig struct {
	excludeHostCandidates            bool
	excludeServerReflexiveCandidates bool
	excludeRelayCandidates           bool
	disableIpv6                      bool
	disableIpv6OnWifi                bool
	stunCandidateKeepAliveIntervalMs int
}

func (config iceGatheringConfig) shouldIgnoreNetwork(network network) bool {
	if config.excludeHostCandidates &&
		config.excludeServerReflexiveCandidates &&
		config.excludeRelayCandidates {
		// Excludes everything
		return true
	}
	if config.disableIpv6 && network.isIpv6() {
		return true
	}
	if config.disableIpv6OnWifi && network.isIpv6() && network.adapterType() == wifi {
		return true
	}
	return false
}

// Like PortAllocator
type iceGatherer struct {
	// Add support for pruning turn ports (like PortAllocator::prune_turn_ports())
	networkManager networkManager
	config         iceGatheringConfig
}

// Like PortAllocatorSession
type iceGatheringSession struct {
	gatherer  *iceGatherer
	state     iceGatheringState
	sequences *iceGatheringSequence
}

// Like PortAllocator::CreateSession
func (gatherer *iceGatherer) startSession(localParams iceParameters) *iceGatheringSession {
	session := &iceGatheringSession{
		gatherer: gatherer,
		state:    gathering,
	}

	networks := session.gatherer.networkManager.networks()
	// like AllocationSequence::OnReadPacket
	for _, network := range networks {
		// TODO: Do AllocationSequence::DisableEquivalentPhases thing
		if gatherer.config.shouldIgnoreNetwork(network) {
			continue
		}
		sequence := &iceGatheringSequence{
			state: gathering,
		}
		udp := network.listenUdp()
		hostCandidate := IceCandidate{
			usernameFragment: localParams.usernameFragment,
			network:          network,
			typ:              host,
			addr:             udp.LocalAddr(),
		}
		go sequence.processUDPPackets(udp)
		if !stun_server_addresses_.empty() {
			SendStunBindingRequests()
		}

		// *** give candidate out via a channel somewhere

		// ****
		udpPort := NewUdpPort(gatherer, network, udp, localParams)
		// ****
		port.AddAddress(udp.LocalAddr().String())

		// *** don't use append here or elsewhere
		session.ports = append(session.ports, udpPort)
		session.sequences = append(session.sequences, sequence)
	}
	networkUpdates := gatherer.networkManager.getNetworkUpdates()
	// Like BasicPortAllocatorSession::OnNetworksChanged
	go func() {
		for update := range networkUpdates {
			// *** Handle network updates
		}
	}()
	return session
}

// Like AllocationSequence
// TODO: Pick a better name
type iceGatheringSequence struct {
	state iceGatheringState
}

func (sequence *iceGatheringSequence) processUDPPackets(udp net.PacketConn) {
	p := make([]byte, 1500)
	n, addr, err := udp.ReadFrom(p)
	if err != nil {
		// TODO: Log warning
		break
	}
	// if addr is a TURN server and it's a
	//  treat it as as TURN packet as in TurnPort::HandleIncomingPacket
	//  if it's channel data, convert the channel ID to the remote address
	//   and process the packet as if it came directly
	//  if it's a data indication, read out the address
	//   and process the packet as if it came directly
	//  if it's a STUN binding response
	//   treat it like it came from a STUN server
	// if addr is a stun server
	//  read out the transaction ID
	//  look for a request in the STUN binding request map
	//  if you find one, parse the whole packet
	//  if it's a success, add a server reflexive candidate as in UDPPort::OnStunBindingRequestSucceeded
	//  if it's an error, handle the error as in StunPort::OnErrorReponse
	//  Send another one until the lifetime passes
	// if from a remote candidate we know of
	//  if it's a STUN binding request
	//   if it's good
	//    handle check as in Connection::HandleBindingRequest
	//   if it's bad
	//    send back an error as in Port::SendBindingErrorResponse
	//  if it's a binding response
	//   if it's good
	//    handle a successful response as in Connection::OnConnectionRequestResponse
	//   if it's malformed
	//    ignore it
	//   if it's an error
	//    if it's a role conflict
	//     deal with a role conflict as in IceTransportAdapterImpl::OnRoleConflict
	//    if it's non-recoverable
	//     go to a failed state and delete the candidate pair
	// if we've never seen the remote address before
	//  if it's a stun binding request
	//   handle peer reflexive candidate as in P2PTransportChannel::OnUnknownAddress
	//   handle possible role conflict
}

type iceRole int

const (
	controlling iceRole = iota
	controlled  iceRole = iota
)

type iceConfig struct {
	role       iceRole
	tieBreaker uint64
}

// Like P2PTransportChannel
type iceTransport struct {
	gatherer          *iceGatherer
	gatheringState    iceGatheringState
	gatheringSessions []*iceGatheringSession
}

type iceParameters struct {
	usernameFragment string
	password         string
}

func newIceTransport(gatherer *iceGatherer, config iceConfig) *iceTransport {
	return &iceTransport{
		gatherer:       gatherer,
		gatheringState: gatheringNotStarted,
	}
}

// Like P2PTransportChannel::MaybeStartGathering
func (transport *iceTransport) gather(localParams iceParameters) {
	transport.changeGatheringState(gathering)
	// TODO: Support pooled candidates like PortAllocator::TakePooledSession
	gatheringSession := transport.gatherer.startSession(localParams)
	transport.gatheringSessions = append(transport.gatheringSessions, gatheringSession)
}

func (transport *iceTransport) changeGatheringState(state iceGatheringState) {
	transport.gatheringState = state
}
