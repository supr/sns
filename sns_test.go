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

func (s *S) TestCreateTopic(c *gocheck.C) {
    testServer.PrepareResponse(200, nil, TestCreateTopicXmlOK)

    resp, err := s.sns.CreateTopic("My-Topic")
    req := testServer.WaitRequest()

    c.Assert(req.Method, gocheck.Equals, "GET")
    c.Assert(req.URL.Path, gocheck.Equals, "/")
    c.Assert(req.Header["Date"], gocheck.Not(gocheck.Equals), "")

    c.Assert(resp.Topic.TopicArn, gocheck.Equals, "arn:aws:sns:us-east-1:123456789012:My-Topic")
    c.Assert(resp.ResponseMetadata.RequestId, gocheck.Equals, "a8dec8b3-33a4-11df-8963-01868b7c937a")
    c.Assert(err, gocheck.IsNil)
}
