golofka
=======
Log processing CLI to process stream data using Kafka and store in Elasticsearch for later analysis. golofka provides a command line interface to read the stream of log data and process through kafka queues. mapper and processor are the two functions which the developer will be able to customize based on the requirement. mapper function will be taking each line object from the log file as input and map the column value to an object and return the object. This object will be pushed to kafka queue topics. There could be more than one topic. This kafka queue is configurable by the developers through a JSON or YAML file. 

Once the api takes the object from the queue , it will then be passed to processor function for further computation. Yhis api provides the flexibility to store the object in Elasticsearch for further log analysis and indexing.

####Command Line Interface:
golofka -startserver        //This reads starts the kafka, elastic search and api servers
golofka -initapi            //This initlizes the api by reading the JSON/YAML file and keep the service ready
golofka -version            //prints the version of the api
golofka -streamdata         //Tiis starts streaming the data from the log file and call mapper function before pushing to queue and then call processor after popping from the queue.
