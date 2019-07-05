package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/streadway/amqp"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	//Rabbitmq config
	User      string
	Password  string
	Host      string
	Port      string
	Queue     string
	Consumer  string
	Autoack   bool
	Exclusive bool
	Nolocal   bool
	Nowait    bool
	Args      string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func start_java_app_test() {
	cmd_prep := "echo 'Hello World'"
	cmd_output := exec.Command("bash", "-c", cmd_prep)
	stdout, err := cmd_output.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(stdout)
	fmt.Printf("%s\n", stdout)
}

func main() {
	configuration := Configuration{}
	err := gonfig.GetConf("config.json", &configuration)
	if err != nil {
		panic(err)
	}

	user := configuration.User
	password := configuration.Password
	host := configuration.Host
	port := configuration.Port

	connstring := "amqp://" + user + ":" + password + "@" + host + ":" + port + "/"

	conn, err := amqp.Dial(connstring)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queue := configuration.Queue
	consumer := configuration.Consumer

	msgs, err := ch.Consume(
		queue,    // queue
		consumer, // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			start_java_app_test()
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
