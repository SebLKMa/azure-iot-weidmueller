# GoMqttPubModuleId

Demonstrates Go application publishing MQTT message to Azure IoT Hub Topic with **ModuleId**.  
Example:  
```go
// Topic - devices/{device_id}/modules/{module_id}/messages/events/
const MqttTopic = "devices/weidmueller-model-x/modules/gomqttpubmoduleid/messages/events/"
```

## Edge Device Pre-requisites
Ensure Edge Device is on-line and its Azure IoT Edge Runtime service is up and running.  
This is so that IoT Hub can stream new built images, configurations and desired statuses to the device.  

If the device is up and running and already connected to IoT Hub, there is no need for any physical access to it. The typical steps are:  
1. Write your Go app.  
2. Generate SAS key.
3. Build your app to executable binary.  
4. Deploy the executable binary image to docker hub (Azure registry).   
5. Push the binary image from Azure IoT docker hub to the device.  

##  1. Write your Go app

This section contains sample POC codes extracted from `main.go`.  

```go
// Topic - devices/{device_id}/modules/{module_id}/messages/events/
const MqttTopic = "devices/weidmueller-model-x/modules/gomqttpubmoduleid/messages/events/"

func init() {
	var broker = "<hubname>.azure-devices.net"
	var port = 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetProtocolVersion(4)
    // Set the client ID to {device_id}/{module_id}
	opts.SetClientID("weidmueller-model-x/gomqttpubmoduleid") 
    // <hubname>.azure-devices.net/{device_id}/{module_id}/?api-version=2018-06-30
	opts.SetUsername("<hubname>.azure-devices.net/weidmueller-model-x/gomqttpubmoduleid/?api-version=2020-09-30") 

	// NOTE: This demo manually generate SAS. This can also be part of CI/CD to set environment variable.
	// az iot hub generate-sas-token -d weidmueller-iot-device -m gomqttpubmoduleid -n seb-hub --du <seconds>
	opts.SetPassword(
		"dummy string")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	mqttClient = mqtt.NewClient(opts)
}
```

The first time you deploy this app via Azure IoT Hub web portal's **Set module** functionality, you can set a `"dummy string"` in `SetPassword()` below.  
This is because the command line `az iot hub generate-sas-token -d weidmueller-iot-device -m gomqttpubmoduleid ...` expects the module *gomqttpubmoduleid* to exist in device's list of modules.
In fact, you can even create and deploy a dummy *gomqttpubmoduleid* app. Of course, please set its **Desired Status** to **stopped** as there is no point running this placeholder dummy app.

## 2. Generate SAS Key

Manually generate SAS.  
Examples:  
```sh
az iot hub generate-sas-token -d sebEdgeDevice -m NodeJsModuleId -n seb-hub --du 86400
```
```
az iot hub generate-sas-token -d weidmueller-iot-device -m gomqttpubmoduleid -n seb-hub --du 86400
```

Now update `SetPassword()` in code.  
```go
    opts.SetPassword(
    "SharedAccessSignature sr=<hubname>.azure-devices.net%2Fdevices%2Fweidmueller-model-x%2Fmodules%2Fgomqttpubmoduleid&sig=...")
``` 

NOTE: Actual build script can generate and store in ENV. Code can then read from ENV.  



## 3. Build your app to executable binary

For Linux amd64 like Ubuntu:  
```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gomqttpubmoduleid main.go
```

For Linux arm32 like Raspberry PI, BeagleBone, Weidmueller GW30:  
```sh
GOOS=linux GOARCH=arm GOARM=5 go build -o gomqttpubmoduleid main.go
``` 

## 4. Deploy the executable binary image to docker hub (Azure registry)

Examine your Dockerfile is correctly scripted (e.g. Dockerfile.gw30).  Thencan queue your build to Azure cloud.

Example - Build and Push Go app for arm32v7:  
The Go binary specified in Dockerfile must already be pre-built for arm, so just deploy it to Azure Registry  
```sh
az acr build -t sebregistry.azurecr.io/gomqttpubmoduleid_gw30:<version tag> -r sebregistry . -f Dockerfile.gw30 --platform linux/arm/v7
```

From Azure Registry, verify the image version is successfully deployed.  

## 5. Push the binary image from Azure IoT docker hub to the device

You do this using Azure IoT Hub web portal's **Set module** functionality again.  
You can now **Desired Status** to **running**.
Once successfully deployed to device, the app module will start sending *Hello* messages to the MQTT topic.  

You can use `goeventhub` to see receiving of these messages and Azure IoT Hub web portal to see the device module's log output.  