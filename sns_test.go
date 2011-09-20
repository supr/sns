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

func (s *S) TestDeleteTopic(c *gocheck.C) {
    testServer.PrepareResponse(200, nil, TestDeleteTopicXmlOK)

    t := sns.Topic{nil, "arn:aws:sns:us-east-1:123456789012:My-Topic"}
    resp, err := s.sns.DeleteTopic(t)
    req := testServer.WaitRequest()

    c.Assert(req.Method, gocheck.Equals, "GET")
    c.Assert(req.URL.Path, gocheck.Equals, "/")
    c.Assert(req.Header["Date"], gocheck.Not(gocheck.Equals), "")

    c.Assert(resp.ResponseMetadata.RequestId, gocheck.Equals, "f3aa9ac9-3c3d-11df-8235-9dab105e9c32")
    c.Assert(err, gocheck.IsNil)
}

func (s *S) TestListSubscriptions(c *gocheck.C) {
    testServer.PrepareResponse(200, nil, TestListSubscriptionsXmlOK)

    resp, err := s.sns.ListSubscriptions(nil)
    req := testServer.WaitRequest()

    c.Assert(req.Method, gocheck.Equals, "GET")
    c.Assert(req.URL.Path, gocheck.Equals, "/")
    c.Assert(req.Header["Date"], gocheck.Not(gocheck.Equals), "")

    c.Assert(len(resp.Subscriptions), gocheck.Not(gocheck.Equals), 0)
    c.Assert(resp.Subscriptions[0].Protocol, gocheck.Equals, "email")
    c.Assert(resp.Subscriptions[0].Endpoint, gocheck.Equals, "example@amazon.com")
    c.Assert(resp.Subscriptions[0].SubscriptionArn, gocheck.Equals, "arn:aws:sns:us-east-1:123456789012:My-Topic:80289ba6-0fd4-4079-afb4-ce8c8260f0ca")
    c.Assert(resp.Subscriptions[0].TopicArn, gocheck.Equals, "arn:aws:sns:us-east-1:698519295917:My-Topic")
    c.Assert(resp.Subscriptions[0].Owner, gocheck.Equals, "123456789012")
    c.Assert(err, gocheck.IsNil)
}
