# CLI2API

The project allows you to HTTPS to run command line utilities securely. You can have configure each command that can be 
executed, and you can define different secret keys per command to provide different access to different clients.
See more details in configs/cli2api.yaml

The API request syntax looks like this:
  
    curl --header "Content-Type: application/json" \
      --cacert cacert.pem \
      --request POST \
      --data '{"args":["-la", "/Users/username"]}' \
      https://localhost:8000/cli/ls

## Generating SSL Keys
The cli2api requires TLS, as otherwise it would be too insecure to execute commands / send things via plain text.
You might want to set up LetsEncrypt and use it to have fully validated certs, but you can also use self signed 
certificates.

    openssl req  -new  -newkey rsa:2048  -nodes  -keyout localhost.key  -out localhost.csr
    openssl  x509  -req  -days 365  -in localhost.csr  -signkey localhost.key  -out localhost.crt

## Using Curl to send request
If using self signed certificates, start by getting certificate PEM file so curl doesn't produce error:
    
    
    API_HOST=localhost:8000
    echo quit | openssl s_client -showcerts -servername "${API_HOST}" -connect "${API_HOST}" > cacert.pem
  
Use curl to make requests: 
  
    API_HOST=localhost:8000
    curl --header "Content-Type: application/json" \
      --cacert cacert.pem \
      --request POST \
      --data '{"args":["-la", "/Users/username"]}' \
      https://${API_HOST}/cli/ls
      
      
## What is next
Few things that needs to be done
* Right now both STDOUT & STDERR from the command execution combined and returned as the result. We might want to
provide control over it to administrator. Maybe wrap stdout/stderr in json.
* Add RunAsUser/RunAsGroup config option for commands, so that server can be started as root, and would drop 
permissions/change user when needed
* Provide ability to set up max concurrent instances of a particular command, as well as total number of commands
server can run concurrently. Ask user to retry later when concurrent limit reached.
* Add command execution timeouts parameter.
* Allow to set limits json body passed via POST.
* Allow to pass environment variables, and current working directory via request.
* Allow to set environment variables and current working directory per command in config.
* Allow to set arguments as config, as addition to arguments passed in request, or instead of them (preventing 
arguments via request)

