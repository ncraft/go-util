[![Build Status](https://travis-ci.com/ncraft/machinery.svg?branch=master)](https://travis-ci.com/ncraft/machinery) [![GoDoc](https://godoc.org/github.com/ncraft/machinery?status.svg)](http://godoc.org/github.com/ncraft/machinery)

# Nextcraft machinery

Tools for logging, configuration, abstractions for execution flow, http basic auth etc.

## Dependent command execution

Can be used, for example, to delete a Kubernetes resource if it exists. It abstracts away the repeating pattern of checking if a resource exists and creating/deleting it, depending on it's existence.

Use an `ExistenceDependentOperation` to conditionally run API calls, e.g. against Kubernetes. In the example bellow a nginx deployment is deleted if it previously existed. Otherwise, no action will be undertaken as a delete operation would fail if the resource does not exist:

```go
func deleteNginxDeploymentIfExists() {
	err := flow.NewOperation(&flow.Options{
		TargetObjectName: "nginx",
		Execute:          deleteDeployment,
		ExecOnExistence:  true,
		ExistenceCheck: &existenceCheck{
			get: getDeployment,
		},
	}).Run()

	if err != nil {
		log.Fatal(err)
	}
}

type existenceCheck struct {
	get func(name string) (flow.NamedObject, error)
}

func (c *existenceCheck) Get(name string) (flow.NamedObject, error) {
	return c.get(name)
}

func (c *existenceCheck) IsNotFoundError(err error) bool {
	statusErr, ok := err.(*errors.StatusError)
	if !ok {
		return false
	}

	return statusErr.Status().Code == http.StatusNotFound
}

func deleteDeployment(name string) error {
	return clientset.AppsV1().Deployments(namespace).Delete(name, deleteOptions())
}

func getDeployment(name string) (flow.NamedObject, error) {
	return clientset.AppsV1().Deployments(namespace).Get(name, v1.GetOptions{})
}
```

## HTTP Basic authentication

Example:
```go
package main

import (
	"fmt"
	"net/http"
	
	httpUtil "github.com/ncraft/machinery/pkg/http"
)

type colors []string

func main() {
	colors := colors{"red", "blue", "green"}

	authHandler := httpUtil.BasicAuthHandler{
		Username:        "user123",
		Password:        "secret",
		Realm:           "Please provide username and password",
		OriginalHandler: http.HandlerFunc(colors.list),
	}

	http.Handle("/colors", authHandler)

	http.ListenAndServe(":8080", nil)
}

func (colors colors) list(w http.ResponseWriter, req *http.Request) {
	for _, c := range colors {
		fmt.Fprintf(w, "%s\n", c)
	}
}
```

## Logging

Example:
```go
package main

import "github.com/ncraft/machinery/pkg/log"

func main() {
	log.SetDebug(true)

	log.Info("running logging example")

	colors := []string{"red", "blue", "green"}

	for _, color := range colors {
		log.Debug("added color %s", color)
	}
}
```

Prints:
```
INFO: 2018/12/07 11:28:23 main.go:8: this is a log example
DEBUG: 2018/12/07 11:28:23 main.go:13: added color one
DEBUG: 2018/12/07 11:28:23 main.go:13: added color two
DEBUG: 2018/12/07 11:28:23 main.go:13: added color three
```
