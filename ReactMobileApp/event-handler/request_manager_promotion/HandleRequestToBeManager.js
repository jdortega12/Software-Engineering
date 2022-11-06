import { END_REQUEST_TO_BE_MANAGER } from "../../GlobalConstants"

export default async function handleRequestToBeManager(message) {
    data = {
        "message": message,
    }

    fetch(END_REQUEST_TO_BE_MANAGER, {
        method:'POST',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    }).then((response) => {
        return response.status
    }).catch((error) => {
        console.error(error)
        return error
    })
}