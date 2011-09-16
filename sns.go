//
// goamz - Go packages to interact with the Amazon Web Services.
//
//   https://wiki.ubuntu.com/goamz
//
// Copyright (c) 2011 Memeo Inc.
//
// Written by Prudhvi Krishna Surapaneni <me@prudhvi.net>
//
package sns

import (
	"launchpad.net/goamz/aws"
	"os"
	"http"
	"time"
	"url"
	"xml"
)

// The SNS type encapsulates operation with an SNS region.
type SNS struct {
	aws.Auth
	aws.Region
	private byte // Reserve the right of using private data.
}

type Topic struct {
	*SNS
	TopicArn string
}

func New(auth aws.Auth, region aws.Region) *SNS {
	return &SNS{auth, region, 0}
}

type Message struct {
	*SNS
	*Topic
	Message [8192]byte
	Subject string
}

type Subscription struct {
    Endpoint string
    Owner string
    Protocol string
    SubscriptionArn string
    TopicArn string
}

func (topic *Topic) Message(message [8192]byte, subject string) *Message {
	return &Message{topic.SNS, topic, message, subject}
}

type ResponseMetadata struct {
	RequestId string
	BoxUsage  float64
}

type ListTopicsResponse struct {
	Topics []Topic `xml:"ListTopicsResult>Topics>member"`
	NextToken string
    ResponseMetadata
}

type CreateTopicResponse struct {
	Topic Topic `xml:"CreateTopicResult"`
	ResponseMetadata
}

type DeleteTopicResponse struct {
	ResponseMetadata
}

type ListSubscriptionsResponse struct {
    Subscriptions []Subscription `xml:"ListSubscriptionsResult>Subscriptions>member"`
    NextToken string
    ResponseMetadata
}

func makeParams(action string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	return params
}

func (sns *SNS) ListTopics(NextToken *string) (resp *ListTopicsResponse, err os.Error) {
	resp = &ListTopicsResponse{}
	params := makeParams("ListTopics")
    if NextToken != nil {
        params["NextToken"] = *NextToken
    }
	err = sns.query(nil, nil, params, resp)
	return
}

func (sns *SNS) CreateTopic(Name string) (resp *CreateTopicResponse, err os.Error) {
	resp = &CreateTopicResponse{}
	params := makeParams("CreateTopic")
    params["Name"] = Name
	err = sns.query(nil, nil, params, resp)
	return
}

func (sns *SNS) DeleteTopic(topic Topic) (resp *DeleteTopicResponse, err os.Error) {
	resp = &DeleteTopicResponse{}
	params := makeParams("DeleteTopic")
	params["TopicArn"] = topic.TopicArn
	err = sns.query(nil, nil, params, resp)
	return
}

func (topic *Topic) Delete() (resp *DeleteTopicResponse, err os.Error) {
	return topic.SNS.DeleteTopic(*topic)
}

func (sns *SNS) ListSubscriptions(NextToken *string) (resp *ListSubscriptionsResponse, err os.Error) {
    resp = &ListSubscriptionsResponse{}
    params := makeParams("ListSubscriptions")
    if NextToken != nil {
        params["NextToken"] = *NextToken
    }
    err = sns.query(nil, nil, params, resp)
    return
}

type Error struct {
	StatusCode int
	Code       string
	Message    string
	RequestId  string
}

func (err *Error) String() string {
	return err.Message
}

type xmlErrors struct {
	RequestId string
	Errors    []Error `xml:"Errors>Error"`
}

func (sns *SNS) query(topic *Topic, message *Message, params map[string]string, resp interface{}) os.Error {
	params["Timestamp"] = time.UTC().Format(time.RFC3339)
	url_, err := url.Parse(sns.Region.SNSEndpoint)
	if err != nil {
		return err
	}

	sign(sns.Auth, "GET", "/", params, url_.Host)
	url_.RawQuery = multimap(params).Encode()
	r, err := http.Get(url_.String())
	if err != nil {
		return err
	}
	defer r.Body.Close()

	//dump, _ := http.DumpResponse(r, true)
	//println("DUMP:\n", string(dump))
	//return nil

	if r.StatusCode != 200 {
		return buildError(r)
	}
	err = xml.Unmarshal(r.Body, resp)
	return err
}

func buildError(r *http.Response) os.Error {
	errors := xmlErrors{}
	xml.Unmarshal(r.Body, &errors)
	var err Error
	if len(errors.Errors) > 0 {
		err = errors.Errors[0]
	}
	err.RequestId = errors.RequestId
	err.StatusCode = r.StatusCode
	if err.Message == "" {
		err.Message = r.Status
	}
	return &err
}

func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		q[k] = []string{v}
	}
	return q
}
