
# Processor Worker FFmpeg

## Basic Functional Pattern

after initialization worker is idle and waits for messages

0. worker periodically receives a tick message and responds with worker_idle
1. then the worker receives a job_assignment
2. validates, if it is really meant for him ($SERVICE_UUID)
3. sends back a job_accepted messages
4. if for any reason another job_assignment arrives after this point, it is silently dropped (just ignore it. it should not happen anyways).
4. starts processing the job
5. sends out a worker_busy message (containing the currently processed job) as a response to each tick message. worker_busy messages can also be sent more frequently - no limit applies.
6. on each progress update received from ffmpeg, we send out an additional worker_busy message containing ffmpeg's current progress.
7. on completion a job_finished message is sent including stdout/err and exitcode data
8. a worker_idle message is sent right after the job_finished message to minimize idle time. 
9. back to the top

## Prerequisites

* Use the supplied Dockerfile for testing and development. 
* It does already include up-to-date ffmpeg binaries (USE THESE! don't bring your own!).
* We are using multistage build, so we can e.g. compile ffmpeg in one step and then just copy over the resulting binaries into the final image. Which is exactly, what we are doing here.
* I presume node.js is the easiest way forward in this case, so our final Docker image derives from the official Docker node image. If you use another language, feel free to change that. But it would make sense not to in this case!
* If you need additional tools, add their setup instructions to the Dockerfile.
* We use Alpine Docker images (don't change that; we love alpine).
* Make sure ffmpeg is run with the same ENV as its parent processor-worker-ffmpeg.
* All output from processor-worker-ffmpeg needs to follow our json-pipe-protocol - one json object per line. Details in next topic.
* Everything going to stdout will end up on the message bus as is. Each line represents on message on the bus. The message should already contain all necessary fields, no additional fields are added  upstream, by the service-adapter listening on the other side of stdout. A message MUST have a set topic field.
* Everything going to stderr will be sent to a special logging topic. Log messages are much more simple. $SERVICE_UUID etc. are added upstream.

## JSON Pipe Protocol

* This services communicates through stdin, stdout and stderr. 
* One line of input/output delimited by '\n' represents one message.
* stdin/stdout messages are json objects.
* The general schema of a json-pipe-protocol message is defined below:
```
{
  "topic": "$NAMESPACE/topic_name",
  "service_uuid": "$SERVICE_UUID",
  "service_name": "$SERVICE_NAME",
  "service_host": "$SERVICE_HOST",
  "created_at": "iso_time: now",
  "tick_reference": {
    "uuid": "uuid: of the referenced tick",
    "created_at": "iso_time of timestamp.microseconds: of the referenced tick"
  },
  "payload": {}
}
```
* Of course, the real message has no newlines between fields, only one after the final '}' which thereby terminates the message.
* $NAMESPACE is actually either $NAMESPACE_LISTENER or $NAMESPACE_PUBLISHER, depending on whether the message is read from stdin or written to stdout.
* The SERVICE as well as the NAMESPACE variables are set in the environment and can be retrieved and inserted that way.
* A tick is a periodic *hello* message, consisting of the uuid of the tick and a created_at timestamp.
* Services can respond to such a tick, in which case, they will carry the original tick uuid and timestamp as a reference.
* stderr messages are json objects with a simplified schema. Topic, service attributes etc will be set upstream and therefore don't need to be added here. The general stderr message schema looks like this:
```
{
  "log_level": "one of critical|error|warning|info|debug",
  "log_message": "string: make sure \"double quotes\" are escaped in this string, if there are any.",
  "created_at": "iso_time"
}
```
Same here, no newlines except after final '}' to terminate the message.

## Message Schemata

Included in this git tree are schema files for every message type handled by this service.

### input

* tick
* job_assignment

### output

* job_accepted
* job_finished
* worker_idle
* worker_busy

### error

* general error schema

## Implementation recommendations

There are at least two javascript ffmpeg progress wrappers available, which can be used in order to retrieve progress data from stdout/err output.

* https://github.com/eugeneware/ffmpeg-progress-stream
* https://github.com/legraphista/ffmpeg-progress-wrapper
