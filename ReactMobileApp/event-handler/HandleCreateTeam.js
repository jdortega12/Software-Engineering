import React from 'react'

// Sends HTTP request to server to create account 
export default function handleCreateTeam(Name, TeamLocation) {
    const teamInfo = {Name: Name, TeamLocation: TeamLocation,}
    console.log(JSON.stringify(teamInfo))
    try{
        fetch("http://10.0.2.2:8080/api/v1/createTeam", {
        method: "POST",
        headers: {
          'Content-type': 'application/json',
        },
        body: JSON.stringify(teamInfo),
          })
        } catch(err){
            console.log(err)
        }
    
}
