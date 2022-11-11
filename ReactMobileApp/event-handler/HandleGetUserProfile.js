// Gets the data for a user's profile and displays it 
export default async function handlGetUserProfile(username) {
    let userData = null

    try {
        const response = await fetch('http://10.0.2.2:8080/api/v1/get-user/' + username)
        userData = await response.json()
    } 
    catch (error) {
        console.log(error)
    }
    
    return userData
}

