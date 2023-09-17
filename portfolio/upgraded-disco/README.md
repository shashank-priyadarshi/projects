Backend for [my portfolio website](ssnk.in)
<br>
This backend has several packages:

<i>Endpoints:

1. graphql: POST <b>planned</b>
2. graphdata: GET
3. githubdata: POST
4. trigger: POST : triggers plugins/integrations with websites like Chess.com, GitHub, etc
5. credentials: POST : currently used to handle login ops
6. register: POST <b>planned</b>
7. resetpassword: POST <b>planned</b>
8. temptoken: POST <b>planned</b>
9. login: POST <b>planned</b>
10. logout: POST <b>planned</b>
    </i>
    <br>

<b>Note:

    - Only the primary server accepts external requests, todos and ghintegration micro-services accept requests from the
      server only
    - All REST endpoints will eventually disallow external connections, only /graphql will be available
    - Internal communication between micro services will migrate from REST to gRPC
    - middleware package might be removed, dependent on requirements after moving to gin
    - [gin-contrib/logger](https://github.com/gin-contrib/logger) for path
      logging, [gin-contrib/cors](https://github.com/gin-contrib/cors) for CORS
      </b>
