package sns_test

var TestListTopicsXmlOK = `
<?xml version="1.0"?>
<ListTopicsResponse xmlns="http://sns.amazonaws.com/doc/2010-03-31/">
  <ListTopicsResult>
    <Topics>
      <member>
        <TopicArn>arn:aws:sns:us-west-1:331995417492:Transcoding</TopicArn>
      </member>
    </Topics>
  </ListTopicsResult>
  <ResponseMetadata>
    <RequestId>bd10b26c-e30e-11e0-ba29-93c3aca2f103</RequestId>
  </ResponseMetadata>
</ListTopicsResponse>
`

var TestCreateTopicXmlOK = `
<?xml version="1.0"?>
<CreateTopicResponse xmlns="http://sns.amazonaws.com/doc/2010-03-31/">
  <CreateTopicResult>
    <TopicArn>arn:aws:sns:us-east-1:123456789012:My-Topic</TopicArn>
  </CreateTopicResult>
  <ResponseMetadata>
    <RequestId>a8dec8b3-33a4-11df-8963-01868b7c937a</RequestId>
  </ResponseMetadata>
</CreateTopicResponse>
`

var TestDeleteTopicXmlOK = `
<DeleteTopicResponse xmlns="http://sns.amazonaws.com/doc/2010-03-31/">
  <ResponseMetadata>
    <RequestId>f3aa9ac9-3c3d-11df-8235-9dab105e9c32</RequestId>
  </ResponseMetadata>
</DeleteTopicResponse>
`

var TestListSubscriptionsXmlOK = `
<ListSubscriptionsResponse xmlns="http://sns.amazonaws.com/doc/2010-03-31/">
  <ListSubscriptionsResult>
    <Subscriptions>
      <member>
        <TopicArn>arn:aws:sns:us-east-1:698519295917:My-Topic</TopicArn>
        <Protocol>email</Protocol>
        <SubscriptionArn>arn:aws:sns:us-east-1:123456789012:My-Topic:80289ba6-0fd4-4079-afb4-ce8c8260f0ca</SubscriptionArn>
        <Owner>123456789012</Owner>
        <Endpoint>example@amazon.com</Endpoint>
      </member>
    </Subscriptions>
  </ListSubscriptionsResult>
  <ResponseMetadata>
    <RequestId>384ac68d-3775-11df-8963-01868b7c937a</RequestId>
  </ResponseMetadata>
</ListSubscriptionsResponse>
`

var TestGetTopicAttributesXmlOK = `
<GetTopicAttributesResponse xmlns="http://sns.amazonaws.com/doc/2010-03-31/">
  <GetTopicAttributesResult>
     <Attributes>
       <entry>
         <key>Owner</key>
         <value>123456789012</value>
       </entry>
       <entry>
         <key>Policy</key>
         <value>{"Version":"2008-10-17","Id":"us-east-1/698519295917/test__default_policy_ID","Statement" : [{"Effect":"Allow","Sid":"us-east-1/698519295917/test__default_statement_ID","Principal" : {"AWS": "*"},"Action":["SNS:GetTopicAttributes","SNS:SetTopicAttributes","SNS:AddPermission","SNS:RemovePermission","SNS:DeleteTopic","SNS:Subscribe","SNS:ListSubscriptionsByTopic","SNS:Publish","SNS:Receive"],"Resource":"arn:aws:sns:us-east-1:698519295917:test","Condition" : {"StringLike" : {"AWS:SourceArn": "arn:aws:*:*:698519295917:*"}}}]}</value>
       </entry>
       <entry>
         <key>TopicArn</key>
         <value>arn:aws:sns:us-east-1:123456789012:My-Topic</value>
       </entry>
     </Attributes>
  </GetTopicAttributesResult>
  <ResponseMetadata>
    <RequestId>057f074c-33a7-11df-9540-99d0768312d3</RequestId>
  </ResponseMetadata>
</GetTopicAttributesResponse>
`

var TestPublishXmlOK = `
<PublishResponse xmlns="http://sns.amazonaws.com/doc/2010-03-31/">
  <PublishResult>
    <MessageId>94f20ce6-13c5-43a0-9a9e-ca52d816e90b</MessageId>
  </PublishResult>
  <ResponseMetadata>
    <RequestId>f187a3c1-376f-11df-8963-01868b7c937a</RequestId>
  </ResponseMetadata>
</PublishResponse>
`

var TestSetTopicAttributesXmlOK = `
<SetTopicAttributesResponse xmlns="http://sns.amazonaws.com/doc/2010-03-31/">
  <ResponseMetadata>
    <RequestId>a8763b99-33a7-11df-a9b7-05d48da6f042</RequestId>
  </ResponseMetadata>
</SetTopicAttributesResponse>
`
