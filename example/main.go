// Copyright 2017 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/grafeas/client-go"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/name"
	"log"
)

func main() {
	client := swagger.NewGrafeasApi()
	nPID := "best-vuln-scanner"
	nID := "CVE-2014-9911"
	n := note(nPID, nID)
	createdN, _, err := client.CreateNote(nPID, nID, *n)
	if err != nil {
		log.Fatalf("Error creating note %v", err)
	} else {
		log.Printf("Succesfully created note: %v", createdN)
	}

	if got, _, err := client.GetNote(nPID, nID); err != nil {
		log.Fatalf("Error getting note %v", err)

	} else {
		log.Printf("Succesfully got note: %v", got)
	}

	oPID := "scanning-customer"
	o := Occurrence(createdN.Name)
	createdO, _, err := client.CreateOccurrence(oPID, *o)
	if err != nil {
		log.Fatalf("Error creating occurrence %v", err)
	} else {
		log.Printf("Succesfully created occurrence: %v", createdO)
	}

	_, oID, pErr := name.ParseOccurrence(createdO.Name)
	if pErr != nil {
		log.Fatalf("Unable to get occurenceId from occurrence name %v: %v", createdO.Name, err)
	}
	if got, _, err := client.GetOccurrence(oPID, oID); err != nil {
		log.Printf("Error getting occurrence %v", err)
	} else {
		log.Printf("Succesfully got occurrence: %v", got)
	}

}

func note(pID, nID string) *swagger.Note {
	return &swagger.Note{
		Name:             fmt.Sprintf("projects/%v/notes/%v", pID, nID),
		ShortDescription: "CVE-2014-9911",
		LongDescription:  "NIST vectors: AV:N/AC:L/Au:N/C:P/I:P",
		Kind:             "PACKAGE_VULNERABILITY",
		VulnerabilityType: swagger.VulnerabilityType{
			CvssScore: 7.5,
			Severity:  "HIGH",
			Details: []swagger.Detail{
				{
					CpeUri:   "cpe:/o:debian:debian_linux:7",
					Package_: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: swagger.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:7",
						Package_: "icu",
						Version: swagger.Version{
							Name:     "4.8.1.1",
							Revision: "12+deb7u6",
						},
					},
				},
				{
					CpeUri:   "cpe:/o:debian:debian_linux:8",
					Package_: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: swagger.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: swagger.Version{
							Name:     "52.1",
							Revision: "8+deb8u4",
						},
					},
				},
				{
					CpeUri:   "cpe:/o:debian:debian_linux:9",
					Package_: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: swagger.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:9",
						Package_: "icu",
						Version: swagger.Version{
							Name:     "55.1",
							Revision: "3",
						},
					},
				},
				{
					CpeUri:   "cpe:/o:canonical:ubuntu_linux:14.04",
					Package_: "andriod",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: swagger.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "MEDIUM",

					FixedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:canonical:ubuntu_linux:14.04",
						Package_: "andriod",
						Version: swagger.Version{
							Kind: "MAXIMUM",
						},
					},
				},
			},
		},
		RelatedUrl: []swagger.RelatedUrl{
			{
				Url:   "https://security-tracker.debian.org/tracker/CVE-2014-9911",
				Label: "More Info",
			},
			{
				Url:   "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9911",
				Label: "More Info",
			},
		},
	}
}

func Occurrence(noteName string) *swagger.Occurrence {
	return &swagger.Occurrence{
		ResourceUrl: "gcr.io/foo/bar",
		NoteName:    noteName,
		Kind:        "PACKAGE_VULNERABILITY",
		VulnerabilityDetails: swagger.VulnerabilityDetails{
			Severity:  "HIGH",
			CvssScore: 7.5,
			PackageIssue: []swagger.PackageIssue{
				swagger.PackageIssue{
					SeverityName: "HIGH",
					AffectedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: swagger.Version{
							Name:     "52.1",
							Revision: "8+deb8u3",
						},
					},
					FixedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: swagger.Version{
							Name:     "52.1",
							Revision: "8+deb8u4",
						},
					},
				},
			},
		},
	}
}
