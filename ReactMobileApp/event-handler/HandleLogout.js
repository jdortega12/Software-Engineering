import React from "react-native"

// Sends HTTP request to server to logout user/
// wipe current session. 
export default function handleLogout() {
    try {
        fetch("http://10.0.2.2:8080/api/v1/logout", {
            method: "POST",
        })

        sessionStorage.clear()
    } catch(err) {
        console.log(err)
    }
}