package sns_test

import (
        "launchpad.net/gocheck"
        "launchpad.net/goamz/aws"
        "launchpad.net/goamz/sns"
)

var _ = gocheck.Suite(&S{})

type S struct {
    HTTPSuite
    sns *sns.SNS
}

func (s *S) SetUpSuite(c *gocheck.C) {
    s.HTTPSuite.SetUpSuite(c)
    auth := aws.Auth{"abc", "123"}
    s.sns = sns.New(auth, aws.Region{SNSEndpoint: testServer.URL})
}

func (s *S) TestListTopicsOK(c *gocheck.C) {
    testServer.PrepareResponse(200, nil, TestListTopicsXmlOK)

    resp, err := s.sns.ListTopics(nil)
    req := testServer.WaitRequest()

    c.Assert(req.Method, gocheck.Equals, "GET")
    c.Assert(req.URL.Path, gocheck.Equals, "/")
    c.Assert(req.Header["Date"], gocheck.Not(gocheck.Equals), "")

    c.Assert(resp.ResponseMetadata.RequestId, gocheck.Equals, "bd10b26c-e30e-11e0-ba29-93c3aca2f103")
    c.Assert(err, gocheck.IsNil)
}
