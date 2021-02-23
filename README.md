# client_standardclient
- client에서 RetryMax(5), RetryWaitMin(3 second)를 설정한 상태로 실행한다. 
- server가 실행중인 상태가 아니므로 요청을 재시도 한다.
- client가 재시도를 하는 중간에 server를 실행한다.
- client

# client_timeout
- server에 PUT 요청이 들어오면 5초뒤에 응답을 한다.
- client는 timeout을 3초로 설정하고 server에 PUT요청을 한다.
