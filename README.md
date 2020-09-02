# abstract-runner
 The abstract-runner was a gitlab-runner like tool which was worked for pulling the project from the remote git server, running and debugging in a k8s pod.

## git to do
- install git
- combine git commands
- needs a git hook or custom trigger for triggering the worker to take operations
- `Important` git must remove the local branch which has been removed by others.

## worker to do 
- async queue
- status
- full log 
- mode (e.g. auto or manual)

## api to do
- http
- tcp