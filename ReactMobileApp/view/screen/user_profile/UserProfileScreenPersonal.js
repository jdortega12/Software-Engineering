import React from "react"
import UserProfileScreen from "./UserProfileScreen"

export default function UserProfileScreenPersonal({username}) {
    return (
        <UserProfileScreen username={username} isSelf={true} />
    )
}