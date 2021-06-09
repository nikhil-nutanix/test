deployer.go


Start(uuid) {
	ch <- uuid
}

wait() {
	m.Resume(<- ch)
}

Resume(mcUuid) {
	isStateValidForResume()
	if mcState == INIT {
		err = doMulticlusterPrismElementInitialization()
		if err {
			idfState = PE_ERROR
			return
		}
		validateContainerExists()
		err = doMulticlusterEndpointManagerInitialization()
		if err {
			idfState = ERROR
			return
		}
		idfState = COMPLETE
	}
	if mcState == DELETING {
		err = DoMulticlusterPrismElementDeletion()
		if err {
			idfState = DELETE_ERROR
		}
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
	if mcState == DELETING {
		return
	}
	if mcState == COMPLETE {
		allOtherMcInComplete()
		allMcHealthy()
	}
	// If backend exists, delete it.
	if mcState == COMPLETE || (mcState == ERROR && backendExists) {
		err = DoMulticlusterEndpointManagerDelete
		if err {
			return err
		}
		return
	}
	// Backend doesn't exists so start manageability workflow.
	if mcState == {PE_ERROR, DELETE_ERROR, ERROR} {
		!backendExists()
		idfState = COMPLETE
		go deployer.Start(uuid)
		// Or if we want synchronous, we can instead do: 
		// err = DoMulticlusterPrismElementDeletion()
		// if err {
		// 	   idfState = DELETE_ERROR
		// }
		return
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
		for all mc {
			if mc state == DELETING {
				go deployer.Start(uuid)
			}
		}
	}
}

initMulticlusterDeployer() {
	go wait()
}
