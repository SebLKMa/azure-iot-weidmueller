# nodejsmodule

The source code in this module is adapted from the official [Azure Node.js documentation](https://docs.microsoft.com/en-us/azure/iot-hub/quickstart-control-device-node).

Demonstrates both Device-to-Cloud and Cloud-to-Device using Azure IoT SDK.  
For Cloud-to-Device, Azure IoT SDK **Direct Method** is used. *Direct Methods* can thus be used as *Facades* to actual module running in the device.

This application sends Temperature and Humidity to cloud Azure Event Hub.  
It also responds to *Direct Method* for setting the send interval.  

Azure IoT SDK **Direct Method** handles the underlying messaging reliability and durability.  

## Build and Push NodeJs app for arm to Azure Registry(a DockerHub):
This example builds a NodeJs app for arm-based hardware:  
```sh
az acr build -t sebregistry.azurecr.io/nodejsmodule_gw30:0.0.1 -r sebregistry . -f Dockerfile.gw30.arm --platform linux/arm/v7
```

Verify the image is deployed in Azure Registry.  

## Push the Deployed Image to Edge Device

From `IoT Hub | IoT Edge | Devices | <the device id> | Set Modules`  
- add the image to be pushed down to device.  

**If deploying the same tag version**, e.g. sebregistry.azurecr.io/nodejsmodule:0.0.1, **you have to delete existing deployment from Set Modules**, then add the same deployment again. This is so that the device must fetch new image from Azufre Registry again.  

Verify the image is deployed and running (via cloud Azure IoT Hub and edge device's Portainer).  

## Monitor Device-to-Cloud(D2C) messages
Use [goeventhub](https://github.com/seblkma/azure-iot-weidmueller/tree/master/goeventhub) to monitor D2C messages. 

## Cloud-to-Device(C2D)

Use [cloudbackend](https://github.com/seblkma/azure-iot-weidmueller/tree/master/cloudbackend) to demonstrate C2D.
