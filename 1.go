package main

import ("amqp"

)
func main() {
	conn, _:= amqp.Dial("amqp://guest:guest@localhost:5672/")

	ch, _ :=conn.Channel()

	q,_ :=ch.QueueDeclare(
		"firrst.queue",
		true,
		false,
		false,
		false,
		nil, )

	msg := amqp.Publishing{
		Body: []byte("my first message"),
	}

	ch.Publish(
		"",
		q.Name,
		false,
		false,
		msg)

	msgs, _ := ch.Consume (
		q.Name,
		""	,
		true,
		false,
		false,
		false,
		nil)

for m := range msgs {
	println(string(m.Body))
 }
  }
