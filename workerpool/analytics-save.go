package workerpool

import (
	"fmt"
	"github.com/todzuko/analytics-service/database/models"
	"os"
	"strconv"
)

var workerQueue chan models.AnalyticsData

func StartWorkerPool() {
	workerCount, _ := strconv.Atoi(os.Getenv("WORKER_CNT"))
	workerQueue = make(chan models.AnalyticsData, workerCount)
	for i := 0; i < workerCount; i++ {
		go saveAnalytics()
	}
}

func QueueAnalytics(analytics models.AnalyticsData) {
	workerQueue <- analytics
}

func saveAnalytics() {
	for analytics := range workerQueue {
		fmt.Println("worker working work")
		models.Save(&analytics)
	}
}
