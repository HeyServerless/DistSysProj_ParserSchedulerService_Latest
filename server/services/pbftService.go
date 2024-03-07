package services

type Task struct {
	ID         int
	Expression string
}

type PrePrepareMessage struct {
	ViewNumber     int
	SequenceNumber int
	Result         float64
	TaskID         int
}

type PrepareMessage struct {
	ViewNumber     int
	SequenceNumber int
	Result         float64
	TaskID         int
	ReplicaID      int
}

type CommitMessage struct {
	ViewNumber     int
	SequenceNumber int
	Result         float64
	TaskID         int
	ReplicaID      int
}

type ViewChangeMessage struct {
	ViewNumber int
	ReplicaID  int
}

type PBFTNode struct {
	ReplicaID      int
	ViewNumber     int
	PrimaryReplica int
	NumReplicas    int
	Threshold      int
	ReceivedData   map[int][]float64
	Results        chan float64
	Consensus      chan float64
	CommitSequence int
	Commit         map[int]float64
}

// Add the implementations of handlePrePrepare, handlePrepare, handleCommit, and handleViewChange functions

// func TriggerPBFT() {
// 	// Initialize the components
// 	// httpService := /* Initialize the Http service */
// 	// sqsService := /* Initialize the SQS service */
// 	// scheduler := /* Initialize the Scheduler */

// 	// Initialize the RPC servers and PBFT nodes
// 	numReplicas := 4
// 	threshold := numReplicas / 3
// 	pbftNodes := make([]*PBFTNode, numReplicas)
// 	for i := 0; i < numReplicas; i++ {
// 		pbftNodes[i] = &PBFTNode{
// 			ReplicaID:      i,
// 			ViewNumber:     0,
// 			PrimaryReplica: 0,
// 			NumReplicas:    numReplicas,
// 			Threshold:      threshold,
// 			ReceivedData:   make(map[int][]float64),
// 			Results:        make(chan float64),
// 			Consensus:      make(chan float64),
// 		}
// 	}

// 	// Run the RPC servers in parallel
// 	var wg sync.WaitGroup
// 	for i := 0; i < numReplicas; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			// rpcServer := /* Initialize and run the RPC server for replica i */
// 			// rpcServer.ProcessTasks(pbftNodes[i])
// 		}(i)
// 	}

// 	// // Run the components and start processing tasks
// 	// httpService.Run()
// 	// sqsService.Run()
// 	// scheduler.Run()

// 	wg.Wait()
// }

// func (n *PBFTNode) handlePrePrepare(msg *PrePrepareMessage) {
// 	// Validate the message
// 	if msg.ViewNumber != n.ViewNumber || msg.SequenceNumber <= 0 {
// 		fmt.Printf("Invalid PrePrepare message from primary replica: %+v\n", msg)
// 		return
// 	}

// 	// If this node is not the primary replica, broadcast PrepareMessage to all replicas
// 	if n.ReplicaID != n.PrimaryReplica {
// 		prepareMsg := &PrepareMessage{
// 			ViewNumber:     msg.ViewNumber,
// 			SequenceNumber: msg.SequenceNumber,
// 			Result:         msg.Result,
// 			TaskID:         msg.TaskID,
// 			ReplicaID:      n.ReplicaID,
// 		}
// 		n.broadcast(prepareMsg)
// 	}
// }

// func (n *PBFTNode) handlePrepare(msg *PrepareMessage) {

// 	// Validate the message
// 	if msg.ViewNumber != n.ViewNumber || msg.SequenceNumber <= 0 {
// 		fmt.Printf("Invalid Prepare message from replica %d: %+v\n", msg.ReplicaID, msg)
// 		return
// 	}

// 	// Store the received Prepare message
// 	n.ReceivedData[msg.TaskID] = append(n.ReceivedData[msg.TaskID], msg.Result)

