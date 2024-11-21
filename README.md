# File share
![image](https://github.com/user-attachments/assets/44769acf-879a-409f-9a41-cf7f2b4d4392)
Simple file transfer server written in golang. Transfer files between devices in a local network.

## Setup
```
git clone https://github.com/me7alix/file-share-go
cd file-share-go
go build app.go
```

## Usage
To transfer a file, you need to start the server. Find out what the IP address of your server is (on Linux you can find it out with the command `ip addr`). Then open the website on a device that is on the local network using the IP address.
Now you can send files from a device to the server. To send files in the other direction you need to run the server on a device too. On Android devices you can do this through the Termux application. Uploaded files will be saved to the `uploads` folder. You can change the folder in the app.go
