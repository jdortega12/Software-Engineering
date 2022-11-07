import { END_GET_PROMOTION_TO_MANAGER_REQS } from "../../GlobalConstants"

// Sends GET request to server for the promotion requests in the db
export default async function handleGetPromotionToManagerRequests() {
    requests = {}

    try {
        response = await fetch(END_GET_PROMOTION_TO_MANAGER_REQS)
        requests = await response.json()
    }
    catch(error) {
        console.log(error)
    }

    return requests
}