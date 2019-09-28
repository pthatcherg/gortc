package gortc

import (
	"testing"
)

func TestGather(t *testing.T) {
	t.Errorf("Nothing implemented.")
}

// From p2p_transport_channel_unittest.cc
// for nat1, nat2, expected1, expected2, timeout := range [
//  [OPEN,                OPEN,                HOST,             HOST,             1000],
//  [OPEN,                FULL_CONE,           HOST,             SERVER_REFLEXIVE, 1000],
//  [OPEN,                ADDR_RESTRICTED,     HOST,             SERVER_REFLEXIVE, 1000],
//  [OPEN,                PORT_RESTRICTED,     HOST,             SERVER_REFLEXIVE, 1000],
//  [OPEN,                SYMMETRIC,           HOST,             PEER_REFLEXIVE,   1000],
//  [OPEN,                DOUBLE_CONE,         HOST,             SERVER_REFLEXIVE, 1000],
//  [OPEN,                SYMMETRIC_THEN_CONE, HOST,             PEER_REFLEXIVE,   1000],

//  [FULL_CONE,           OPEN,                SERVER_REFLEXIVE, HOST,             1000],
//  [FULL_CONE,           FULL_CONE,           SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [FULL_CONE,           ADDR_RESTRICTED,     SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [FULL_CONE,           PORT_RESTRICTED,     SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [FULL_CONE,           SYMMETRIC,           SERVER_REFLEXIVE, PEER_REFLEXIVE,   1000],
//  [FULL_CONE,           DOUBLE_CONE,         SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [FULL_CONE,           SYMMETRIC_THEN_CONE, SERVER_REFLEXIVE, PEER_REFLEXIVE,   1000],

//  [ADDR_RESTRICTED,     OPEN,                SERVER_REFLEXIVE, HOST,             1000],
//  [ADDR_RESTRICTED,     FULL_CONE,           SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [ADDR_RESTRICTED,     ADDR_RESTRICTED,     SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [ADDR_RESTRICTED,     PORT_RESTRICTED,     SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [ADDR_RESTRICTED,     SYMMETRIC,           SERVER_REFLEXIVE, PEER_REFLEXIVE,   1000],
//  [ADDR_RESTRICTED,     DOUBLE_CONE,         SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [ADDR_RESTRICTED,     SYMMETRIC_THEN_CONE, SERVER_REFLEXIVE, PEER_REFLEXIVE,   1000],

//  [PORT_RESTRICTED,     OPEN,                SERVER_REFLEXIVE, HOST,             1000],
//  [PORT_RESTRICTED,     FULL_CONE,           SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [PORT_RESTRICTED,     ADDR_RESTRICTED,     SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [PORT_RESTRCITED,     PORT_RESTRICTED,     SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [PORT_RESTRCITED,     SYMMETRIC,           RELAY,            PEER_REFLEXIVE,   2000],
//  [PORT_RESTRCITED,     DOUBLE_CONE,         SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [PORT_RESTRCITED,     SYMMETRIC_THEN_CONE, RELAY,            PEER_REFLEXIVE,   2000],

//  [SYMMETRIC,           OPEN,                PEER_REFLEXIVE,   HOST,             1000],
//  [SYMMETRIC,           FULL_CONE,           PEER_REFLEXIVE,   SERVER_REFLEXIVE, 1000],
//  [SYMMETRIC,           ADDR_RESTRICTED,     PEER_REFLEXIVE,   SERVER_REFLEXIVE, 1000],
//  [SYMMETRIC,           PORT_RESTRICTED,     PEER_REFLEXIVE,   RELAY,            2000],
//  [SYMMETRIC,           SYMMETRIC,           PEER_REFLEXIVE,   RELAY,            2000],
//  [SYMMETRIC,           DOUBLE_CONE,         PEER_REFLEXIVE,   SERVER_REFLEXIVE, 1000],
//  [SYMMETRIC,           SYMMETRIC_THEN_CONE, PEER_REFLEXIVE,   RELAY,            2000],

