
### Basic Functional Pattern

after initialization worker is idle and waits for messages

0. worker periodically receives a tick message and responds with worker_idle
1. then the worker receives a job_assignment
2. validates, if it is really meant for him
3. sends back a job_accepted messages
4. if for any reason another job_assignment arrives after this point, it is silently dropped (just ignore it. it should not happen anyways).
4. starts processing the job
5. sends out a worker_progress message as a response to each tick message
6. worker_progress can also be sent more often - no limit applies
7. on completion a job_finished message is sent including stdout/err and exitcode data
8. a worker_idle message is sent right after the job_finished message to minimize idle time. 
9. back to the top

### Prerequisites

* Use the supplied Dockerfile for testing and development. 
* It does already include up-to-date ffmpeg binaries (USE THAT! don't bring your own!)
* Our Docker image derives from the official Docker node image (use whatever you like here; if you like to drop node.js and use something else entirely, so be it. But it would make sense not to in this case!)
* If you need additional tools, add their setup instructions to the Dockerfile
* The Docker image is based on alpine (don't change that; we love alpine).
* Make sure ffmpeg is run with the same env as its parent processor-worker-ffmpeg
* All output from processor-worker-ffmpeg needs to be json messages following the json-pipe-protocol - one message object per line
* Everything going to stdout will end up on the message bus as is.
* Everything going to stderr will be sent to a special logging topic.

There are at least two javascript ffmpeg progress wrappers available, which can be used in order to retrieve progress data from stdout/err output.

* https://github.com/eugeneware/ffmpeg-progress-stream
* https://github.com/legraphista/ffmpeg-progress-wrapper