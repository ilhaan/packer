package main

import (
	"fmt"

	"github.com/hashicorp/packer/packer"
)

type mapOfProvisioner map[string]packer.Provisioner

func (mop mapOfProvisioner) Get(provisioner string) (packer.Provisioner, error) {
	p, found := mop[provisioner]
	var err error
	if !found {
		err = fmt.Errorf("Unknown provisioner %s", provisioner)
	}
	return p, err
}

func (mod mapOfProvisioner) List() []string {
	res := []string{}
	for k := range mod {
		res = append(res, k)
	}
	return res
}

type mapOfPostProcessor map[string]packer.PostProcessor

func (mop mapOfPostProcessor) Get(provisioner string) (packer.PostProcessor, error) {
	p, found := mop[provisioner]
	var err error
	if !found {
		err = fmt.Errorf("Unknown post-processor %s", provisioner)
	}
	return p, err
}

func (mod mapOfPostProcessor) List() []string {
	res := []string{}
	for k := range mod {
		res = append(res, k)
	}
	return res
}

type mapOfBuilder map[string]packer.Builder

func (mob mapOfBuilder) Get(builder string) (packer.Builder, error) {
	d, found := mob[builder]
	var err error
	if !found {
		err = fmt.Errorf("Unknown builder %s", builder)
	}
	return d, err
}

type mapOfCommunicator map[string]packer.ConfigurableCommunicator

func (mob mapOfCommunicator) Get(communicator string) (packer.ConfigurableCommunicator, error) {
	c, found := mob[communicator]
	var err error
	if !found {
		err = fmt.Errorf("Unknown communicator %s", communicator)
	}
	return c, err
}
