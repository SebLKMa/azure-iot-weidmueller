# azure-iot-poc-3
This repository contains the codes used for POC 3.  
Below are the summary of steps.  
Please refer to the READMEs from their respective directories.  

## Step 1 - goeventhub
This is a cloud application subscribing to Azure Event Hub.  
Please refer to the README in its respective directory.  

## Step 2 - gomqttpubmodule
This is just a lightweight Golang MQTT binary demonstating Device-to-Cloud(D2C).  
Please refer to the README in its respective directory.  

## Step 3 - nodejsmodule
This demonstrates the **device side** of Cloud-to-Device(C2D).  
Please refer to the README in its respective directory.  

## Step 4 - cloudbackend
This demonstrates the **cloud side** of Cloud-to-Device(C2D).  
Please refer to the README in its respective directory.  

## gohttpmodule
This demonstrates a small lightweight Go HTTP server providing APIs for local internal apps in same device.  
It also shows its potential to provide app reliability, as a small watchdog to pre-empt local apps failure caused by insufficient system resources (e.g. disk space).  

## gomqttpubmoduleid
This is just another lightweight Golang MQTT binary demonstating Device-to-Cloud(D2C).  
It demonstrates Go application publishing MQTT message to Azure IoT Hub Topic with **ModuleId**. 
Please refer to the README in its respective directory. 
