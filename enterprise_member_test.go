package deploygate

import (
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func Test_ListEnterpriseMember(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/enterprise/list_enterprise_member")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.ListEnterpriseMembers(&ListEnterpriseMembersRequest{
		Enterprise: "test-enterprise",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}

func Test_AddEnterpriseMember(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/enterprise/add_enterprise_member")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	// TODO: Success response is empty and couldn't decode.
	_, err = c.AddEnterpriseMember(&AddEnterpriseMemberRequest{
		Enterprise: "test-enterprise",
		User:       "naoto-fukuda-test",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_RemoveEnterpriseMember(t *testing.T) {
	t.Parallel()

	c, err := NewClient("user_api_token")
	if err != nil {
		t.Fatal(err)
	}

	r, err := recorder.New("fixtures/enterprise/remove_enterprise_member")
	if err != nil {
		t.Fatal(err)
	}

	c.httpClient.Transport = r
	defer r.Stop()

	resp, err := c.RemoveEnterpriseMember(&RemoveEnterpriseMemberRequest{
		Enterprise: "test-enterprise",
		User:       "naoto-fukuda-test",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error("response caused error")
	}
}
