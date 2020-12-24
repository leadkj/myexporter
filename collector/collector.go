package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	ip2 "myexporter/ip"
	"os"
	"sync"
)

//定义指标结构体
type Metrics struct {
	metrics map[string]*prometheus.Desc
	mutex sync.Mutex
}

//创建指标
func CreateMetric(namespace string,metricName string,docString string,labels []string) *prometheus.Desc{
	return  prometheus.NewDesc(
		namespace+"_"+metricName,
		docString,
		labels,
		nil)

}


func NewMetrics(namespace string) *Metrics{
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			"server_name":CreateMetric(namespace,"server_name","the name of server",[]string{"name"}),
			"server_ip":CreateMetric(namespace,"server_ip","the ip of server",[]string{"ip"}),
		},
	}
}

//定义方法
func (m Metrics) Describe(ch chan <- *prometheus.Desc){
	for _,m := range m.metrics{
		ch <- m
	}
}

func (m Metrics) Collect(ch chan<-prometheus.Metric){
	m.mutex.Lock()
	defer m.mutex.Unlock()
	Infodata,Ipdata := m.GetData()
	for hostname,currentValue:= range Infodata{
		ch<- prometheus.MustNewConstMetric(m.metrics["server_name"],prometheus.CounterValue,float64(currentValue),hostname)
	}
	for ip,currectValue:=range Ipdata{
		ch<-prometheus.MustNewConstMetric(m.metrics["server_ip"],prometheus.CounterValue,float64(currectValue),ip)
	}
}

func (m Metrics) GetData()(Infodata map[string]float64,Ipdata map[string]float64){
	ip :=ip2.GetOutboundIP()
	hname,_:=os.Hostname()

	Infodata= map[string]float64{hname:0}
	Ipdata = map[string]float64{ip:0}
	return

}

