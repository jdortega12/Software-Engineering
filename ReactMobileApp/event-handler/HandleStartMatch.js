export default async function handleStartMatch(homeTeam, awayTeam, location) {
    body = JSON.stringify({
        home_team_name: homeTeam,
        away_team_name: awayTeam,
        location: location,
    })

    try {
        response = await fetch("http://10.0.2.2:8080/api/v1/start-match", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: body,
        })

        return response.status
    } 
    catch(err) {
        console.log(err)
        return err
    }
}