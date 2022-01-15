start-sqs:
	docker run -p 9324:9324 -p 9325:9325 -v `pwd`/queues:/opt/elasticmq.conf softwaremill/elasticmq

start-gim:
	docker run -p 8080:8080 \
		-e GIMCONFIG='{"gimconsumer": {"consumer": [{"queue": {"region": "eu-central-1","endpoint": "http://localhost:9324","queueName": "pink_panther"},"forwardUrl": "http://localhost:9999/echo","processingTimeout": "2s","pollTimeout": "3s","ackTimeout": "3s","bulkReadSize": 10},{"queue": {"region": "eu-central-1","endpoint": "http://localhost:9324","queueName": "blue_rabbit"},"forwardUrl": "http://localhost:9999/echo","processingTimeout": "2s","pollTimeout": "3s","ackTimeout": "3s","bulkReadSize": 10}]},"gimproducer": {"server": {"port": "8080"},"producer": [{"queue": {"region": "eu-central-1","endpoint": "http://localhost:9324","queueName": "blue_rabbit"},"producerName": "bluerabbit","sendTimeout": "3s"},{"queue": {"region": "eu-central-1","endpoint": "http://localhost:9324","queueName": "pink_panther"},"producerName": "pinkpanther","sendTimeout": "3s"}]}, "gimlogger": {"level": "debug" }}' \
		-e AWS_ACCESS_KEY_ID='foo' \
		-e AWS_SECRET_ACCESS_KEY='bar' \
		gim
