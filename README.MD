http://localhost:9001/api/users/register

{
    "age": -435309917,
    "email": "dilshoddilmurodov112@gmail.com", /haqiqiry email bulishi kk
    "name": "qui",
    "password": "ex dolore"
}

http://localhost:9001/api/users/verify

{
    "email":"dilshoddilmurodov915@gmail.com",
    "code": "604527"
}

http://localhost:9001/api/users/login


{
    "email": "",
    "password": ""
}



http://localhost:9001/api/users/2   GET  TOKEN BILAN




http://localhost:9001/api/users/2 DELETE TOKEN BILAN





HOTEL 

http://localhost:8080/api/hotels/create

{
    "address": "laborum",
    "location": "aliqua",
    "name": "exercitation eiusmod Lorem ad",
    "rating": 5,
    "rooms": [
        {
            "pricePerNight": 234,
            "availability": true,
            "roomType": "lux"
        }
    ]
}

HOTEL GET WITH ID
http://localhost:9001/api/hotels/3


HOTEL GET ALL HOTELS 
http://localhost:9001/api/hotels



AVAILABLE HOTEL
HTTP://localhost:9001/api/hotels/rooms/availabilty/:id



http:localhost:9001/api/booking

{
    "user_id": 2,
    "hotel_id": 4,
    "roomtype": "viphttp",
    "check_in_date": {
        "seconds": 1640995200,
        "nanos": 0
    },
    "check_out_date": {
        "seconds": 1642211600,
        "nanos": 0
    },
    "total_amount": 250.75
}


http://localhost9001/update/booking
{
    "booking_id": 67,
    "check_in_date": {
        "nanos": -24758650,
        "seconds": "8427977"
    },
    "check_out_date": {
        "nanos": 1796217347,
        "seconds": "11672827"
    },
    "room_type": "mollit non"
}