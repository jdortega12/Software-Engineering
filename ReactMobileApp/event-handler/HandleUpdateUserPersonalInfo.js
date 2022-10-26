import React from 'react'

// Grabs user info from form and submits to server as JSON.
export default function handleUpdateUserPersonalInfo(firstname, lastname, height, weight) {
    const personalInfo = {
        firstname: firstname,
        lastname: lastname,
        height: height,
        weight: weight,
    }

    try {
        fetch("http://10.0.2.2:8080/api/v1/updatePersonalInfo", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
              },
            body: JSON.stringify(personalInfo),
        })
    } catch(err) {
        console.log(err)
    }
}
