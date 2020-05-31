package taillog

// 专门收集日志文件模块
import (
	"fmt"
	"go_git/Go_learn/studygo/logagent/kafka"

	"context"

	"github.com/hpcloud/tail"
)

// TailTask 一个手机日志的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// 为了能实现退出t.run()
	ctx        context.Context
	cancelFunc context.ConcelFunc
}

func newTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		concelFunc: cancel,
	}
	tailObj.init() // 根据路径去打开对于的日志
	return
}

func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的那个地方开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}

	// 当goroutine执行的函数退出，goroutine退出
	go t.run() // 直接去采集日志发送到kafka
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printlf"tail task:%s_%s 结束了...\n",t.path,t.topic)
			return
		case line := <-t.instance.Lines: // 从tailObj的通道中一行一行的读取日志
			// 3.2发往kafka
			// kafka.SendToKafka(t.topic, line.Text) // 函数调用函数
			// 先把日志数据发送到一个通道中
			kafka.SnedToChan(t.topic, line.Text)
			// kafka包中有单独的goroutine去取日志数据发送到kafka
		}
	}
}
