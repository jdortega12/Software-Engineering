import React from 'react'

// Sends HTTP request to server to create account 
export default function handleCreateAccount(username, email, password) {
    const accountInfo = {username: username, email: email, password: password,}
    console.log(JSON.stringify(accountInfo))
    try{
        fetch("http://10.0.2.2:8080/api/v1/createAccount", {
        method: "POST",
        headers: {
          'Content-type': 'application/json',
        },
        body: JSON.stringify(accountInfo),
          })
        } catch(err){
            console.log(err)
        }
    
}