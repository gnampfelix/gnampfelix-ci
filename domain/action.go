package domain

import (
    "os/exec"
)

//  An action maps two scripts and one deployment method to a arbitrary number of
//  branches. See RepoConfig for an example.
type Action struct {
    Branches []string
    PreTest string      //  This is the script that should run before the tests.
    Test string         //  The script that triggers the actual test.
    Deploy Deployment
}

//  Runs the action in the specified ci root. First, the PreTest script is executed,
//  then the Test script. Both scripts run in their own shell. Any output of Both
//  scripts and any errors are returned.
func (a Action)Run(ciRoot string) ([]byte, error) {
    cmd := exec.Command("/bin/sh", ciRoot + a.PreTest)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return output, err
    }

    cmd = exec.Command("/bin/sh", ciRoot + a.Test)
    testOutput, err := cmd.CombinedOutput()
    output  = append(output, testOutput...)
    return output, err
}
