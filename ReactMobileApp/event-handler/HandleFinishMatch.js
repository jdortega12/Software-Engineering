export default async function handleFinishMatch(id) {
    body = JSON.stringify({
        id: id,
    })

    try {
        response = await fetch("http://10.0.2.2:8080/api/v1/finish-match", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: body,
        })

        return response.status
    }
    catch(error) {
        console.log(error)
        return error
    }
}