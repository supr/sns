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