//  [DOUBLE_CONE,         OPEN,                SERVER_REFLEXIVE, HOST,             1000],
//  [DOUBLE_CONE,         FULL_CONE,           SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [DOUBLE_CONE,         ADDR_RESTRICTED,     SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [DOUBLE_CONE,         PORT_RESTRICTED,     SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [DOUBLE_CONE,         SYMMETRIC,           SERVER_REFLEXIVE, PEER_REFLEXIVE,   1000],
//  [DOUBLE_CONE,         DOUBLE_CONE,         SERVER_REFLEXIVE, SERVER_REFLEXIVE, 1000],
//  [DOUBLE_CONE,         SYMMETRIC_THEN_CONE, SERVER_REFLEXIVE, PEER_REFLEXIVE,   1000],

//  [SYMMETRIC_THEN_CONE, OPEN,                PEER_REFLEXIVE,   HOST,             1000],
//  [SYMMETRIC_THEN_CONE, FULL_CONE,           PEER_REFLEXIVE,   SERVER_REFLEXIVE, 1000],
//  [SYMMETRIC_THEN_CONE, ADDR_RESTRICTED,     PEER_REFLEXIVE,   SERVER_REFLEXIVE, 1000],
//  [SYMMETRIC_THEN_CONE, PORT_RESTRICTED,     PEER_REFLEXIVE,   RELAY,            2000],
//  [SYMMETRIC_THEN_CONE, SYMMETRIC,           PEER_REFLEXIVE,   RELAY,            2000],
//  [SYMMETRIC_THEN_CONE, DOUBLE_CONE,         PEER_REFLEXIVE,   SERVER_REFLEXIVE, 1000],
//  [SYMMETRIC_THEN_CONE, SYMMETRIC_THEN_CONE, PEER_REFLEXIVE,   RELAY,            2000],
// ]
//   network := VirtualSocketServer()  // aka vss_
//   nat := NATSocketServer(network)  // aka nss_
//   firewall := FirewallSocketServer(nat)  // aka ss_
//   stunAddr := "99.99.99.1:3478"
//   turnAddr := "99.99.99.3:3478"
//   networkManager1 := FakeNetworkManager()
//   allocator1 := CreateBasicPortAllocator({networkManager1, stunAddr, turnAddr)
//    allocator := NewBasicPortAllocator(networkManager1)
//    allocator.Initialize()
//    candidatePoolSize = 0
//    pruneTurnPorts = false
//    alloactor.SetConfiguration(stunAddr, turnAddr, pruneTurnPorts, candidatePoolSize)
//    allocator.setFlags(PORTALLOCATOR_ENABLE_SHARED_SOCKET)
//    allocator.setStepDelay(kMinimumStepDelay)
//   endpoint1 := newEndpoint{
//     networkManager: networkManager1
//     allocator: allocator1
//     role: controlling,
//     tieBreaker: 11111,
//     readyToSend: false,
//   })
//   networkManager2 := FakeNetworkManager()
//   allocator1 := CreateBasicPortAllocator(networkManager2, stunAddr, turnAddr)
//     ... same as above ...
//   endpoint2 := newEndpoint{
//     networkManager: networkManager2
//     allocator: allocator2
//     role: controlled,
//     tieBreaker: 22222,
//     readyToSend: false,
//     // ... roleConflict, saveCandidates, savedCandidates, iceRegatheringCounter, receivedPackets
//   })
//   ConfigureEndpoint(endpoint1, nat1)
//    switch nat1
//     case OPEN
//      endpoint1.networkManager.addInterface("11.11.11.11:0")
//     case FULL_CONE, ADDR_RESTRICTED, PORT_RESTRCITED, SYMMETRIC
//      endpoint1.networkManager.addInterface("192.168.1.11", 0)
//      nat.addTranslator("11.11.11.11:0", "192.168.1.1:0", nat1).addClient("192.168.1.11:0")
//     case DOUBLE_CONE
//      endpoint1.networkManager.addInterface("192.168.10.11")
//      nat.addTranslator("11.11.11.11:0", "192.168.1.1:0", FULL_CONE).addTranslator("192.168.1.11:0", "192.168.10.11:0", FULL_CONE).addClient("192.168.10.11:0")
//     case SYMMETRIC_THEN_CONE
//      endpoint1.networkManager.addInterface("192.168.10.11")
//      nat.addTranslator("11.11.11.11:0", "192.168.1.1:0", SYMMETRIC).addTranslator("192.168.1.11:0", "192.168.10.11:0", FULL_CONE).addClient("192.168.10.11:0")
//   ConfigureEndpoint(endpoint2, nat2)
//    switch nat2
//     case OPEN
//      endpoint2.networkManager.addInterface("22.22.22.22:0")
//     case FULL_CONE, ADDR_RESTRICTED, PORT_RESTRCITED, SYMMETRIC
//      endpoint2.networkManager.addInterface("192.168.2.22", 0)
//      nat.addTranslator("22.22.22.22:0", "192.168.2.2:0", nat2).addClient("192.168.2.22:0")
//     case DOUBLE_CONE
//      endpoint2.networkManager.addInterface("192.168.20.22")
//      nat.addTranslator("22.22.22.22:0", "192.168.2.2:0", FULL_CONE).addTranslator("192.168.2.22:0", "192.168.20.22:0", FULL_CONE).addClient("192.168.20.22:0")
//     case SYMMETRIC_THEN_CONE
//      endpoint2.networkManager.addInterface("192.168.20.22")
//      nat.addTranslator("22.22.22.22:0", "192.168.2.2:0", SYMMETRIC).addTranslator("192.168.2.22:0", "192.168.20.22:0", FULL_CONE).addClient("192.168.20.22:0")

