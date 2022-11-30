// Gets the teams from handlers 
export default async function handlGetUserProfile() {
    let teams = null

    try {
        const response = await fetch('http://10.0.2.2:8080/api/v1/getTeams/')
        teams = await response.json()
    } 
    catch (error) {
        console.log(error)
    }
    return teams
}