// 	// Check for 2f matching PrepareMessages
// 	matchingResults := 0
// 	for _, result := range n.ReceivedData[msg.TaskID] {
// 		if result == msg.Result {
// 			matchingResults++
// 		}
// 	}

// 	if matchingResults >= 2*n.Threshold {
// 		// Broadcast CommitMessage to all replicas
// 		commitMsg := &CommitMessage{
// 			ViewNumber:     msg.ViewNumber,
// 			SequenceNumber: msg.SequenceNumber,
// 			Result:         msg.Result,
// 			TaskID:         msg.TaskID,
// 			ReplicaID:      n.ReplicaID,
// 		}
// 		n.broadcast(commitMsg)
// 	}
// }

// func (n *PBFTNode) handleCommit(msg *CommitMessage) {
// 	// Validate the message
// 	if msg.ViewNumber != n.ViewNumber || msg.SequenceNumber <= 0 {
// 		fmt.Printf("Invalid Commit message from replica %d: %+v\n", msg.ReplicaID, msg)
// 		return
// 	}

// 	// Store the received Commit message
// 	n.CommitData[msg.TaskID] = append(n.CommitData[msg.TaskID], msg.Result)

// 	// Check for 2f+1 matching CommitMessages
// 	matchingResults := 0
// 	for _, result := range n.CommitData[msg.TaskID] {
// 		if result == msg.Result {
// 			matchingResults++
// 		}
// 	}

// 	if matchingResults >= (2*n.Threshold + 1) {
// 		// Send the agreed-upon result to the Consensus channel
// 		n.Consensus <- msg.Result
// 	}
// }

// func (n *PBFTNode) handleViewChange(msg *ViewChangeMessage) {
// 	// Validate the message
// 	if msg.ViewNumber <= n.ViewNumber || msg.ReplicaID < 0 {
// 		fmt.Printf("Invalid ViewChange message from replica %d: %+v\n", msg.ReplicaID, msg)
// 		return
// 	}

// 	// Store the received ViewChange message
// 	n.ViewChangeData[msg.ViewNumber] = append(n.ViewChangeData[msg.ViewNumber], msg.ReplicaID)

// 	// Check for 2f+1 matching ViewChangeMessages
// 	matchingReplicas := len(n.ViewChangeData[msg.ViewNumber])

// 	if matchingReplicas >= (2*n.Threshold + 1) {
// 		// Update the view number and elect a new primary replica
// 		n.ViewNumber = msg.ViewNumber
// 		n.PrimaryReplica = (n.PrimaryReplica + 1) % n.NumReplicas

// 		fmt.Printf("View change to %d. New primary replica: %d\n", n.ViewNumber, n.PrimaryReplica)
// 	}
// }

// func (n *PBFTNode) broadcast(msg interface{}) {
// 	// Broadcast the message to all replicas
// 	for i := 0; i < n.NumReplicas; i++ {
// 		if i != n.ReplicaID {
// 			// Send the message to replica i
// 		}
// 	}
// }

// func (n *PBFTNode) ProcessTasks() {
// 	// Process tasks from the Results channel
// 	for result := range n.Results {
// 		// Process the result
// 	}
// }

// func (n *PBFTNode) ProcessConsensus() {
// 	// Process consensus results from the Consensus channel
// 	for result := range n.Consensus {
// 		// Process the result
// 	}
// }

// func (n *PBFTNode) ProcessViewChange() {
// 	// Process view change messages from the ViewChange channel
// 	for msg := range n.ViewChange {
// 		n.handleViewChange(msg)
// 	}
// }

// func (n *PBFTNode) ProcessPrePrepare() {
// 	// Process PrePrepare messages from the PrePrepare channel
// 	for msg := range n.PrePrepare {
// 		n.handlePrePrepare(msg)
// 	}
// }

// func (n *PBFTNode) CommitData() {
// 	// Process Commit messages from the Commit channel
// 	for msg := range n.Commit {
// 		n.handleCommit(msg)
// 	}
// }
