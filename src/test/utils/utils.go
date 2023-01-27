/* vi: set ts=4 sw=4 noexpandtab : */

/*
 * Copyright contributors to the IBM Security Verify Directory Operator project
 */

package utils

/*****************************************************************************/

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo/v2" //nolint:golint,revive
)

/*****************************************************************************/

const (
	certmanagerVersion = "v1.4.0"
	certmanagerURLTmpl = "https://github.com/jetstack/cert-manager/releases/download/%s/cert-manager.yaml"
)

/*****************************************************************************/

/*
 * Diplay a warning message for the supplied error.
 */

func warnError(err error) {
	fmt.Fprintf(GinkgoWriter, "warning: %v\n", err)
}

/*****************************************************************************/

/*
 * Execute the provided command within this context.
 */

func Run(cmd *exec.Cmd) ([]byte, error) {
	dir, _ := GetProjectDir()

	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GO111MODULE=on")

	command := strings.Join(cmd.Args, " ")

	fmt.Fprintf(GinkgoWriter, "running: %s\n", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("%s failed with error: (%v) %s", 
									command, err, string(output))
	}

	return output, nil
}

/*****************************************************************************/

/*
 * Uninstall the cert manager.
 */

func UninstallCertManager() {
	url := fmt.Sprintf(certmanagerURLTmpl, certmanagerVersion)
	cmd := exec.Command("kubectl", "delete", "-f", url)

	if _, err := Run(cmd); err != nil {
		warnError(err)
	}
}

/*****************************************************************************/

/*
 * Install the cert manager bundle.
 */

func InstallCertManager() error {
	url := fmt.Sprintf(certmanagerURLTmpl, certmanagerVersion)
	cmd := exec.Command("kubectl", "apply", "-f", url)

	if _, err := Run(cmd); err != nil {
		return err
	}

	/*
	 * Wait for the cert-manager-webhook to be ready, which can take time if 
	 * the cert-manager was re-installed after uninstalling on a cluster.
	 */

	cmd = exec.Command("kubectl", "wait", 
		"deployment.apps/cert-manager-webhook",
		"--for", "condition=Available",
		"--namespace", "cert-manager",
		"--timeout", "5m",
	)

	_, err := Run(cmd)

	return err
}

/*****************************************************************************/

/*
 * Load a local docker image to the kind cluster.
 */

func LoadImageToKindClusterWithName(name string) error {
	cluster := "kind"

	if v, ok := os.LookupEnv("KIND_CLUSTER"); ok {
		cluster = v
	}

	kindOptions := []string{"load", "docker-image", name, "--name", cluster}

	cmd := exec.Command("kind", kindOptions...)

	_, err := Run(cmd)

	return err
}

/*****************************************************************************/

/*
 * Convert the given command output string into individual objects according 
 * to line breakers, and ignore the empty elements in it.
 */

func GetNonEmptyLines(output string) []string {
	var res []string

	elements := strings.Split(output, "\n")

	for _, element := range elements {
		if element != "" {
			res = append(res, element)
		}
	}

	return res
}

/*****************************************************************************/

/*
 * Return the directory where the project resides.
 */

func GetProjectDir() (string, error) {
	wd, err := os.Getwd()

	if err != nil {
		return wd, err
	}

	wd = strings.Replace(wd, "/test/e2e", "", -1)

	return wd, nil
}

/*****************************************************************************/

/*
 * Replace all instances of old with new in the file at the specified path.
 */

func ReplaceInFile(path, old, new string) error {

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if !strings.Contains(string(b), old) {
		return errors.New("unable to find the content to be replaced")
	}

	s := strings.Replace(string(b), old, new, -1)

	err = os.WriteFile(path, []byte(s), info.Mode())
	if err != nil {
		return err
	}

	return nil
}

/*****************************************************************************/

