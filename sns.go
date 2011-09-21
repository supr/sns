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
	"strconv"
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
	Endpoint        string
	Owner           string
	Protocol        string
	SubscriptionArn string
	TopicArn        string
}

func (topic *Topic) Message(message [8192]byte, subject string) *Message {
	return &Message{topic.SNS, topic, message, subject}
}

type ResponseMetadata struct {
	RequestId string
	BoxUsage  float64
}

type ListTopicsResponse struct {
	Topics    []Topic `xml:"ListTopicsResult>Topics>member"`
	NextToken string
	ResponseMetadata
}

type CreateTopicResponse struct {
	Topic Topic `xml:"CreateTopicResult>"`
	ResponseMetadata
}

type DeleteTopicResponse struct {
	ResponseMetadata
}

type ListSubscriptionsResponse struct {
	Subscriptions []Subscription `xml:"ListSubscriptionsResult>Subscriptions>member"`
	NextToken     string
	ResponseMetadata
}

type AttributeEntry struct {
	Key, Value string
}

type GetTopicAttributesResponse struct {
	Attributes []AttributeEntry `xml:"GetTopicAttributesResult>Attributes>entry"`
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

func (sns *SNS) GetTopicAttributes(TopicArn string) (resp *GetTopicAttributesResponse, err os.Error) {
	resp = &GetTopicAttributesResponse{}
	params := makeParams("GetTopicAttributes")
	params["TopicArn"] = TopicArn
	err = sns.query(nil, nil, params, resp)
	return
}

type PublishOpt struct {
	Message          string
	MessageStructure string
	Subject          string
	TopicArn         string
}

type PublishResponse struct {
	MessageId string `xml:"PublishResult>MessageId"`
	ResponseMetadata
}

func (sns *SNS) Publish(options *PublishOpt) (resp *PublishResponse, err os.Error) {
	resp = &PublishResponse{}
	params := makeParams("Publish")

	if options.Subject != "" {
		params["Subject"] = options.Subject
	}

	if options.MessageStructure != "" {
		params["MessageStructure"] = options.MessageStructure
	}

	if options.Message != "" {
		params["Message"] = options.Message
	}

	if options.TopicArn != "" {
		params["TopicArn"] = options.TopicArn
	}

	err = sns.query(nil, nil, params, resp)
	return
}

type SetTopicAttributesResponse struct {
	ResponseMetadata
}

func (sns *SNS) SetTopicAttributes(AttributeName, AttributeValue, TopicArn string) (resp *SetTopicAttributesResponse, err os.Error) {
	resp = &SetTopicAttributesResponse{}
	params := makeParams("SetTopicAttributes")

	if AttributeName == "" || TopicArn == "" {
		return nil, os.NewError("Invalid Attribute Name or TopicArn")
	}

	params["AttributeName"] = AttributeName
	params["AttributeValue"] = AttributeValue
	params["TopicArn"] = TopicArn

	err = sns.query(nil, nil, params, resp)
	return
}

type SubscribeResponse struct {
	SubscriptionArn string `xml:"SubscribeResult>SubscriptionArn"`
	ResponseMetadata
}

func (sns *SNS) Subscribe(Endpoint, Protocol, TopicArn string) (resp *SubscribeResponse, err os.Error) {
	resp = &SubscribeResponse{}
	params := makeParams("Subscribe")

	params["Endpoint"] = Endpoint
	params["Protocol"] = Protocol
	params["TopicArn"] = TopicArn

	err = sns.query(nil, nil, params, resp)
	return
}

type UnsubscribeResponse struct {
	ResponseMetadata
}

func (sns *SNS) Unsubscribe(SubscriptionArn string) (resp *UnsubscribeResponse, err os.Error) {
	resp = &UnsubscribeResponse{}
	params := makeParams("Unsubscribe")

	params["SubscriptionArn"] = SubscriptionArn

	err = sns.query(nil, nil, params, resp)
	return
}

type ConfirmSubscriptionResponse struct {
	SubscriptionArn string `xml:"ConfirmSubscriptionResult>SubscriptionArn"`
	ResponseMetadata
}

type ConfirmSubscriptionOpts struct {
	AuthenticateOnUnsubscribe string
	Token                     string
	TopicArn                  string
}

func (sns *SNS) ConfirmSubscription(options *ConfirmSubscriptionOpts) (resp *ConfirmSubscriptionResponse, err os.Error) {
	resp = &ConfirmSubscriptionResponse{}
	params := makeParams("ConfirmSubscription")

	if options.AuthenticateOnUnsubscribe != "" {
		params["AuthenticateOnUnsubscribe"] = options.AuthenticateOnUnsubscribe
	}

	params["Token"] = options.Token
	params["TopicArn"] = options.TopicArn

	err = sns.query(nil, nil, params, resp)
	return
}

type Permission struct {
	ActionName string
	AccountId  string
}

type AddPermissionResponse struct {
	ResponseMetadata
}

func (sns *SNS) AddPermission(permissions []Permission, Label, TopicArn string) (resp *AddPermissionResponse, err os.Error) {
	resp = &AddPermissionResponse{}
	params := makeParams("AddPermission")

	for i, p := range permissions {
		params["AWSAccountId.member."+strconv.Itoa(i+1)] = p.AccountId
		params["ActionName.member."+strconv.Itoa(i+1)] = p.ActionName
	}

	params["Label"] = Label
	params["TopicArn"] = TopicArn

	err = sns.query(nil, nil, params, resp)
	return
}

type RemovePermissionResponse struct {
	ResponseMetadata
}

func (sns *SNS) RemovePermission(Label, TopicArn string) (resp *RemovePermissionResponse, err os.Error) {
	resp = &RemovePermissionResponse{}
	params := makeParams("RemovePermission")

	params["Label"] = Label
	params["TopicArn"] = TopicArn

	err = sns.query(nil, nil, params, resp)
	return
}

type ListSubscriptionByTopicResponse struct {
	Subscriptions []Subscription `xml:"ListSubscriptionsByTopicResult>Subscriptions>member"`
	ResponseMetadata
}

type ListSubscriptionByTopicOpt struct {
	NextToken string
	TopicArn  string
}

func (sns *SNS) ListSubscriptionByTopic(options *ListSubscriptionByTopicOpt) (resp *ListSubscriptionByTopicResponse, err os.Error) {
	resp = &ListSubscriptionByTopicResponse{}
	params := makeParams("ListSbubscriptionByTopic")

	if options.NextToken != "" {
		params["NextToken"] = options.NextToken
	}

	params["TopicArn"] = options.TopicArn

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
