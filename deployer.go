deployer.go


Start(uuid) {
	ch <- uuid
}

wait() {
	m.Resume(<- ch)
}

Resume(mcUuid) {
	isStateValidForResume()
	if idfState == INIT {
		err = doMulticlusterPrismElementInitialization()
		if err {
			idfState = PE_ERROR
			return
		}
		validateStateIsPeDone
	}
	if idfState == PE_DONE {
		err = doMulticlusterEndpointManagerInitialization()
		if err {
			idfState = ERROR
			return
		}
		idfState = COMPLETE
	}
}

UpdateMulticlusterIdfStateAndErrMsgList() {}

doMulticlusterPrismElementInitialization() {
	ips = getMspIps
	containerExists = StorageContainerExistsByPrismEndpoint()
	if !containerExists {
		GetNumNodesV2()
		GetRedundancyFactorV2()
		GetTargetVersionV2()
		GetDataContainerOptions()
		CreateStorageContainerByPrismEndpoint()
	}
	idfState = PE_DONE
}

DoMulticlusterPrismElementDeletion(proto){
	if containerExists {
		deleteContainer()
	}
	ips = getMspIps
	unwhitelistIps(ips)
	err = deleteIdfEntity()
	if err != nil {
		err = deleteIdfEntity 
		if err != nil && err.contains("not found") {
			log
		} else {
			return err
		}
	}
}

doMulticlusterEndpointManagerInitialization() {}
doMulticlusterEndpointManagerUpdate() {}
doMulticlusterEndpointManagerCancelOp() {}
DoMulticlusterEndpointManagerDelete() {}

DeleteMulticluster() {
	isStateValid()
	isMcPrimary()
	if mcState == {PE_ERROR, DELETING} {
		!backendExists()
		err = DoMulticlusterPrismElementDeletion
		if err {
			idfState = DELETE_ERROR
		}
		return
	}
	// remaining states == {COMPLETE, ERROR, DELETE_ERROR}
	allOtherMcStatesComplete()
	DoMulticlusterEndpointManagerDelete()
	if !McEntityExistsInBackend {
		DoMulticlusterEndpointManagerDelete()
	}
}

maybeDeleteMulticluster() {
	if any mc state == MC_DELETE_STATE_DELETING {
		err = DoMulticlusterPrismElementDeletion()
		if err {
			idfState = DELETE_ERROR
		}
	}
}

startMulticlusterDeleteWatch() {
	every 10 min {
		maybeDeleteMulticluster()
	}
}

initMulticlusterDeployer() {
	go wait()
}
