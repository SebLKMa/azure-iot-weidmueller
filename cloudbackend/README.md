# cloudbackend

The source code in this module is adapted from the official [Azure Node.js documentation](https://docs.microsoft.com/en-us/azure/iot-hub/quickstart-control-device-node).

This demonstrates Cloud-to-Device(C2D) using Azure IoT **Direct Method** calls.  
Please note to the `connectionString` and `deviceId` variables used in the codes.  

Example cloud application setting weidmuller-iot-device its send interval to 20 seconds:  
```sh
$ node CloudToGW30.js 
Response from SetTelemetryInterval on weidmueller-iot-device:  
{
  "status": 200,
  "payload": "Telemetry interval set: 20"
}
```
