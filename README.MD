# Discord as CDN
![logo](./logo.jpeg)


## How to run your project
first create a `.env` file and add your bot token, you can follow with the `.env.example` file

then run the following commands

```bash
go mod download
go run main.go
```
then you can do a call to the ``/sendFile`` endpoint, passing the file as form-data with the name of ``file``, you can check the example in the ``example`` folder

when you send the file, the server will answer with the url of the file like this:

```json
{
"url": "link"
}

```
## Why?

It's a free service to store your photos dynamically, instead of hosting on your server and spending more money. 

As said one of my friends([NicolasLopes7](https://github.com/NicolasLopes7)) and owner of the [repo](https://github.com/NicolasLopes7/dontUseCDN-UseDiscord) that I took inspiration from.


## How to get .env information

