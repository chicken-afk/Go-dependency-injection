package main

import "fmt"

type Credentials struct {
	Username string
	Password string
}

type ThirdPartyPayment interface {
	Auth(username string, password string) string
	Payment() bool
}

type ThirdParty struct {
	Credential Credentials
	Payment    ThirdPartyPayment
}

func NewThirdParty(Credential Credentials, payment ThirdPartyPayment) *ThirdParty {
	return &ThirdParty{
		Credential: Credential,
		Payment:    payment,
	}
}

type ThirdPartyXendit struct {
}

func (tp *ThirdPartyXendit) Auth(username string, password string) string {
	if username == "xendit" && password == "xendit" {
		fmt.Println("Authenticated with Xendit")
	} else {
		fmt.Println("Failed authenticated with Xendit")
	}
	return "Auth"
}

func (tp *ThirdPartyXendit) Payment() bool {
	fmt.Println("Payment to xendit..")
	return true
}

type ThirdPartyMCP struct {
}

func (tp *ThirdPartyMCP) Auth(username string, password string) string {
	if username == "mcp" && password == "mcp" {
		fmt.Println("Authenticated with MCP")
	} else {
		fmt.Println("Failed authenticated with MCP")
	}
	return "Auth"
}

func (tp *ThirdPartyMCP) Payment() bool {
	fmt.Println("Payment to mcp successfully")
	return true
}

func main() {
	credentialXendit := Credentials{
		Username: "xendit",
		Password: "123456",
	}
	tp := NewThirdParty(credentialXendit, &ThirdPartyXendit{})
	tp.Payment.Auth(tp.Credential.Username, tp.Credential.Password)
	tp.Payment.Payment()

	credentialMCP := Credentials{
		Username: "mcp",
		Password: "mcp",
	}
	mcp := NewThirdParty(credentialMCP, &ThirdPartyMCP{})
	mcp.Payment.Auth(mcp.Credential.Username, mcp.Credential.Password)
	mcp.Payment.Payment()
}
