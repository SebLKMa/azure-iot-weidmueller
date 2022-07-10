# goeventhub
Azure Event Hubs automatically capture streaming data and save it to a storage account. You can have data retention up to 7 days. You can horizontal scale using partitions. For details, please refer to:  
https://azure.microsoft.com/en-us/services/event-hubs/  
https://azure.microsoft.com/en-us/pricing/details/event-hubs/  

This module demonstrates a Cloud application receiving messages from Device. 
The codes used in this module have been adapted from these official sites:  
https://github.com/Azure/azure-event-hubs-go  
https://docs.microsoft.com/en-us/azure/event-hubs/event-hubs-go-get-started-send  



Before building the app, please do your own `go mod init` and `go mod tidy`.  

Start `goeventhub`  to start receiving messages.  
```sh
$ ./goeventhub 
Connected to event hub - sb://...namespace.servicebus.windows.net/
Hello from weidmueller-model-x 2021.05.18 09:34:13
Hello from weidmueller-model-x 2021.05.18 09:34:24
Hello from weidmueller-model-x 2021.05.18 09:34:35
Hello from weidmueller-model-x 2021.05.18 09:34:45
...
```

## Using az CLI
Another general way to monitor IoT Hub Events is to use `az` CLI:  
```sh
az iot hub monitor-events -n seb-hub
```
