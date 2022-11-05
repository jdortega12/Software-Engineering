import React from "react"
import UserProfileScreen from "./UserProfileScreen"

export default function UserProfileScreenNotPersonal({username}) {
    return (
        <UserProfileScreen username={username} isSelf={false} />
    )
}