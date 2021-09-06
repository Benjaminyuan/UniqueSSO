# UniqueSSO
> Single Sign On for UniqueStudio

## Diagram 

The UniqueSSO is nearly a standard implementation of CAS.

Below is the cas diagram.

![CAS Diagram](https://apereo.github.io/cas/4.2.x/images/cas_flow_diagram.png)
## Big picture

1. login at `POST /cas/login?service=${redirectURI}` with body 
2. validate ticket at `GET /cas/p3/serviceValidate?ticket=${ticket}`

for login, there are four ways to login:

1. phone number with password

2. phone sms

3. email address with password

4. wechat oauth

store state in cookie, which persisted by redis.

The user info is stored in PostgreSQL with database named `sso`, and the table name is `user`

## How to access

The UniqueSSO is nearly a standard implementation of CAS. This is the [cas link](https://apereo.github.io/cas/4.2.x/protocol/CAS-Protocol.html)

1. Redirect to UniqueSSO login page `https://sso.hustuniuqe.com/cas/login` with service, which is the redirectURI from SSO.
2. If user login successfully, the `UniqueSSO` will redirct the page to `service` specified in step 1 and with the ticket. Like this:
   > `https://bbs.hustunique.com?ticket=${TICKET}`
   
   For most cases, the ticket will expire after 3 minutes. In addition, the ticket is just valid at the first time whether validate successfully or not.
3. Validate ticket by sending HTTP GET request to `https://sso.hustuniuqe.com/cas/p3/serviceValidate?ticket=${ticket}&service=${service}`. If success, sso will return the user info
   >  The service here is used to fiter not redirect.

## Deployment

1. edit the backend config file 


## TODO list

- [ ] Access APM systems

## Uniform Response

```json
{
  "serviceResponse": {
    "authenticationFailure": {
      "code": "",
      "description": ""
    },
    "authenticationSuccess": {
      "user": "${UID}",
      "attributes": {
        "uid": "",
        "name": "",
        "phone": "",
        "email": ""
      }
    }
  }
}
```