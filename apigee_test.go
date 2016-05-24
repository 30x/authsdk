package authsdk

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Apigee", func() {

	It("Parse Token", func() {

		testToken := `eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJhZGFkYjllMS02MWFmLTQ3YzYtYjdkNi00OTM3NTI4ODA3OTUiLCJzdWIiOiIzODkyNzcyNy05MzJhLTQ3MDYtYWNlYy01NmQzODIwNzdmMzQiLCJzY29wZSI6WyJzY2ltLm1lIiwib3BlbmlkIiwicGFzc3dvcmQud3JpdGUiLCJhcHByb3ZhbHMubWUiLCJvYXV0aC5hcHByb3ZhbHMiXSwiY2xpZW50X2lkIjoiZWRnZWNsaSIsImNpZCI6ImVkZ2VjbGkiLCJhenAiOiJlZGdlY2xpIiwiZ3JhbnRfdHlwZSI6InBhc3N3b3JkIiwidXNlcl9pZCI6IjM4OTI3NzI3LTkzMmEtNDcwNi1hY2VjLTU2ZDM4MjA3N2YzNCIsIm9yaWdpbiI6InVzZXJncmlkIiwidXNlcl9uYW1lIjoidG5pbmVAYXBpZ2VlLmNvbSIsImVtYWlsIjoidG5pbmVAYXBpZ2VlLmNvbSIsImF1dGhfdGltZSI6MTQ2NDEzMDk5MywicmV2X3NpZyI6IjZkZTJjODk2IiwiaWF0IjoxNDY0MTMwOTkzLCJleHAiOjE0NjQxMzI3OTMsImlzcyI6Imh0dHBzOi8vbG9naW4uYXBpZ2VlLmNvbS9vYXV0aC90b2tlbiIsInppZCI6InVhYSIsImF1ZCI6WyJlZGdlY2xpIiwic2NpbSIsIm9wZW5pZCIsInBhc3N3b3JkIiwiYXBwcm92YWxzIiwib2F1dGgiXX0.UF5ObmUpeqFuEv58lBYq5O59wqYRt_mb5a5Sh96n2h4DsSk_g7WQ9gupn2BTZGn2ZVyjlrXVlSR5hDTuvJ-x1dCYh3f4wZASBZwqeRqa2quYiSmrVPPXk_lUbaj379O6HOLSfkaRNyGwfGS_QaS3_odMFyoPFF3rnpa4ZpeyOY66esLHAh0JhLHtE9vBr-0q3wps7GcxOIZUe3Xne-katxX3F4e9YenADDQ7CCvZiznASdDF1mH1QQ1nNA0lCz2HPZhqLc5Axj7TGJbugHT34eNg7mltOQEgWfnXCe_IpMlgac3YHjofa8ww8imtMuqyHNE5XetJ1sCJTLNBWt_ogQ`

		jwtToken, err := NewApigeeJWTToken(testToken)

		Expect(err).Should(BeNil(), "Should not return an error creating a valid token")

		//if could not find directory, it's a fail

		Expect(jwtToken.GetSubject()).Should(Equal("38927727-932a-4706-acec-56d382077f34"))

		Expect(jwtToken.GetEmail()).Should(Equal("tnine@apigee.com"))

		Expect(jwtToken.GetUsername()).Should(Equal("tnine@apigee.com"))

	})

})
