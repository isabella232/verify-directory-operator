# Introduction

This document contains the release process which should be followed when generating a new release of the IBM Security Verify Directory operator.

## Version Number

The version number should be of the format: `v<year>.<month>.0`, for example: `v22.2.0`.  There should be no leading zero's in any part of the version number. For exampe, v22.02.0 should be v22.2.0.


# Generating a GitHub Release

In order to generate a new version of the operator a new GitHub release should be created: [https://github.com/IBM-Security/verify-directory-operator/releases/new](https://github.com/IBM-Security/verify-directory-operator/releases/new). 

The fields for the release should be:

|Field|Description
|-----|----------- 
|Tag | The version number, e.g. `v21.2.0`
|Release title | The version number, e.g. `v21.2.0`
|Release description | The resources associated with the \<version\-number> IBM Security Verify Directory operator release.

After the release has been created the GitHub actions workflow ([https://github.com/IBM-Security/verify-directory-operator/actions/workflows/build.yml](https://github.com/IBM-Security/verify-directory-operator/actions/workflows/build.yml)) will be executed to generate the build.  This build process will include:

* publishing the generated docker images to the IBM Container Registry;
* adding the manifest zip and bundle.yaml files to the release artifacts in GitHub.

# Publishing

Once a new GitHub release has been generated the updated operator bundle needs to be published so that it can appear on OperatorHub.io, OpenShift Container Platform and OKD.  Detailed information on how to do this can be found at the following URL: [https://k8s-operatorhub.github.io/community-operators/](hhttps://k8s-operatorhub.github.io/community-operators/).

At a high level you need to (taken from: [https://k8s-operatorhub.github.io/community-operators/contributing-via-pr/]()):

1. Test the operator locally.
2. Add the operator bundle to the GitHub repository.  OperatorHub.io and OpenShift use the same process but different GitHub repositories to store the supported operators.  The GitHub repository for OperatorHub.io is [https://github.com/k8s-operatorhub/community-operators]()  and the GitHub repository for OpenShift is [https://github.com/redhat-openshift-ecosystem/community-operators-prod]().  For each GitHub repository:
	2. Fork the GitHub project.
	3. Add the operator bundle to the verify-directory-operator sub-directory.
	4. Push a 'signed' commit of the changes.  See [https://k8s-operatorhub.github.io/community-operators/contributing-prerequisites/](https://k8s-operatorhub.github.io/community-operators/contributing-prerequisites/).  The easiest way to sign the commit is to use the `git commit -s -m '<description>'` command to commit the changes. **NB:** The updates to the `community-operators` and `upstream-community-operators` directories must be in different pull requests.
	5. Contribute the changes back to the main GitHub repository (using the 'Contribute' button in the GitHub console).  This will have the effect of creating a new pull request against the main GitHub repository.
	6. Monitor the 'checks' against the pull request to ensure that all of the automated test cases pass.
	7. Wait for the pull request to be merged.  This will usually happen overnight.

