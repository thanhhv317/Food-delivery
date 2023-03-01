package subscriber

import (
	"context"
	"golang/common"
	"golang/component/appctx"
	"golang/component/asyncjob"
	"golang/pubsub"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type subscriberEngine struct {
	appCtx appctx.AppContext
}

func NewEngine(appContext appctx.AppContext) *subscriberEngine {
	return &subscriberEngine{appCtx: appContext}
}

func (engine *subscriberEngine) Start() error {
	engine.setup()

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *subscriberEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("Setup subscriber for:", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
			//return nil
		}
	}

	go func() {
		defer common.Recover()
		for {
			msg := <-c

			// Link array of consumer jobs => array of async job handlers

			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
