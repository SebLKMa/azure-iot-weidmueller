# gomqttpubmodule

This Go application uses the open-source Go MQTT package as well as these official Azure MQTT documentation:    
[Using MQTT directly as an Azure IoT Device](https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-mqtt-support#using-the-mqtt-protocol-directly-as-a-device)  
[Using MQTT directly as an Azure IoT Module](https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-mqtt-support#using-the-mqtt-protocol-directly-as-a-module)  

Before building the app, please do your own `go mod init` and `go mod tidy`.  

## Generating SAS Key

Manually generate SAS for now and paste into code.  
Actual build script can generate and store in ENV. Code can then read from ENV.  
E.g.  
```sh
az iot hub generate-sas-token -d weidmueller-model-x -n <hubname> --du 86400
```
Helper notes:  
1 week = 60 * 60 * 24 * 7  
1 year = 8760 hrs  
Using VS Code `AZURE IOT HUB | device | right-click | Generate SAS Token for Device | 9000 hours`  
```sh
[SASToken] SAS token for [device name] is generated and copied to clipboard:  
SharedAccessSignature sr=<hubname>.azure-devices.net%2Fdevices%2F<device name>&sig=...se=...
```

## Build Go binary for arm

For Linux arm32 like Raspberry Pi, BeagleBone and Weidmueller-GW30:  
Example:  
```sh
GOOS=linux GOARCH=arm GOARM=5 go build -o gomqttpubgw30 main.go
```

## Push to Azure Registry using Az CLI command

Build and Push Go binary for arm32v7 to Azure Registry(a DockerHub):  
The Go binary specified in Dockerfile is already be pre-built for arm, so just deploy it to Azure Registry  
```sh
az acr build -t sebregistry.azurecr.io/gomqttpubmodule_gw30:0.0.1 -r sebregistry . -f Dockerfile.gw30 --platform linux/arm/v7
```

Verify the image is deployed in Azure Registry.  

## Push the Deployed Image to Edge Device

From `IoT Hub | IoT Edge | Devices | <the device id> | Set Modules`  
- add the image to be pushed down to device.  

**If deploying the same tag version**, e.g. sebregistry.azurecr.io/gomqttpubmodule_arm32v7:0.0.1-arm32v7, **you have to delete existing deployment from Set Modules**, then add the same deployment again. This is so that the device must fetch new image from Azufre Registry again.  

Verify the image is deployed and running (via cloud Azure IoT Hub and edge device's Portainer).  

## Monitor Device-to-Cloud(D2C) messages
Use [goeventhub](https://github.com/seblkma/azure-iot-weidmueller/tree/master/goeventhub) to monitor D2C messages.  
