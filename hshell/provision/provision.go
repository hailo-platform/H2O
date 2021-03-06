package provision

import (
	log "github.com/cihub/seelog"
	"github.com/hailo-platform/H2O/platform/client"
	"github.com/hailo-platform/H2O/platform/raven"
	"github.com/hailo-platform/H2O/hshell/login"
	"github.com/hailo-platform/H2O/hshell/util"
	"github.com/hailo-platform/H2O/protobuf/proto"
	create "github.com/hailo-platform/H2O/provisioning-manager-service/proto/create"
	del "github.com/hailo-platform/H2O/provisioning-manager-service/proto/delete"
	search "github.com/hailo-platform/H2O/provisioning-manager-service/proto/search"
	restart "github.com/hailo-platform/H2O/provisioning-service/proto/restart"
	restartaz "github.com/hailo-platform/H2O/provisioning-service/proto/restartaz"
	"strconv"
)

var AmqpHost string

var bufferedServices []string = make([]string, 0)

func GetAvailableServicesBuffered() []string {
	if len(bufferedServices) == 0 || util.CacheReload["proveservicesbuffered"] {
		GetAvailableServices()
	} else {
		go GetAvailableServices()
	}

	return bufferedServices
}

func GetAvailableServices() []string {
	res, _ := callSearch("", "")
	srvs := make([]string, 0)
	for _, srv := range res {
		srvs = append(srvs, *srv.ServiceName)
	}
	bufferedServices = srvs
	util.CacheReload["proveservicesbuffered"] = false
	return srvs
}

var bufferedVersions []string = make([]string, 0)

func GetAvailableVersionsBuffered(service string) []string {
	if len(bufferedVersions) == 0 || util.CacheReload["provversionsbuffered"] {
		GetAvailableVersions(service)
	} else {
		go GetAvailableVersions(service)
	}

	return bufferedVersions
}

func GetAvailableVersions(service string) []string {
	res, _ := callSearch(service, "")
	versions := make([]string, 0)

	for _, vers := range res {
		versions = append(versions, strconv.FormatUint(*vers.ServiceVersion, 10))
	}
	bufferedVersions = versions
	util.CacheReload["provversionsbuffered"] = false
	return versions
}

func CreateService(service string, class string, version int) error {
	return callCreate(service, class, version)
}

func DeleteService(service string, class string, version int) error {
	return callDelete(service, class, version)
}

var bufferedMachineClass []string = make([]string, 0)

func GetServiceMachineClassesBuffered(service string) []string {
	if len(bufferedVersions) == 0 || util.CacheReload["machineclassbuffered"] {
		GetServiceMachineClasses(service)
	} else {
		go GetServiceMachineClasses(service)
	}

	return bufferedMachineClass
}

func GetServiceMachineClasses(service string) []string {
	searchRes, _ := callSearch(service, "")
	classArr := make([]string, 0)
	for _, val := range searchRes {
		classArr = append(classArr, val.GetMachineClass())
	}
	bufferedMachineClass = classArr
	util.CacheReload["machineclassbuffered"] = false
	return classArr
}

type ServiceVersionMachine struct {
	Service      string
	Version      int
	MachineClass string
}

func GetServiceVersionMachineClass() ([]*ServiceVersionMachine, error) {
	searchRes, err := callSearch("", "")
	if err != nil {
		return nil, err
	}
	svm := make([]*ServiceVersionMachine, len(searchRes))
	for i, res := range searchRes {
		svm[i] = &ServiceVersionMachine{}
		svm[i].MachineClass = res.GetMachineClass()
		svm[i].Service = res.GetServiceName()
		svm[i].Version = int(res.GetServiceVersion())
	}

	return svm, err
}

func callCreate(service string, class string, version int) error {
	if len(AmqpHost) != 0 {
		raven.AmqpUri = AmqpHost
	}

	request, _ := client.NewRequest(
		"com.hailocab.kernel.provisioning-manager",
		"create",
		&create.Request{
			ServiceName:    proto.String(service),
			MachineClass:   proto.String(class),
			ServiceVersion: proto.Uint64(uint64(version)),
		},
	)
	request.SetSessionID(login.Session)
	request.SetFrom(login.FromService)
	response := &create.Response{}
	if err := util.SendRequest(request, response); err != nil {
		log.Critical(err)
		return err
	}

	return nil
}

func callDelete(service string, class string, version int) error {
	if len(AmqpHost) != 0 {
		raven.AmqpUri = AmqpHost
	}

	request, _ := client.NewRequest(
		"com.hailocab.kernel.provisioning-manager",
		"delete",
		&del.Request{
			ServiceName:    proto.String(service),
			MachineClass:   proto.String(class),
			ServiceVersion: proto.Uint64(uint64(version)),
		},
	)
	request.SetSessionID(login.Session)
	request.SetFrom(login.FromService)
	response := &del.Response{}
	if err := util.SendRequest(request, response); err != nil {
		log.Critical(err)
		return err
	}

	return nil
}

func callSearch(service string, class string) ([]*search.Result, error) {
	if len(AmqpHost) != 0 {
		raven.AmqpUri = AmqpHost
	}

	request, _ := client.NewRequest(
		"com.hailocab.kernel.provisioning-manager",
		"search",
		&search.Request{
			ServiceName:  proto.String(service),
			MachineClass: proto.String(class),
		},
	)
	request.SetSessionID(login.Session)
	request.SetFrom(login.FromService)
	response := &search.Response{}
	if err := util.SendRequest(request, response); err != nil {
		log.Critical(err)
		return nil, err
	}

	return response.Results, nil
}

func PubRestart(service, class string, version uint64, azname string) error {
	pub, err := client.NewPublication(
		"com.hailocab.kernel.provisioning.restart",
		&restart.Request{
			ServiceName:    proto.String(service),
			ServiceVersion: proto.Uint64(version),
			MachineClass:   proto.String(class),
			AzName:         proto.String(azname),
		},
	)
	if err != nil {
		return err
	}
	pub.SetSessionID(login.Session)
	return util.SendPublication(pub)
}

func PubRestartAz(azname string) error {
	pub, err := client.NewPublication(
		"com.hailocab.kernel.provisioning.restartaz",
		&restartaz.Request{
			AzName: proto.String(azname),
		},
	)
	if err != nil {
		return err
	}
	pub.SetSessionID(login.Session)
	return util.SendPublication(pub)
}
