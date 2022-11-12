import React from 'react'

// Sends HTTP request to server to accept or deny 
export default function handleRemovePlayer(managerUsername, playerUsername, accepted) {
    if(accepted === 'Y')
    {
    const data = {managername: managerUsername, playername: playerUsername}
    console.log(data.managername)
    console.log(data.playername)
    console.log(JSON.stringify(data))
    try{
        fetch("http://10.0.2.2:8080/api/v1/removePlayer", {
        method: "POST",
        headers: {
          'Content-type': 'application/json',
        },
        body: JSON.stringify(data),
          })
        } catch(err){
            console.log(err)
        }
    }
    else
    {
        console.log("Player removal was invalid")
    }
}