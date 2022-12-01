import React from 'react'

// Sends HTTP request to server to create account 
export default function handleChangeRoster(userID, teamname) {
    const info = {userID: userID, teamname: teamname}
    console.log(JSON.stringify(info))
    try{
        fetch("http://10.0.2.2:8080/api/v1/changeRoster", {
        method: "POST",
        headers: {
          'Content-type': 'application/json',
        },
        body: JSON.stringify(info),
          })
        } catch(err){
            console.log(err)
        }
    
}