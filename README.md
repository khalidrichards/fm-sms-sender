# fm-sms-sender
A WIP AWS Lambda Function meant to be used to send text messages. 

## To-Dos

### 1. Set up Secrets using AWS Secrets Manager
The Twilio API uses environment variables to validate the SMS/MMS request. We should make sure the environment variables are set up
using secrets manager, then write the appropriate code to get the correct environment variables.

### 2. Set up data store to store send lists
While our initial use-case for the SMS/MMS sender will be for food pantry events, we'll want to be flexible and have the capability to send SMS/MMS messages to various groups of people. We shoudl set up a data store that will make it easy to fetch who is in any particular 
send list.

### 3. Create a Queue to start processing SMS.
According to the Twilio API, [we cannot send multiple numbers the same SMS using one API call](https://help.twilio.com/articles/223181548-Sending-multiple-or-bulk-messages-with-a-single-API-request-from-a-list-of-recipients). As a result, whenever we want to send an SMS/MMS, we should actually send this to a lambda function that will grab the relevant send list, queue up batches of 10-20 numbers to send messages to, then have another lambda function send those messages. 

### 4. Send multiple users the same SMS.
According to the Twilio API, [we cannot send multiple numbers the same SMS using one API call](https://help.twilio.com/articles/223181548-Sending-multiple-or-bulk-messages-with-a-single-API-request-from-a-list-of-recipients). As a result, we should refactor the code
in `sendsms.go` to accept a batch of 10-20 numbers to iterate and send SMS to (this will depending on things like timings and time-outs). These batches shoudl be picked up from the queue created in the previous To-Do item.

### 5. Send multiple users the same MMS.
Eventually, we will want to send users MMS with things like flyers for food pantry events. We'll likely want some way of storing the MMS assets (presumably in S3), adding a path to the asset in the queueing process (we can make the assumption that MMS messages will the image of the asset followed by text), then the other information needed by SMS.
