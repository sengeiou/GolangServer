swagger: "2.0"
info:
  title: Sample API
  description: API description in Markdown.
  version: 1.0.0

host: localhost:5000
basePath: /
schemes:
  - https
securityDefinitions:
  Bearer:
    type: apiKey
    name: Authorization
    in: header
paths:
  /api/user/manual-signup:
    post:
      tags:
      - "Auth / Manula signup"
      summary: "manual signup"
      description: ""
      operationId: "addPet"
      consumes:
      - "application/json"
      - "application/xml"
      produces:
      - "application/xml"
      - "application/json"
      parameters:
      - in: "body"
        name: "body"
        description: "Pet object that needs to be added to the store"
        required: true
        schema:
              type: object
              properties:
                firstName:
                  type: string
                lastName:
                  type: string
                dateOfBirth:
                  type: string
                email:
                  type: string
                password:
                  type: string
              example:   # Sample object
                firstName: Ahmed
                lastName : Khalaf
                dateOfBirth : "15-8-1996"
                email: ahmed@ahmed.com
                password: ahmed5747723

      responses:
        "405":
          description: "Invalid input"
        "400":
          description: "Invalid input"
        "208":
          description: "Email Already Exist"
        "500":
          description: "Server Error Can't crate Now User"

  api/user/manual-signin:
    post:
      tags:
      - "Auth / Manula login"
      summary: "manual login"
      description: ""
      operationId: "addPet"
      consumes:
      - "application/json"
      - "application/xml"
      produces:
      - "application/xml"
      - "application/json"
      parameters:
      - in: "body"
        name: "body"
        description: "Pet object that needs to be added to the store"
        required: true
        schema:
              type: object
              properties:
                email:
                  type: string
                password:
                  type: string
              example:   # Sample object
                email: ahmed@ahmed.com
                password: ahmed5747723

      responses:
        "405":
          description: "Invalid input"
        "400":
          description: "Invalid input"
        "208":
          description: "Email Already Exist"
        "500":
          description: "Server Error Can't crate Now User"
  
    
  
  /api/user/social-Google-Apple-Facebook:
    post:
      tags:
      - "Auth /Social login & Register"
      summary: "Social login & Register"
      description: ""
      operationId: "addPet"
      consumes:
      - "application/json"
      - "application/xml"
      produces:
      - "application/xml"
      - "application/json"
      parameters:
      - in: "body"
        name: "body"
        description: "Pet object that needs to be added to the store"
        required: true
        schema:
              type: object
              properties:
                email:
                  type: string
                firstName:
                  type: string
                lastName:
                  type: string
                token:
                  type: string
              example:   # Sample object
                email: ahmed@ahmed.com
                firstName: ahmed
                lastName: khalaf
                token: "tokendataExample || NOT Requerd"

      responses:
        "405":
          description: "Invalid input"
        "400":
          description: "Invalid input"
        "403":
          description: "can't Register new user"
        "208":
          description: "Email Already Exist"
        "500":
          description: "Server Error Can't login "
  
    
  /api/user/Auth-Phone-Number:
    post:
      tags:
      - "Auth / Phone Number login & Register"
      summary: "PhoneNumber login & Register"
      description: ""
      operationId: "addPet"
      consumes:
      - "application/json"
      - "application/xml"
      produces:
      - "application/xml"
      - "application/json"
      parameters:
      - in: "body"
        name: "body"
        description: "Pet object that needs to be added to the store"
        required: true
        schema:
              type: object
              properties:
                phoneNumber:
                  type: string
                firstName:
                  type: string
                lastName:
                  type: string
                password:
                  type: string
              example:   # Sample object
                email: "01278592995"
                firstName: sara
                lastName: belosy
                password: "sara55555"

      responses:
        "405":
          description: "Invalid input"
        "400":
          description: "Invalid input"
        "403":
          description: "can't Register new user"
        "208":
          description: "Email Already Exist"
        "500":
          description: "Server Error Can't login "
  

  /api/user/UpdateInfo/{id}:

    post:
      security:
      - Bearer: []
      tags:
      - "Updated User Data"
      summary: "Update user data"
      description: ""
      operationId: "addPet"
      consumes:
      - "application/json"
      - "application/xml"
      produces:
      - "application/xml"
      - "application/json"
      parameters:
      - name: "id"
        in: "path"
      - in: "body"
        name: "body"
        description: "Pet object that needs to be added to the store"
        required: true
        schema:
              type: object
              properties:
                # phoneNumber:
                #   type: string
                # firstName:
                #   type: string
                # lastName:
                #   type: string
                # password:
                #   type: string

                firstname: 
                 type: string
                firstName: 
                  type: string
                dateOfBirth: 
                  type: string
                email: 
                  type: string
                phoneNumber: 
                  type: string
                statusOfUser: 
                  type: string
                userProfilePhoto: 
                  type: string
                userHeight: 
                  "33"
                Gender: 
                  type: string
                orientation: 
                  type: string
                userRelationship: 
                  type: string
                userLookFor: 
                  type: string
                userRole: 
                  type: string
                school: 
                  type: string
                placeOfJob: 
                  type: string
                jobTitle: 
                  type: string
                userCurrentMood: 
                  type: string
                userMediaPhoto:
                  type: array
                userMediaVideo:
                   type: array
                userLocation:
                  type: array
                userLanguages :
                  type: array
                userHobbyes :
                  type: array
                userZodlac :
                  type: array
                userPets :
                  type: array
                isUserSmoking :
                  type: array
                physlcalAttraction :
                  type: array
                userTurnON :
                  type: array
                userStyle :
                  type: array
                userBody :
                  type: array
                userSesson :
                  type: array
                userMovies :
                  type: array
                userDrink :
                  type: array


              example:   # Sample object
                firstname: "Ahmed "
                firstName: "Updated Ahmed"
                dateOfBirth: "7-4-9988"
                email: "ahmed@ahmed.com"
                phoneNumber: "testing..any data"
                statusOfUser: "testing..any data"
                userProfilePhoto: "testing..any data"
                userHeight: "33"
                Gender: "testing..any data"
                orientation: "testing..any data"
                userRelationship: "testing..any data"
                userLookFor: "testing..any data"
                userRole: "testing..any data"
                school: "testing..any data"
                placeOfJob: "testing..any data"
                jobTitle: "testing..any data"
                userCurrentMood: "testing..any data"
                userMediaVideo: ["first media"]
                userLocation: ["country","USA", "city","NYC" ]
                userLanguages: ["test"]
                userHobbyes : ["test"]
                userZodlac : ["test"]
                userPets : ["test"]
                isUserSmoking : ["test"]
                physlcalAttraction : ["test"]
                userTurnON : ["test"]
                userStyle : ["test"]
                userBody : ["test"]
                userSesson : ["test"]
                userMovies : ["test"]
                userDrink : ["test"]
                

      responses:
        "405":
          description: "Invalid input"
        "400":
          description: "Invalid input"
        "403":
          description: "can't Register new user"
        "208":
          description: "Email Already Exist"
        "500":
          description: "Server Error Can't Update Userdata "
  
    
    
    
