MainEndPoint URL =  http://ec2-35-90-201-145.us-west-2.compute.amazonaws.com:5000

# manual signup
Method [Post]  End Point "URL + api/user/manual-signup"
pyload Example 
Shold be JsonData
{
		firstName   string
		lastName    string
		dateOfBirth string
		email       string
		password    string	
}

# manual signin
Method [Post]	EndPoint "URL + api/user/manual-signin"
pyload Example 
Shold be JsonData
{
		email       string
		password    string	
}

# Social login & Register At The Same Time using Same EndPoint

Method [Post]	EndPoint "URL + api/user/social-Google-Apple-Facebook"
pyload Example 
Shold be JsonData
{
		email       string
		firstName   string
		lastName    string
        token string ! Not Required 
}

# Auth login & Register With Phone Number

Method [Post]	EndPoint "URL + api/user/Auth-Phone-Number"

# pyload Example 
Shold be JsonData
{
		phoneNumber string
		Password    string
		firstName   string
		lastName    string
}

# Update User Data Any Time or After Siginup
# we should have the token bearar and send it as Authorization header & user ID
 
Method [Patch]	EndPoint "URL + api/user/UpdateInfo/:id"

Sending Headers to the end points with peare token 
{Authorization:"Bearer (TokenString)"}
# Example

{Authorization:"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJleHAiOjE2Njgw
NjIxNTAsImlzcyI6IjYzNmM3ZTQ0MzVkOTRiODdjNGIxVhYyJ9.Em2c3IBrXz1J_7tdZhP5Q6L_dc5ps0EfV332hFS8RX0"}

# payload Example 1
{
    "firstname": "Ahmed ",
    "lastname":"test",
    "dateOfBirth": "7-4-9988"
}

# Pyload Example 2 ! you can change any number of felds as you need just passed to the end point 
Shold be JsonData

{
    "firstname": "Ahmed ",
    "lastname":"test",
    "dateOfBirth": "7-4-9988",
    "email": "s@s.com",
    "phoneNumber": "testing..any data",
    "statusOfUser": "testing..any data",
    "userProfilePhoto": "testing..any data",
    "userHeight": "33",
    "Gender": "testing..any data",
    "orientation": "testing..any data",
    "userRelationship": "testing..any data",
    "userLookFor": "testing..any data",
    "userRole": "testing..any data",
    "school": "testing..any data",
    "placeOfJob": "testing..any data",
    "jobTitle": "testing..any data",
    "userCurrentMood": "testing..any data",
    "userMediaPhoto":["test"],
    "userMediaVideo": ["first media"],
    "userLocation": ["country","USA", "city","NYC" ],
    "userLanguages" :["test"],
    "userHobbyes" :["test"],
    "userZodlac" :["test"],
    "userPets" :["test"],
    "isUserSmoking" :["test"],
    "physlcalAttraction" :["test"],
    "userTurnON" :["test"],
    "userStyle" :["test"],
    "userBody" :["test"],
    "userSesson" :["test"],
    "userMovies" :["test"],
    "userDrink" :["test"]
}


