import React from 'react'

// Sends HTTP request to server to create account
export default function handleLogin(username, password) {
    accountInfo = {username: username, password: password}
    try {
        fetch("http://10.0.2.2:8080/api/v1/login", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
              },
            body: JSON.stringify(accountInfo),
        })
    } catch(err) {
        console.log(err)
    }
    }
