package hcl2template

import (
	"testing"
)

func TestParse_build(t *testing.T) {
	defaultParser := getBasicParser()

	tests := []parseTest{
		{"basic build",
			defaultParser,
			parseTestArgs{"testdata/build/basic.pkr.hcl"},
			&PackerConfig{
				Builds: Builds{
					&Build{
						Froms: []SourceRef{
							{
								Type: "amazon-ebs",
								Name: "ubuntu-1604",
							},
							{
								Type: "virtualbox-iso",
								Name: "ubuntu-1204",
							},
						},
						ProvisionerGroups: ProvisionerGroups{
							&ProvisionerGroup{
								CommunicatorRef: CommunicatorRef{
									Type: "ssh",
									Name: "vagrant",
								},
								Provisioners: []Provisioner{
									{Provisioner: basicMockProvisioner},
									{Provisioner: basicMockProvisioner},
								},
							},
							&ProvisionerGroup{
								CommunicatorRef: CommunicatorRef{
									Type: "ssh",
									Name: "secure",
								},
							},
						},
						PostProvisionerGroups: PostProcessorGroups{
							&PostProcessorGroup{
								PostProcessors: []PostProcessor{
									{basicMockPostProcessor},
								},
							},
						},
					},
				}},
			false,
		},
	}
	testParse(t, tests)
}
