package toolLib

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/streadway/amqp"
)
//定义结构体
type RabbitMq struct {
	amqpUrl string
	connect *amqp.Connection
	channel *amqp.Channel
}

//初始化
//@param  interface params 可变参数，若传值，则传user、pwd,host,port 四个参数
//@return *RabbitMq
func  NewRabbitMq(params ...interface{})(*RabbitMq,error){
	var url string
	var err error
	if len(params) == 0{
         url = fmt.Sprintf("amqp://%s:%s@%s:%s/",beego.AppConfig.String("rabbitmq.user"),
         	beego.AppConfig.String("rabbitmq.pwd"),
         	beego.AppConfig.String("rabbitmq.host"),
         	beego.AppConfig.String("rabbitmq.port"))
	}else{
		if len(params) != 4{
			err = errors.New("param must four.")
		}
		url = fmt.Sprintf("amqp://%s:%s@%s:%s/",params[0],
			params[1],
			params[2],
			params[3])
	}
	connect,err:= amqp.Dial(url)
	return &RabbitMq{
		amqpUrl: url,
		connect: connect,
	},err
}

//生产者
//@param  string exchange 交换机名称
//@param  string content  队列内容
//@param  string params   可变参数 只接受一个第一个参数，用于赋值给routingkey
//@return error
func (r *RabbitMq) Publish(exchange string,content string,params ...string)(error){
	var routingKey string
	if len(params) != 0{
		routingKey = params[0]
	}
	ch, err := r.connect.Channel()
	if err != nil{
		return errors.New(err.Error())
	}
	r.channel = ch
	defer r.close()
	err = ch.Publish(
		exchange,                   //exchange fanout
		routingKey,                 // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			DeliveryMode: 2,
			ContentType:  "text/plain",
			Body:         []byte(content),
		})
	return err
}

//消费者
//@param string    queueName 队列名称
//@param interface params    可变参数，若传值、则只传一个bool值即可，用于是否收到消息需要确认，默认是需要确认
//@return <-chan amqp.Delivery
func (r *RabbitMq) Consume(queueName string,params ...interface{})(<-chan amqp.Delivery,error){
	ch, err := r.connect.Channel()
	if err != nil{
		return nil,errors.New(err.Error())
	}
	r.channel = ch
	defer  r.close()
	var autoAck bool
	if len(params) != 0{
		autoAck = params[0].(bool)
	}
	msgs, err := ch.Consume(
		queueName, // queue
		"",     // consumer
		autoAck,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	return msgs,err
}


//关闭channel
func (r *RabbitMq) close(){
	err := r.channel.Close()
	if err != nil{
		fmt.Println("close channel fail,",err.Error())
	}
}