//
//   ConfigureEndpoint(endpoint2, nat2)
//   remoteParametersKnown := true  // Not known if testing ICE check received before signaling
//   Test(expected1, expected2, timeout)
//    fakeClock := createFakeClock();
//    CreateChannels();
//     var default_config IceConfig
//     renomination := false
//     config1 := default_config
//     config2 := default_config
//     CreateChannels(config1, config2, renomination)
//      params1 := IceParams{"UFRAG1", "PASSWORD0000000000000001", renomination}
//      params2 := IceParams{"UFRAG2", "PASSWORD0000000000000002", renomination}
//      channel1 := CreateChannel(1, config1, params1, params2)
//       localParams := params1
//       remoteParams := params2
//       channel := NewP2pTransportChannel(endpoint1.allocator, endpoint1.role, endpoint1.tieBreaker, endpoint.asyncResolverFactory, localParams)
//       ... listen to events of OnReadyToSend, OnCandidateGathered, OnCandidatesRemoved, OnReadPacket, OnRoleConflict, OnNetworkRouteChanged, OnSentPacket
//       if remoteParametersKnown  // Not known if testing ICE check received before signaling
//        channel.setRemoteParams(remoteParams)
//      channel2 := CreateChannel(2, config2, params2, params1)
//      channel1.MaybeStartGathering()
//      channel2.MaybeStartGathering()
//
//    defer DestroyChannels();
//     channel1.destroy()
//     channel2.destroy()
//    wait at most (timeout + 1000)ms with fakeClock for (
//      ep1_ch1() != NULL &&m
//      ep2_ch1() != NULL &&
//      ep1_ch1()->receiving() &&
//      ep1_ch1()->writable() &&
//      ep2_ch1()->receiving() &&
//      ep2_ch1()->writable())
//    if (ep1_ch1()->selected_connection() && ep2_ch1()->selected_connection())
//      wait at most 2000ms with fakeClock for (
//        expected1 == LocalCandidate(channel1).type &&
//        expected2 == RemoteCandidate(channel1).type &&
//        expected2 == LocalCandidate(channel2).type &&
//        expected1 == RemoteCandidate(channel2).type
//      )
//      // ... Do the same check above as an "expect"
//    TestSendRecv(clock);
//    for 10 times:
//     msg := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
//     sent := channel1.SendPacket(msg)
//     check sent == length(msg)
//     received := ReceiveData(channel2)
//     check received == msg
//     sent := channel1.SendPacket(msg)
//     check sent == length(msg)
//     received := channel1.ReceiveData(channel1)
//      get from channel1.receivedPackets somehow
//     check received == msg
