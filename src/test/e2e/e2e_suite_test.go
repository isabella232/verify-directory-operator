/* vi: set ts=4 sw=4 noexpandtab : */

/*
 * Copyright contributors to the IBM Security Verify Directory Operator project
 */

package e2e

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*****************************************************************************/

/*
 * Run the e2e tests using the Ginkgo runner.
 */

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	fmt.Fprintf(GinkgoWriter, "Starting Verify Directory Operator suite\n")
	RunSpecs(t, "Verify Directory operator e2e suite")
}

/*****************************************************************************/

