package main

import agentv0 "github.com/codefly-dev/core/generated/go/codefly/services/agent/v0"

// validationCapabilities describes this agent's workload operations. It requires
// a caller that skips explicitly unsupported test suites; the agent's own Go
// suites run separately in its repository.
//
// image_sbom remains unadvertised: the pinned Core cannot represent that phase
// in effective-input discovery. TestAdvertisedPhasesAreSchedulable guards this
// independently of the Builder.SBOM image-scope RPC.
func validationCapabilities() *agentv0.ValidationCapabilities {
	return &agentv0.ValidationCapabilities{
		Audit:         workspaceOperation(),
		ArtifactBuild: workspaceOperation(),
		Sbom:          workspaceOperation(),
		// There is no derived source to regenerate: everything the build writes
		// into the service tree is ignored by the .gitignore the agent scaffolds
		// and by the one it stages beside the bootstrap tree. Sync therefore
		// reports the same authoritative absence of drift in either mode.
		Sync: workspaceOperation(),
	}
}

// workspaceOperation validates the whole service tree. Postgres owns migrations
// and its service manifest, not language packages or independently checkable
// files, so no narrower scope has a meaning this agent could honour.
func workspaceOperation() *agentv0.ValidationOperationCapability {
	return &agentv0.ValidationOperationCapability{
		Supported: true,
		Scopes:    []agentv0.ValidationScope{agentv0.ValidationScope_VALIDATION_SCOPE_WORKSPACE},
	}
}
